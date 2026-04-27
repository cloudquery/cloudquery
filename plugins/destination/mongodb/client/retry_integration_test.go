package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// flakyProxy is a TCP proxy that forwards traffic to upstream and can be put
// into a "failing" mode where all new connections are dropped immediately and
// existing connections are continuously severed. While failing, the MongoDB
// driver observes the same NetworkError-labeled "broken pipe" / EOF errors
// that surface in production (ENG-3281). When failing is toggled off, the
// proxy forwards normally so retries can succeed.
type flakyProxy struct {
	upstream string
	listener net.Listener
	failing  atomic.Bool

	mu      sync.Mutex
	conns   []net.Conn
	stopped bool
}

func newFlakyProxy(t *testing.T, upstream string) *flakyProxy {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	p := &flakyProxy{upstream: upstream, listener: l}
	go p.serve()
	t.Cleanup(p.close)
	return p
}

func (p *flakyProxy) addr() string { return p.listener.Addr().String() }

// failFor puts the proxy into failing mode for d, then restores pass-through.
// Returns a channel that closes when the failing window ends.
func (p *flakyProxy) failFor(d time.Duration) <-chan struct{} {
	done := make(chan struct{})
	p.failing.Store(true)
	p.severAll()
	go func() {
		defer close(done)
		deadline := time.Now().Add(d)
		// Continuously sever new connections that snuck in during the window.
		for time.Now().Before(deadline) {
			time.Sleep(10 * time.Millisecond)
			p.severAll()
		}
		p.failing.Store(false)
	}()
	return done
}

func (p *flakyProxy) severAll() {
	p.mu.Lock()
	conns := p.conns
	p.conns = nil
	p.mu.Unlock()
	for _, c := range conns {
		_ = c.Close()
	}
}

func (p *flakyProxy) close() {
	p.mu.Lock()
	if p.stopped {
		p.mu.Unlock()
		return
	}
	p.stopped = true
	conns := p.conns
	p.conns = nil
	p.mu.Unlock()
	_ = p.listener.Close()
	for _, c := range conns {
		_ = c.Close()
	}
}

func (p *flakyProxy) trackConn(c net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.stopped {
		_ = c.Close()
		return
	}
	p.conns = append(p.conns, c)
}

func (p *flakyProxy) serve() {
	for {
		client, err := p.listener.Accept()
		if err != nil {
			return
		}
		if p.failing.Load() {
			_ = client.Close()
			continue
		}
		server, err := net.Dial("tcp", p.upstream)
		if err != nil {
			_ = client.Close()
			continue
		}
		p.trackConn(client)
		p.trackConn(server)
		go func() {
			_, _ = io.Copy(server, client)
			_ = server.Close()
		}()
		go func() {
			_, _ = io.Copy(client, server)
			_ = client.Close()
		}()
	}
}

func upstreamHostPort(t *testing.T, connectionString string) string {
	t.Helper()
	// Strip the "mongodb://" scheme and any path / query, return host:port.
	u, err := url.Parse(connectionString)
	require.NoError(t, err)
	host := u.Host
	if host == "" {
		host = strings.TrimPrefix(connectionString, "mongodb://")
		if i := strings.IndexAny(host, "/?"); i >= 0 {
			host = host[:i]
		}
	}
	if !strings.Contains(host, ":") {
		host += ":27017"
	}
	return host
}

func newRetryReproClient(t *testing.T, retry *spec.WriteRetryConfig) (*Client, *flakyProxy) {
	t.Helper()
	upstream := upstreamHostPort(t, getTestConnection())
	proxy := newFlakyProxy(t, upstream)

	ctx := context.Background()
	s := &spec.Spec{
		// Pool size 1 keeps the topology simple; serverSelectionTimeoutMS keeps
		// the test fast when the proxy is in failing mode.
		ConnectionString: "mongodb://" + proxy.addr() + "/?maxPoolSize=1&serverSelectionTimeoutMS=2000",
		Database:         "destination_mongodb_retry_repro_test",
		WriteRetry:       retry,
	}
	specBytes, err := json.Marshal(s)
	require.NoError(t, err)

	logger := zerolog.New(zerolog.NewTestWriter(t))
	pc, err := New(ctx, logger, specBytes, plugin.NewClientOptions{})
	require.NoError(t, err)
	c := pc.(*Client)
	t.Cleanup(func() {
		_ = c.client.Database(s.Database).Drop(ctx)
		_ = pc.Close(ctx)
	})
	return c, proxy
}

var retryReproTable = &schema.Table{
	Name: "retry_repro",
	Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "val", Type: arrow.BinaryTypes.String},
	},
}

// TestRetryAbsorbsConnectionDrop is a deterministic regression test for
// ENG-3281. While the driver is busy executing a write, a TCP proxy in front
// of MongoDB severs all connections for ~300ms (the same NetworkError shape
// the customer sees in production). The retry wrapper must absorb the
// failures and complete the write once the proxy stops failing.
func TestRetryAbsorbsConnectionDrop(t *testing.T) {
	maxBackoff := configtype.NewDuration(500 * time.Millisecond)
	c, proxy := newRetryReproClient(t, &spec.WriteRetryConfig{MaxAttempts: 8, MaxBackoff: &maxBackoff})

	docs := []any{bson.M{"id": int64(1), "val": "a"}}

	// Start the write and the failure window concurrently. The write will
	// observe broken pipes / EOF on its first attempt(s), back off, and
	// eventually retry past the failure window onto a healthy connection.
	failingDone := proxy.failFor(300 * time.Millisecond)

	require.NoError(t, c.overwriteTableBatch(context.Background(), retryReproTable, docs))
	<-failingDone
}

// TestFailureInjectionReachesWritePath is the negative control for
// TestRetryAbsorbsConnectionDrop: with retries disabled (MaxAttempts=1) the
// failing proxy must cause the write to fail. This guards against the
// failure injection being silently absorbed by background driver activity
// (heartbeats, topology probes) and never actually reaching the write path.
func TestFailureInjectionReachesWritePath(t *testing.T) {
	maxBackoff := configtype.NewDuration(50 * time.Millisecond)
	c, proxy := newRetryReproClient(t, &spec.WriteRetryConfig{MaxAttempts: 1, MaxBackoff: &maxBackoff})

	proxy.failFor(2 * time.Second)
	err := c.overwriteTableBatch(context.Background(), retryReproTable, []any{bson.M{"id": int64(1)}})
	require.Error(t, err, "expected failure without retry, but write succeeded")
}

// TestRetryGivesUpWhenAllAttemptsFail asserts that when failures outlast the
// configured retry budget, the error propagates with a network-error shape.
func TestRetryGivesUpWhenAllAttemptsFail(t *testing.T) {
	maxBackoff := configtype.NewDuration(50 * time.Millisecond)
	c, proxy := newRetryReproClient(t, &spec.WriteRetryConfig{MaxAttempts: 2, MaxBackoff: &maxBackoff})

	docs := []any{bson.M{"id": int64(1)}}

	// Hold the failure mode for longer than any plausible retry budget.
	proxy.failFor(10 * time.Second)

	err := c.overwriteTableBatch(context.Background(), retryReproTable, docs)
	require.Error(t, err)
	// The final error chain should still be classifiable as retryable;
	// we just ran out of attempts.
	require.True(t,
		isRetryableWriteError(err) || errors.Is(err, context.DeadlineExceeded),
		"expected network/timeout error, got: %v", err,
	)
}

package client

import (
	"bytes"
	"context"
	"encoding/json"
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

type flakyProxy struct {
	upstream string
	listener net.Listener

	mu           sync.Mutex
	failuresLeft int
	drops        int
	conns        []net.Conn
	stopped      bool
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

func (p *flakyProxy) dropNext(n int) {
	p.mu.Lock()
	p.failuresLeft = n
	p.drops = 0
	conns := p.conns
	p.conns = nil
	p.mu.Unlock()
	for _, c := range conns {
		_ = c.Close()
	}
}

func (p *flakyProxy) dropsCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.drops
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

func (p *flakyProxy) serve() {
	for {
		client, err := p.listener.Accept()
		if err != nil {
			return
		}
		p.mu.Lock()
		if p.failuresLeft > 0 {
			p.failuresLeft--
			p.drops++
			p.mu.Unlock()
			_ = client.Close()
			continue
		}
		if p.stopped {
			p.mu.Unlock()
			_ = client.Close()
			return
		}
		p.mu.Unlock()

		server, err := net.Dial("tcp", p.upstream)
		if err != nil {
			_ = client.Close()
			continue
		}
		p.mu.Lock()
		p.conns = append(p.conns, client, server)
		p.mu.Unlock()
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

type retryLogCounter struct {
	inner io.Writer
	n     atomic.Int32
}

func (w *retryLogCounter) Write(p []byte) (int, error) {
	if bytes.Contains(p, []byte("retrying MongoDB write")) {
		w.n.Add(1)
	}
	if w.inner != nil {
		return w.inner.Write(p)
	}
	return len(p), nil
}

func (w *retryLogCounter) Count() int { return int(w.n.Load()) }

func newRetryReproClient(t *testing.T, retry *spec.WriteRetryConfig) (*Client, *flakyProxy, *retryLogCounter) {
	t.Helper()
	upstream := upstreamHostPort(t, getTestConnection())
	proxy := newFlakyProxy(t, upstream)

	ctx := context.Background()
	s := &spec.Spec{
		ConnectionString: "mongodb://" + proxy.addr() + "/?retryWrites=false&maxPoolSize=1&serverSelectionTimeoutMS=200&directConnection=true",
		Database:         "destination_mongodb_retry_repro_test",
		WriteRetry:       retry,
	}
	specBytes, err := json.Marshal(s)
	require.NoError(t, err)

	counter := &retryLogCounter{inner: zerolog.NewTestWriter(t)}
	logger := zerolog.New(counter)
	pc, err := New(ctx, logger, specBytes, plugin.NewClientOptions{})
	require.NoError(t, err)
	c := pc.(*Client)
	t.Cleanup(func() {
		_ = c.client.Database(s.Database).Drop(ctx)
		_ = pc.Close(ctx)
	})
	return c, proxy, counter
}

var retryReproTable = &schema.Table{
	Name: "retry_repro",
	Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "val", Type: arrow.BinaryTypes.String},
	},
}

func TestRetryAbsorbsConnectionDrop(t *testing.T) {
	const drops = 2

	maxBackoff := configtype.NewDuration(20 * time.Millisecond)
	c, proxy, retries := newRetryReproClient(t, &spec.WriteRetryConfig{
		MaxAttempts: 100,
		MaxBackoff:  &maxBackoff,
	})

	proxy.dropNext(drops)

	require.NoError(t, c.overwriteTableBatch(
		context.Background(), retryReproTable,
		[]any{bson.M{"id": int64(1), "val": "a"}},
	))
	require.Equal(t, drops, proxy.dropsCount(), "proxy should have consumed its drop budget")
	require.GreaterOrEqual(t, retries.Count(), drops, "retry-go should fire at least once per dropped connection")
}

func TestFailureInjectionReachesWritePath(t *testing.T) {
	maxBackoff := configtype.NewDuration(50 * time.Millisecond)
	c, proxy, retries := newRetryReproClient(t, &spec.WriteRetryConfig{MaxAttempts: 1, MaxBackoff: &maxBackoff})

	proxy.dropNext(100)
	err := c.overwriteTableBatch(context.Background(), retryReproTable, []any{bson.M{"id": int64(1)}})
	require.Error(t, err, "expected failure without retry, but write succeeded")
	require.True(t, isRetryableWriteError(err), "expected a retryable network error, got: %v", err)
	require.Equal(t, 0, retries.Count(), "OnRetry must not fire when MaxAttempts=1")
}

func TestRetryGivesUpWhenAllAttemptsFail(t *testing.T) {
	maxBackoff := configtype.NewDuration(50 * time.Millisecond)
	const maxAttempts = 3
	c, proxy, retries := newRetryReproClient(t, &spec.WriteRetryConfig{MaxAttempts: maxAttempts, MaxBackoff: &maxBackoff})

	proxy.dropNext(10_000)

	err := c.overwriteTableBatch(context.Background(), retryReproTable, []any{bson.M{"id": int64(1)}})
	require.Error(t, err)
	require.True(t, isRetryableWriteError(err), "expected a retryable network error, got: %v", err)
	require.Equal(t, maxAttempts, retries.Count(), "OnRetry should fire once per failing attempt")
}

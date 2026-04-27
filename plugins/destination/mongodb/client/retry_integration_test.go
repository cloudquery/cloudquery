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

// flakyProxy is a TCP proxy that forwards traffic to upstream and can be told
// to drop a specific number of inbound connections (closing them immediately
// before any bytes flow). While failures are armed, the MongoDB driver
// observes the same NetworkError-labeled "broken pipe" / EOF errors that
// surface in production (ENG-3281). Once the failure budget is consumed, the
// proxy forwards normally so retries can succeed. Tests assert the exact
// number of drops to prove the retry layer actually fired.
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

// dropNext arms the proxy to drop the next n inbound connections. Existing
// connections are also severed so the driver must dial fresh ones (which then
// consume the drop budget).
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

// drops returns how many inbound connections have been dropped since the last
// dropNext call.
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

// retryLogCounter counts log lines emitted by the retryWrite OnRetry callback
// so tests can assert how many times the app-level retry actually fired
// (independent of how many connection attempts the driver made internally).
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
		// retryWrites=false disables the driver's own single retry so the
		// retry-go layer is the only place writes get a second chance.
		// serverSelectionTimeoutMS=200 makes each failing op() return quickly
		// (the driver's internal server-selection loop will time out fast and
		// surface the error to retry-go rather than masking it).
		ConnectionString: "mongodb://" + proxy.addr() + "/?retryWrites=false&maxPoolSize=1&serverSelectionTimeoutMS=200",
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

// TestRetryAbsorbsConnectionDrop is a deterministic regression test for
// ENG-3281. The proxy drops enough connections that the driver's
// server-selection layer can't recover within a single op() call, forcing
// retry-go to do the recovery. The drop budget is then exhausted and a later
// op() succeeds. We assert that retry-go's OnRetry callback fired -- that is
// the only direct signal that maps to retry-go invocations (drop counts get
// inflated by the driver's internal server-selection retry loop).
func TestRetryAbsorbsConnectionDrop(t *testing.T) {
	const drops = 2
	// Empirically each drop produces two failed op() invocations against
	// driver v2.5.0: the first dials a fresh connection (hits the proxy
	// drop), and the second sees the now-broken pool entry and fails fast
	// without dialing. Then the pool dials a fresh conn that the proxy lets
	// through, and op() succeeds. If a future driver upgrade changes pool
	// dial-on-checkout behavior this constant may need updating.
	const expectedRetries = drops * 2

	maxBackoff := configtype.NewDuration(20 * time.Millisecond)
	c, proxy, retries := newRetryReproClient(t, &spec.WriteRetryConfig{
		MaxAttempts: expectedRetries + 5, // headroom so retry-go isn't the bottleneck
		MaxBackoff:  &maxBackoff,
	})

	proxy.dropNext(drops)

	require.NoError(t, c.overwriteTableBatch(
		context.Background(), retryReproTable,
		[]any{bson.M{"id": int64(1), "val": "a"}},
	))
	require.Equal(t, drops, proxy.dropsCount(), "proxy should have consumed its drop budget")
	require.Equal(t, expectedRetries, retries.Count(), "retry-go OnRetry should fire exactly twice per drop")
}

// TestFailureInjectionReachesWritePath is the negative control for
// TestRetryAbsorbsConnectionDrop: with retries disabled (MaxAttempts=1) the
// failing proxy must cause the write to fail. This guards against the
// failure injection being silently absorbed by background driver activity
// (heartbeats, topology probes) and never actually reaching the write path.
func TestFailureInjectionReachesWritePath(t *testing.T) {
	maxBackoff := configtype.NewDuration(50 * time.Millisecond)
	c, proxy, retries := newRetryReproClient(t, &spec.WriteRetryConfig{MaxAttempts: 1, MaxBackoff: &maxBackoff})

	proxy.dropNext(100) // far more than the single attempt should hit
	err := c.overwriteTableBatch(context.Background(), retryReproTable, []any{bson.M{"id": int64(1)}})
	require.Error(t, err, "expected failure without retry, but write succeeded")
	require.True(t, isRetryableWriteError(err), "expected a retryable network error, got: %v", err)
	require.Equal(t, 0, retries.Count(), "OnRetry must not fire when MaxAttempts=1")
}

// TestRetryGivesUpWhenAllAttemptsFail asserts that when the configured budget
// is exhausted before the proxy runs out of failures, the error propagates.
func TestRetryGivesUpWhenAllAttemptsFail(t *testing.T) {
	maxBackoff := configtype.NewDuration(50 * time.Millisecond)
	const maxAttempts = 3
	c, proxy, retries := newRetryReproClient(t, &spec.WriteRetryConfig{MaxAttempts: maxAttempts, MaxBackoff: &maxBackoff})

	// Drop far more connections than the retry budget can recover from, so
	// every op() invocation fails.
	proxy.dropNext(10_000)

	err := c.overwriteTableBatch(context.Background(), retryReproTable, []any{bson.M{"id": int64(1)}})
	require.Error(t, err)
	require.True(t, isRetryableWriteError(err), "expected a retryable network error, got: %v", err)
	// retry-go's OnRetry fires once per failing attempt -- including the
	// final attempt, even though no retry follows it (see retry-go v5
	// retry.go: r.onRetry(n, err) is called before the last-attempt check).
	require.Equal(t, maxAttempts, retries.Count(), "OnRetry should fire once per failing attempt")
}

package client

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/writers/mixedbatchwriter"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// fakeBatchClient stands in for *Client so the writer can be driven without a
// database. Calls are logged as begin/end pairs for the ordering assertions.
type fakeBatchClient struct {
	insertFn       func(context.Context, message.WriteInserts) error
	migrateFn      func(context.Context, message.WriteMigrateTables) error
	deleteStaleFn  func(context.Context, message.WriteDeleteStales) error
	deleteRecordFn func(context.Context, message.WriteDeleteRecords) error

	mu                sync.Mutex
	events            []string
	insertBatches     [][]string
	migrateSizes      []int
	deleteStaleSizes  []int
	deleteRecordSizes []int

	inFlight    atomic.Int64
	maxInFlight atomic.Int64
	started     atomic.Int64
	finished    atomic.Int64
}

var _ mixedbatchwriter.Client = (*fakeBatchClient)(nil)

func (c *fakeBatchClient) InsertBatch(ctx context.Context, messages message.WriteInserts) error {
	c.started.Add(1)
	c.observe(c.inFlight.Add(1))
	defer c.inFlight.Add(-1)
	defer c.finished.Add(1)
	c.log("insert:begin")

	var err error
	if c.insertFn != nil {
		err = c.insertFn(ctx, messages)
	}
	// Read on the way out: a reused backing array shows up here as the wrong rows.
	ids := recordIDs(messages)

	c.mu.Lock()
	c.insertBatches = append(c.insertBatches, ids)
	c.mu.Unlock()
	c.log("insert:end")
	return err
}

func (c *fakeBatchClient) MigrateTableBatch(ctx context.Context, messages message.WriteMigrateTables) error {
	c.log("migrate:begin")
	defer c.log("migrate:end")
	c.mu.Lock()
	c.migrateSizes = append(c.migrateSizes, len(messages))
	c.mu.Unlock()
	if c.migrateFn != nil {
		return c.migrateFn(ctx, messages)
	}
	return nil
}

func (c *fakeBatchClient) DeleteStaleBatch(ctx context.Context, messages message.WriteDeleteStales) error {
	c.log("delete-stale:begin")
	defer c.log("delete-stale:end")
	c.mu.Lock()
	c.deleteStaleSizes = append(c.deleteStaleSizes, len(messages))
	c.mu.Unlock()
	if c.deleteStaleFn != nil {
		return c.deleteStaleFn(ctx, messages)
	}
	return nil
}

func (c *fakeBatchClient) DeleteRecordsBatch(ctx context.Context, messages message.WriteDeleteRecords) error {
	c.log("delete-record:begin")
	defer c.log("delete-record:end")
	c.mu.Lock()
	c.deleteRecordSizes = append(c.deleteRecordSizes, len(messages))
	c.mu.Unlock()
	if c.deleteRecordFn != nil {
		return c.deleteRecordFn(ctx, messages)
	}
	return nil
}

func (c *fakeBatchClient) log(event string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.events = append(c.events, event)
}

func (c *fakeBatchClient) observe(inFlight int64) {
	for {
		peak := c.maxInFlight.Load()
		if inFlight <= peak || c.maxInFlight.CompareAndSwap(peak, inFlight) {
			return
		}
	}
}

func (c *fakeBatchClient) eventLog() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]string(nil), c.events...)
}

// batches returns the rows of every completed insert batch, in completion order.
func (c *fakeBatchClient) batches() [][]string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([][]string(nil), c.insertBatches...)
}

func recordIDs(messages message.WriteInserts) []string {
	var ids []string
	for _, msg := range messages {
		col := msg.Record.Column(0).(*array.String)
		for i := range int(msg.Record.NumRows()) {
			ids = append(ids, col.Value(i))
		}
	}
	return ids
}

func writerTestTable() *schema.Table {
	return &schema.Table{
		Name:    "concurrent_writer_unit",
		Columns: []schema.Column{{Name: "id", Type: arrow.BinaryTypes.String}},
	}
}

// insertMessages builds n single-row inserts with ids "id-000", "id-001", ...
func insertMessages(n int) []message.WriteMessage {
	table := writerTestTable()
	msgs := make([]message.WriteMessage, n)
	for i := range n {
		bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
		bldr.Field(0).(*array.StringBuilder).Append(fmt.Sprintf("id-%03d", i))
		msgs[i] = &message.WriteInsert{Record: bldr.NewRecordBatch()}
		bldr.Release()
	}
	return msgs
}

func newTestWriter(t *testing.T, c mixedbatchwriter.Client, batchSize, batchSizeBytes, concurrency int64) *concurrentWriter {
	t.Helper()
	logger := zerolog.New(zerolog.NewTestWriter(t)).Level(zerolog.WarnLevel)
	// The hour timeout keeps the ticker out of tests that are not about it.
	return newConcurrentWriter(c, logger, batchSize, batchSizeBytes, concurrency, time.Hour)
}

// writeMessages feeds msgs on its own goroutine, so a writer blocked on a full
// pool still makes progress.
func writeMessages(ctx context.Context, w *concurrentWriter, msgs []message.WriteMessage) error {
	ch := make(chan message.WriteMessage)
	go func() {
		defer close(ch)
		for _, msg := range msgs {
			select {
			case ch <- msg:
			case <-ctx.Done():
				return
			}
		}
	}()
	return w.Write(ctx, ch)
}

func TestConcurrentWriter_FlushesByRowCount(t *testing.T) {
	fake := &fakeBatchClient{}
	w := newTestWriter(t, fake, 3, 1<<30, 2)

	require.NoError(t, writeMessages(context.Background(), w, insertMessages(7)))

	require.ElementsMatch(t, [][]string{
		{"id-000", "id-001", "id-002"},
		{"id-003", "id-004", "id-005"},
		{"id-006"},
	}, fake.batches())
}

func TestConcurrentWriter_FlushesBySizeBytes(t *testing.T) {
	msgs := insertMessages(4)
	perRecord := recordSize(msgs[0].(*message.WriteInsert).Record)
	require.Positive(t, perRecord)

	fake := &fakeBatchClient{}
	// Row count is out of reach, so only the byte limit can split these.
	w := newTestWriter(t, fake, 1000, 2*perRecord, 2)

	require.NoError(t, writeMessages(context.Background(), w, msgs))

	require.ElementsMatch(t, [][]string{
		{"id-000", "id-001"},
		{"id-002", "id-003"},
	}, fake.batches())
}

// Without the clone in flushInserts a worker reads whichever rows are in the
// shared backing array by the time it looks.
func TestConcurrentWriter_HandsEachWorkerItsOwnRows(t *testing.T) {
	fake := &fakeBatchClient{insertFn: func(context.Context, message.WriteInserts) error {
		// Hold the batch open long enough for the writer to fill the next two.
		time.Sleep(50 * time.Millisecond)
		return nil
	}}
	w := newTestWriter(t, fake, 2, 1<<30, 4)

	require.NoError(t, writeMessages(context.Background(), w, insertMessages(6)))

	require.ElementsMatch(t, [][]string{
		{"id-000", "id-001"},
		{"id-002", "id-003"},
		{"id-004", "id-005"},
	}, fake.batches())
}

func TestConcurrentWriter_AppliesBatchesInParallel(t *testing.T) {
	const concurrency = 4
	release := make(chan struct{})
	fake := &fakeBatchClient{insertFn: func(ctx context.Context, _ message.WriteInserts) error {
		select {
		case <-release:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}}
	w := newTestWriter(t, fake, 1, 1<<30, concurrency)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	done := make(chan error, 1)
	go func() { done <- writeMessages(ctx, w, insertMessages(3*concurrency)) }()

	require.Eventually(t, func() bool { return fake.inFlight.Load() == concurrency }, 20*time.Second, time.Millisecond,
		"the pool should reach write_concurrency batches in flight at once")
	// Nothing has been released yet, so an overshoot would show up here.
	require.LessOrEqual(t, fake.maxInFlight.Load(), int64(concurrency))
	close(release)

	require.NoError(t, <-done)
	require.Equal(t, int64(concurrency), fake.maxInFlight.Load())
	require.Len(t, fake.batches(), 3*concurrency)
}

func TestConcurrentWriter_ConcurrencyOfOneStaysSerial(t *testing.T) {
	fake := &fakeBatchClient{insertFn: func(context.Context, message.WriteInserts) error {
		time.Sleep(5 * time.Millisecond)
		return nil
	}}
	w := newTestWriter(t, fake, 1, 1<<30, 1)

	require.NoError(t, writeMessages(context.Background(), w, insertMessages(5)))

	require.Equal(t, int64(1), fake.maxInFlight.Load(), "concurrency 1 must not overlap batches")
	require.Len(t, fake.batches(), 5)
}

// A pool sized at zero would block on its first acquire.
func TestConcurrentWriter_NonPositiveConcurrencyFallsBackToSerial(t *testing.T) {
	for _, concurrency := range []int64{0, -1} {
		t.Run(fmt.Sprint(concurrency), func(t *testing.T) {
			fake := &fakeBatchClient{}
			w := newTestWriter(t, fake, 1, 1<<30, concurrency)
			require.Equal(t, int64(1), w.concurrency)

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			require.NoError(t, writeMessages(ctx, w, insertMessages(3)))
			require.Len(t, fake.batches(), 3)
		})
	}
}

// The barrier that keeps the SDK's migrate -> insert -> delete ordering.
func TestConcurrentWriter_DrainsInsertsBeforeOtherMessageTypes(t *testing.T) {
	fake := &fakeBatchClient{insertFn: func(context.Context, message.WriteInserts) error {
		// Keep flushes in flight while the writer moves on to the next message.
		time.Sleep(20 * time.Millisecond)
		return nil
	}}
	w := newTestWriter(t, fake, 2, 1<<30, 4)

	msgs := make([]message.WriteMessage, 0, 13)
	msgs = append(msgs, &message.WriteMigrateTable{Table: writerTestTable()})
	msgs = append(msgs, insertMessages(6)...)
	msgs = append(msgs, &message.WriteDeleteStale{TableName: "concurrent_writer_unit", SourceName: testSourceName, SyncTime: time.Now()})
	msgs = append(msgs, insertMessages(4)...)
	msgs = append(msgs, &message.WriteDeleteRecord{DeleteRecord: message.DeleteRecord{TableName: "concurrent_writer_unit"}})

	require.NoError(t, writeMessages(context.Background(), w, msgs))

	open := 0
	for _, event := range fake.eventLog() {
		switch event {
		case "insert:begin":
			open++
		case "insert:end":
			open--
		default:
			require.Zero(t, open, "%s ran with %d insert batches still in flight", event, open)
		}
	}
	require.Len(t, fake.batches(), 5, "6 then 4 rows at batch_size 2")
	require.Equal(t, []int{1}, fake.migrateSizes)
	require.Equal(t, []int{1}, fake.deleteStaleSizes)
	require.Equal(t, []int{1}, fake.deleteRecordSizes)
}

func TestConcurrentWriter_FlushesEachRunOfMessages(t *testing.T) {
	fake := &fakeBatchClient{}
	w := newTestWriter(t, fake, 1000, 1<<30, 4)

	table := writerTestTable()
	inserts := insertMessages(2)
	msgs := []message.WriteMessage{
		&message.WriteMigrateTable{Table: table},
		inserts[0],
		&message.WriteMigrateTable{Table: table},
		inserts[1],
	}
	require.NoError(t, writeMessages(context.Background(), w, msgs))

	require.Equal(t, []string{
		"migrate:begin", "migrate:end",
		"insert:begin", "insert:end",
		"migrate:begin", "migrate:end",
		"insert:begin", "insert:end",
	}, fake.eventLog(), "a change of message type must flush what came before it")
}

func TestConcurrentWriter_FlushesOnBatchTimeout(t *testing.T) {
	fake := &fakeBatchClient{}
	logger := zerolog.New(zerolog.NewTestWriter(t)).Level(zerolog.WarnLevel)
	// Row and byte limits are out of reach: only the timeout can flush this.
	w := newConcurrentWriter(fake, logger, 1000, 1<<30, 2, 50*time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ch := make(chan message.WriteMessage)
	done := make(chan error, 1)
	go func() { done <- w.Write(ctx, ch) }()

	msgs := insertMessages(2)
	ch <- msgs[0]
	require.Eventually(t, func() bool { return len(fake.batches()) == 1 }, 20*time.Second, time.Millisecond,
		"a partial batch must be flushed once batch_timeout elapses")

	ch <- msgs[1]
	close(ch)
	require.NoError(t, <-done)

	require.ElementsMatch(t, [][]string{{"id-000"}, {"id-001"}}, fake.batches())
}

func TestConcurrentWriter_NoMessages(t *testing.T) {
	fake := &fakeBatchClient{}
	w := newTestWriter(t, fake, 10, 1<<30, 4)

	ch := make(chan message.WriteMessage)
	close(ch)
	require.NoError(t, w.Write(context.Background(), ch))
	require.Empty(t, fake.eventLog(), "an empty stream must not flush anything")
}

func TestConcurrentWriter_PropagatesErrors(t *testing.T) {
	wantErr := errors.New("boom")
	table := writerTestTable()

	for _, tt := range []struct {
		name   string
		client *fakeBatchClient
		msgs   []message.WriteMessage
	}{
		{
			name:   "insert",
			client: &fakeBatchClient{insertFn: func(context.Context, message.WriteInserts) error { return wantErr }},
			msgs:   insertMessages(4),
		},
		{
			name:   "migrate",
			client: &fakeBatchClient{migrateFn: func(context.Context, message.WriteMigrateTables) error { return wantErr }},
			msgs:   []message.WriteMessage{&message.WriteMigrateTable{Table: table}},
		},
		{
			name:   "delete stale",
			client: &fakeBatchClient{deleteStaleFn: func(context.Context, message.WriteDeleteStales) error { return wantErr }},
			msgs:   []message.WriteMessage{&message.WriteDeleteStale{TableName: table.Name, SyncTime: time.Now()}},
		},
		{
			name:   "delete record",
			client: &fakeBatchClient{deleteRecordFn: func(context.Context, message.WriteDeleteRecords) error { return wantErr }},
			msgs:   []message.WriteMessage{&message.WriteDeleteRecord{DeleteRecord: message.DeleteRecord{TableName: table.Name}}},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			w := newTestWriter(t, tt.client, 1, 1<<30, 4)
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			require.ErrorIs(t, writeMessages(ctx, w, tt.msgs), wantErr)
		})
	}
}

// A failed flush must not leave the rest of the pool writing behind the caller's
// back.
func TestConcurrentWriter_WaitsForInFlightBatchesOnError(t *testing.T) {
	wantErr := errors.New("boom")
	var calls atomic.Int64
	fake := &fakeBatchClient{insertFn: func(context.Context, message.WriteInserts) error {
		time.Sleep(100 * time.Millisecond)
		if calls.Add(1) == 1 {
			return wantErr
		}
		return nil
	}}
	w := newTestWriter(t, fake, 1, 1<<30, 4)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	require.ErrorIs(t, writeMessages(ctx, w, insertMessages(12)), wantErr)

	require.Equal(t, fake.started.Load(), fake.finished.Load(),
		"every batch the writer started must have finished before Write returned")
	require.Zero(t, fake.inFlight.Load())
}

func TestConcurrentWriter_ReturnsContextError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	started := make(chan struct{})
	var once sync.Once
	fake := &fakeBatchClient{insertFn: func(ctx context.Context, _ message.WriteInserts) error {
		once.Do(func() { close(started) })
		<-ctx.Done()
		return ctx.Err()
	}}
	// One slot, so the second batch cannot be submitted until the first is done.
	w := newTestWriter(t, fake, 1, 1<<30, 1)

	done := make(chan error, 1)
	go func() { done <- writeMessages(ctx, w, insertMessages(5)) }()

	<-started
	cancel()

	select {
	case err := <-done:
		require.ErrorIs(t, err, context.Canceled)
	case <-time.After(30 * time.Second):
		t.Fatal("Write did not return after the context was cancelled")
	}
	require.Zero(t, fake.inFlight.Load())
}

func TestAsyncFlusher_BoundsConcurrency(t *testing.T) {
	const concurrency = 3
	var inFlight, peak atomic.Int64
	release := make(chan struct{})
	flusher := newAsyncFlusher(concurrency)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Nothing is released until the pool is full, so peak is the real ceiling
	// rather than whatever the scheduler happened to overlap.
	go func() {
		for inFlight.Load() < concurrency {
			time.Sleep(time.Millisecond)
		}
		close(release)
	}()

	// Submitting from this goroutine mirrors the writer: submit and drain never
	// run concurrently.
	for range 10 {
		require.NoError(t, flusher.submit(ctx, func() error {
			n := inFlight.Add(1)
			defer inFlight.Add(-1)
			for {
				best := peak.Load()
				if n <= best || peak.CompareAndSwap(best, n) {
					break
				}
			}
			<-release
			return nil
		}))
	}
	require.NoError(t, flusher.drain())
	require.Equal(t, int64(concurrency), peak.Load(), "the pool must reach, and never exceed, its size")
}

func TestAsyncFlusher_KeepsFirstError(t *testing.T) {
	first, second := errors.New("first"), errors.New("second")
	ctx := context.Background()
	flusher := newAsyncFlusher(1)

	require.NoError(t, flusher.submit(ctx, func() error { return first }))
	require.ErrorIs(t, flusher.drain(), first)

	// A later failure does not replace it, and the next submit reports it instead
	// of starting more work.
	require.ErrorIs(t, flusher.submit(ctx, func() error { return second }), first)
	require.ErrorIs(t, flusher.drain(), first)
}

func TestAsyncFlusher_DrainWithoutWork(t *testing.T) {
	require.NoError(t, newAsyncFlusher(2).drain())
}

func TestAsyncFlusher_SubmitHonoursContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	flusher := newAsyncFlusher(1)

	blocked := make(chan struct{})
	require.NoError(t, flusher.submit(ctx, func() error {
		<-blocked
		return nil
	}))
	cancel()

	// The single slot is taken, so this can only end on the cancelled context.
	require.ErrorIs(t, flusher.submit(ctx, func() error { return nil }), context.Canceled)
	close(blocked)
	require.NoError(t, flusher.drain())
}

func TestRecordSize(t *testing.T) {
	table := writerTestTable()
	build := func(rows int) arrow.RecordBatch {
		bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
		defer bldr.Release()
		for i := range rows {
			bldr.Field(0).(*array.StringBuilder).Append(fmt.Sprintf("id-%03d", i))
		}
		return bldr.NewRecordBatch()
	}

	small, large := build(1), build(100)
	require.Positive(t, recordSize(small))
	require.Greater(t, recordSize(large), recordSize(small), "size must track the rows held")

	empty := build(0)
	require.GreaterOrEqual(t, recordSize(empty), int64(0))
}

func TestArrayDataSize(t *testing.T) {
	require.Zero(t, arrayDataSize(nil), "a column with no data contributes nothing")

	// A list keeps its values in a child, which a top-level walk would miss.
	bldr := array.NewListBuilder(memory.DefaultAllocator, arrow.BinaryTypes.String)
	defer bldr.Release()
	values := bldr.ValueBuilder().(*array.StringBuilder)
	for i := range 50 {
		bldr.Append(true)
		values.Append(fmt.Sprintf("value-%03d", i))
	}
	list := bldr.NewArray()
	defer list.Release()

	var topLevel int64
	for _, buf := range list.Data().Buffers() {
		if buf != nil {
			topLevel += int64(buf.Len())
		}
	}
	require.Greater(t, arrayDataSize(list.Data()), topLevel, "child buffers must be counted")
}

// --- Tests below here run against a live database. ---

// Has to leave room for write_concurrency, or batches queue on the pool instead
// of running side by side.
const testPoolMaxConns = 10

func concurrentTestConnection() string {
	separator := "?"
	if strings.Contains(getTestConnection(), "?") {
		separator = "&"
	}
	return fmt.Sprintf("%s%spool_max_conns=%d", getTestConnection(), separator, testPoolMaxConns)
}

func concurrentTestTable(t *testing.T) *schema.Table {
	t.Helper()
	return &schema.Table{
		Name: fmt.Sprintf("cq_concurrent_%d", time.Now().UnixNano()),
		Columns: []schema.Column{
			{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true, NotNull: true},
			{Name: "name", Type: arrow.BinaryTypes.String},
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
		},
	}
}

// What delete-stale keys on.
const testSourceName = "src"

func concurrentRecord(table *schema.Table, id, name string, syncTime time.Time) arrow.RecordBatch {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	defer bldr.Release()
	bldr.Field(0).(*array.StringBuilder).Append(id)
	bldr.Field(1).(*array.StringBuilder).Append(name)
	bldr.Field(2).(*array.StringBuilder).Append(testSourceName)
	bldr.Field(3).(*array.TimestampBuilder).Append(arrow.Timestamp(syncTime.UnixMicro()))
	return bldr.NewRecordBatch()
}

func newConcurrentPlugin(t *testing.T, s *spec.Spec) *plugin.Plugin {
	t.Helper()
	p := plugin.NewPlugin("postgresql", "development", New, plugin.WithJSONSchema(spec.JSONSchema))
	p.SetLogger(zerolog.New(zerolog.NewTestWriter(t)).Level(zerolog.WarnLevel))
	b, err := json.Marshal(s)
	require.NoError(t, err)
	require.NoError(t, p.Init(context.Background(), b, plugin.NewClientOptions{}))
	t.Cleanup(func() { _ = p.Close(context.Background()) })
	return p
}

func TestConcurrentWriter_EnabledByWriteConcurrency(t *testing.T) {
	for _, tt := range []struct {
		concurrency int64
		want        int64 // 0 means the SDK writer is used instead
	}{
		{concurrency: 0, want: 0},
		{concurrency: 1, want: 0},
		{concurrency: 2, want: 2},
		{concurrency: 8, want: 8},
	} {
		t.Run(fmt.Sprint(tt.concurrency), func(t *testing.T) {
			ctx := context.Background()
			b, err := json.Marshal(&spec.Spec{
				ConnectionString: getTestConnection(),
				WriteConcurrency: tt.concurrency,
			})
			require.NoError(t, err)

			c, err := New(ctx, zerolog.Nop(), b, plugin.NewClientOptions{})
			require.NoError(t, err)
			t.Cleanup(func() { _ = c.Close(ctx) })

			w := c.(*Client).concurrentWriter
			if tt.want == 0 {
				require.Nil(t, w, "write_concurrency %d must keep the serial SDK writer", tt.concurrency)
				return
			}
			require.NotNil(t, w)
			require.Equal(t, tt.want, w.concurrency)
		})
	}
}

// write_concurrency has to be invisible in the result: every row lands once, and
// a later sync of the same keys updates rather than duplicates them.
func TestConcurrentWriter_WritesEveryRowExactlyOnce(t *testing.T) {
	const rows = 1000
	ctx := context.Background()
	table := concurrentTestTable(t)
	p := newConcurrentPlugin(t, &spec.Spec{
		ConnectionString: concurrentTestConnection(),
		WriteConcurrency: 8,
		BatchSize:        25, // 40 batches over 8 connections
		RetryOnDeadlock:  5,
	})

	syncTime := time.Now().UTC().Truncate(time.Microsecond)
	ids := make([]string, rows)
	msgs := make([]message.WriteMessage, 0, rows+1)
	msgs = append(msgs, &message.WriteMigrateTable{Table: table})
	for i := range rows {
		ids[i] = uuid.New().String()
		msgs = append(msgs, &message.WriteInsert{Record: concurrentRecord(table, ids[i], "first", syncTime)})
	}
	require.NoError(t, p.WriteAll(ctx, msgs))

	got := readAll(ctx, t, p, table)
	require.Len(t, got, rows)
	for _, id := range ids {
		require.Equal(t, "first", got[id])
	}

	rewrite := make([]message.WriteMessage, 0, rows)
	for _, id := range ids {
		rewrite = append(rewrite, &message.WriteInsert{Record: concurrentRecord(table, id, "second", syncTime.Add(time.Second))})
	}
	require.NoError(t, p.WriteAll(ctx, rewrite))

	got = readAll(ctx, t, p, table)
	require.Len(t, got, rows, "concurrent batches must upsert, not duplicate")
	for _, id := range ids {
		require.Equal(t, "second", got[id])
	}
}

// The inserts carry the older sync time on purpose: a delete-stale that overtakes
// them leaves rows behind instead of clearing the table.
func TestConcurrentWriter_DeleteStaleWaitsForInserts(t *testing.T) {
	const rows = 600
	ctx := context.Background()
	table := concurrentTestTable(t)
	p := newConcurrentPlugin(t, &spec.Spec{
		ConnectionString: concurrentTestConnection(),
		WriteConcurrency: 8,
		BatchSize:        20,
		RetryOnDeadlock:  5,
	})

	stale := time.Now().UTC().Add(-time.Hour).Truncate(time.Microsecond)
	msgs := make([]message.WriteMessage, 0, rows+2)
	msgs = append(msgs, &message.WriteMigrateTable{Table: table})
	for range rows {
		msgs = append(msgs, &message.WriteInsert{Record: concurrentRecord(table, uuid.New().String(), "stale", stale)})
	}
	msgs = append(msgs, &message.WriteDeleteStale{
		TableName:  table.Name,
		SourceName: testSourceName,
		SyncTime:   stale.Add(time.Hour),
	})
	require.NoError(t, p.WriteAll(ctx, msgs))

	require.Empty(t, readAll(ctx, t, p, table),
		"every insert must have landed before the delete-stale ran")
}

func TestConcurrentWriter_KeepsRowsFromTheCurrentSync(t *testing.T) {
	const rows = 300
	ctx := context.Background()
	table := concurrentTestTable(t)
	p := newConcurrentPlugin(t, &spec.Spec{
		ConnectionString: concurrentTestConnection(),
		WriteConcurrency: 4,
		BatchSize:        20,
		RetryOnDeadlock:  5,
	})

	previous := time.Now().UTC().Add(-time.Hour).Truncate(time.Microsecond)
	old := make([]message.WriteMessage, 0, rows+1)
	old = append(old, &message.WriteMigrateTable{Table: table})
	for range rows {
		old = append(old, &message.WriteInsert{Record: concurrentRecord(table, uuid.New().String(), "old", previous)})
	}
	require.NoError(t, p.WriteAll(ctx, old))

	current := time.Now().UTC().Truncate(time.Microsecond)
	fresh := make([]message.WriteMessage, 0, rows+1)
	freshIDs := make([]string, rows)
	for i := range rows {
		freshIDs[i] = uuid.New().String()
		fresh = append(fresh, &message.WriteInsert{Record: concurrentRecord(table, freshIDs[i], "new", current)})
	}
	fresh = append(fresh, &message.WriteDeleteStale{TableName: table.Name, SourceName: testSourceName, SyncTime: current})
	require.NoError(t, p.WriteAll(ctx, fresh))

	got := readAll(ctx, t, p, table)
	require.Len(t, got, rows, "the previous sync's rows should be the only ones deleted")
	for _, id := range freshIDs {
		require.Equal(t, "new", got[id])
	}
}

// COPY upserts through a temporary staging table, which several connections now
// build at the same time.
func TestConcurrentWriter_WithCopyFrom(t *testing.T) {
	const rows = 800
	ctx := context.Background()
	table := concurrentTestTable(t)
	p := newConcurrentPlugin(t, &spec.Spec{
		ConnectionString: concurrentTestConnection(),
		WriteConcurrency: 4,
		BatchSize:        200, // above minCopyRows, so every batch takes the COPY path
		UseCopyFrom:      true,
		RetryOnDeadlock:  5,
	})

	syncTime := time.Now().UTC().Truncate(time.Microsecond)
	ids := make([]string, rows)
	msgs := make([]message.WriteMessage, 0, rows+1)
	msgs = append(msgs, &message.WriteMigrateTable{Table: table})
	for i := range rows {
		ids[i] = uuid.New().String()
		msgs = append(msgs, &message.WriteInsert{Record: concurrentRecord(table, ids[i], "copied", syncTime)})
	}
	require.NoError(t, p.WriteAll(ctx, msgs))

	got := readAll(ctx, t, p, table)
	require.Len(t, got, rows)
	for _, id := range ids {
		require.Equal(t, "copied", got[id])
	}
}

// Concurrent upserts over the same keys deadlock in PostgreSQL; retry_on_deadlock
// is what keeps the sync alive, which is why the spec docs pair the two.
func TestConcurrentWriter_OverlappingKeysWithDeadlockRetry(t *testing.T) {
	const (
		keys   = 50
		passes = 20
	)
	ctx := context.Background()
	isPostgres, err := isPostgresDB(ctx)
	require.NoError(t, err)
	if !isPostgres {
		t.Skip("deadlock retry is keyed on the PostgreSQL 40P01 error code")
	}

	table := concurrentTestTable(t)
	p := newConcurrentPlugin(t, &spec.Spec{
		// The default deadlock_timeout costs a second of lock waiting per deadlock.
		ConnectionString: concurrentTestConnection() + "&options=-c%20deadlock_timeout%3D20ms",
		WriteConcurrency: 8,
		BatchSize:        keys,
		RetryOnDeadlock:  10,
	})

	ids := make([]string, keys)
	for i := range keys {
		ids[i] = uuid.New().String()
	}

	syncTime := time.Now().UTC().Truncate(time.Microsecond)
	msgs := make([]message.WriteMessage, 0, keys*passes+1)
	msgs = append(msgs, &message.WriteMigrateTable{Table: table})
	for pass := range passes {
		// Each pass touches the same keys in a different order, so overlapping
		// batches take their row locks in conflicting orders.
		rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })
		for _, id := range ids {
			msgs = append(msgs, &message.WriteInsert{
				Record: concurrentRecord(table, id, fmt.Sprintf("pass-%02d", pass), syncTime),
			})
		}
	}
	require.NoError(t, p.WriteAll(ctx, msgs))

	// Which pass wins is undefined once batches overlap.
	require.Len(t, readAll(ctx, t, p, table), keys)
}

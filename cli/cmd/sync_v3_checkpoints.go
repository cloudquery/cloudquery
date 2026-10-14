package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"slices"
	"strings"
	gosync "sync"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/cloudquery/cloudquery/cli/v6/internal/transformerpipeline"
)

// Checkpoints closer together than this are persisted together, to avoid flushing destinations too often.
// A variable so that tests can shorten it.
var checkpointInterval = 30 * time.Second

// errCheckpointDue is returned by the checkpointer's receiver when the pending state is due while the source is idle.
var errCheckpointDue = errors.New("state checkpoint due")

// State is written to the backend in records of about this size, well within the 200 MiB gRPC message limit.
const stateRecordMaxBytes = 64 * 1024 * 1024

const (
	stateKeyColumn     = "key"
	stateValueColumn   = "value"
	stateVersionColumn = "version"
)

type writeClient interface {
	Send(*plugin.Write_Request) error
	CloseAndRecv() (*plugin.Write_Response, error)
}

// flushingWriteClient is a WriteStream client. Like safeWriteClient, it is safe for concurrent use.
type flushingWriteClient struct {
	name        string
	client      grpc.BidiStreamingClient[plugin.Write_Request, plugin.Write_Response]
	mu          gosync.Mutex
	lastFlushID uint64
}

func newFlushingWriteClient(name string, client grpc.BidiStreamingClient[plugin.Write_Request, plugin.Write_Response]) *flushingWriteClient {
	return &flushingWriteClient{name: name, client: client}
}

func (s *flushingWriteClient) Send(req *plugin.Write_Request) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.client.Send(req)
}

// flushFailedError is a destination's report that it could not commit some messages sent before a flush.
type flushFailedError struct {
	reason string
}

func (e *flushFailedError) Error() string { return e.reason }

// Flush returns once the destination commits every message sent before it. If the destination could not commit some
// of them, it returns a *flushFailedError, and the stream stays usable.
func (s *flushingWriteClient) Flush() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastFlushID++
	err := s.client.Send(&plugin.Write_Request{
		Message: &plugin.Write_Request_Flush{Flush: &plugin.Write_MessageFlush{Id: s.lastFlushID}},
	})
	// io.EOF means the stream ended, and Recv returns the reason
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	res, err := s.client.Recv()
	if err != nil {
		return err
	}
	if res.FlushId != s.lastFlushID {
		return fmt.Errorf("destination acknowledged flush %d, expected %d", res.FlushId, s.lastFlushID)
	}
	if res.FlushError != "" {
		return &flushFailedError{reason: res.FlushError}
	}
	return nil
}

func (s *flushingWriteClient) CloseAndRecv() (*plugin.Write_Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.client.CloseSend(); err != nil {
		return nil, err
	}
	for {
		if _, err := s.client.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				return &plugin.Write_Response{}, nil
			}
			return nil, err
		}
	}
}

type stateRow struct {
	value   string
	version uint64
}

// checkpointer persists the state from source checkpoints once every destination commits the records synced before them.
type checkpointer struct {
	destinations []*flushingWriteClient
	pipelines    []*transformerpipeline.TransformerPipeline
	backend      plugin.PluginClient
	// the backend supports WriteStream, so the state is committed when persisted
	backendFlushes bool
	backendStream  *flushingWriteClient
	stateTable     string
	// approximate size of each state record written to the backend
	maxRecordBytes int

	schema     *arrow.Schema
	pending    map[string]stateRow
	persisted  map[string]uint64
	lastCommit time.Time
	received   bool
	// destinations that failed to commit records, by name. The state then summarizes lost records, so the sync saves no
	// more of it.
	failed map[string]error
	// set when the state backend failed to commit state. The sync then saves no more state, as it does for destinations.
	backendFailed error
}

func newCheckpointer(destinations []*flushingWriteClient, pipelines []*transformerpipeline.TransformerPipeline, backend plugin.PluginClient, backendFlushes bool, stateTable string) *checkpointer {
	return &checkpointer{
		destinations:   destinations,
		pipelines:      pipelines,
		backend:        backend,
		backendFlushes: backendFlushes,
		stateTable:     stateTable,
		maxRecordBytes: stateRecordMaxBytes,
		pending:        make(map[string]stateRow),
		persisted:      make(map[string]uint64),
		failed:         make(map[string]error),
		lastCommit:     time.Now(),
	}
}

// add queues the state rows of a checkpoint, keeping the newest version of each key.
func (c *checkpointer) add(recordBytes []byte) error {
	record, err := plugin.NewRecordFromBytes(recordBytes)
	if err != nil {
		return fmt.Errorf("failed to decode checkpoint: %w", err)
	}
	tableName, _ := record.Schema().Metadata().GetValue(schema.MetadataTableName)
	if tableName != c.stateTable {
		return fmt.Errorf("checkpoint is for table %q, expected the state table %q", tableName, c.stateTable)
	}
	keys, err := requiredStateColumn[*array.String](record, stateKeyColumn, arrow.BinaryTypes.String)
	if err != nil {
		return err
	}
	values, err := requiredStateColumn[*array.String](record, stateValueColumn, arrow.BinaryTypes.String)
	if err != nil {
		return err
	}
	versioned := record.Schema().HasField(stateVersionColumn)
	var versions *array.Uint64
	if versioned {
		if versions, err = requiredStateColumn[*array.Uint64](record, stateVersionColumn, arrow.PrimitiveTypes.Uint64); err != nil {
			return err
		}
	}
	if !c.saving() {
		return nil
	}
	if !c.received {
		c.received = true
		log.Info().Str("table", c.stateTable).Msg("Source sends state checkpoints: state is persisted after destinations commit the records synced before it")
	}
	c.schema = record.Schema()
	for i := range int(record.NumRows()) {
		row := stateRow{value: values.Value(i)}
		key := keys.Value(i)
		if versioned {
			row.version = versions.Value(i)
			if pending, ok := c.pending[key]; ok && pending.version > row.version {
				continue
			}
			if persisted, ok := c.persisted[key]; ok && persisted >= row.version {
				continue
			}
		}
		c.pending[key] = row
	}
	return nil
}

func requiredStateColumn[T arrow.Array](record arrow.RecordBatch, name string, typ arrow.DataType) (T, error) {
	var zero T
	indices := record.Schema().FieldIndices(name)
	switch len(indices) {
	case 0:
		return zero, fmt.Errorf("checkpoint has no %q column", name)
	case 1:
	default:
		return zero, fmt.Errorf("checkpoint has %d %q columns", len(indices), name)
	}
	col, ok := record.Column(indices[0]).(T)
	if !ok || !arrow.TypeEqual(col.DataType(), typ) {
		return zero, fmt.Errorf("checkpoint column %q is %s, expected %s", name, record.Column(indices[0]).DataType(), typ)
	}
	return col, nil
}

// due reports whether the pending state should be persisted now.
func (c *checkpointer) due() bool {
	return c.saving() && len(c.pending) > 0 && time.Since(c.lastCommit) >= checkpointInterval
}

// saving reports whether the sync still saves state, which it stops once a destination or the state backend fails to
// commit what it was sent.
func (c *checkpointer) saving() bool {
	return len(c.failed) == 0 && c.backendFailed == nil
}

// failure returns why the sync saves no more state, if a destination or the state backend failed to commit.
func (c *checkpointer) failure() error {
	errs := make([]error, 0, len(c.failed)+1)
	for _, name := range slices.Sorted(maps.Keys(c.failed)) {
		errs = append(errs, fmt.Errorf("destination %s: %w", name, c.failed[name]))
	}
	if c.backendFailed != nil {
		errs = append(errs, fmt.Errorf("state backend: %w", c.backendFailed))
	}
	return errors.Join(errs...)
}

// stateNotSavedError is returned when a sync finishes, but saved no state since a destination or the state backend
// failed to commit.
type stateNotSavedError struct {
	err error
	// whether the source sent checkpoints. A source that does not saves its state itself, which may then be past the
	// records that were lost.
	checkpointed bool
}

func (e *stateNotSavedError) Error() string {
	if e.checkpointed {
		return fmt.Sprintf("sync finished without saving its state, as it failed to commit: %v", e.err)
	}
	return fmt.Sprintf("sync finished, but failed to commit, and the source sent no state checkpoints, so state it saved itself may be past the lost records: %v", e.err)
}

func (e *stateNotSavedError) Unwrap() error { return e.err }

// failedDestination reports whether the destination failed to commit records.
func (c *checkpointer) failedDestination(name string) bool {
	_, ok := c.failed[name]
	return ok
}

type syncResult struct {
	response *plugin.Sync_Response
	err      error
}

// receiver wraps recv to return errCheckpointDue when the pending state falls due while the source sends nothing,
// so the caller commits between the messages it processes. recv runs on its own goroutine.
func (c *checkpointer) receiver(ctx context.Context, recv func() (*plugin.Sync_Response, error)) func() (*plugin.Sync_Response, error) {
	results := make(chan syncResult)
	go func() {
		for {
			response, err := recv()
			select {
			case results <- syncResult{response: response, err: err}:
			case <-ctx.Done():
				return
			}
			if err != nil {
				return
			}
		}
	}()
	return func() (*plugin.Sync_Response, error) {
		result, received, err := c.waitForResult(ctx, results)
		if err != nil {
			return nil, err
		}
		if !received {
			return nil, errCheckpointDue
		}
		return result.response, result.err
	}
}

// waitForResult returns the next result, or received=false once the pending state falls due.
func (c *checkpointer) waitForResult(ctx context.Context, results <-chan syncResult) (result syncResult, received bool, err error) {
	var due <-chan time.Time
	if len(c.pending) > 0 {
		timer := time.NewTimer(time.Until(c.lastCommit.Add(checkpointInterval)))
		defer timer.Stop()
		due = timer.C
	}
	select {
	case result := <-results:
		return result, true, nil
	case <-due:
		return syncResult{}, false, nil
	case <-ctx.Done():
		return syncResult{}, false, ctx.Err()
	}
}

// commit persists the pending state after every destination commits the records sent so far.
func (c *checkpointer) commit(ctx context.Context) error {
	if len(c.pending) == 0 || !c.saving() {
		return nil
	}
	return c.flushAndPersist(ctx)
}

// finish commits at the end of the sync. It flushes every destination that has not failed yet, even without pending
// state, so that the sync learns about every destination that failed to commit records.
func (c *checkpointer) finish(ctx context.Context) error {
	return c.flushAndPersist(ctx)
}

func (c *checkpointer) flushAndPersist(ctx context.Context) error {
	for _, pipeline := range c.pipelines {
		if err := pipeline.WaitDelivered(ctx); err != nil {
			return fmt.Errorf("failed to wait for records to reach destinations: %w", err)
		}
	}
	if err := c.flushDestinations(); err != nil {
		return fmt.Errorf("failed to flush destinations: %w", err)
	}
	if !c.saving() {
		clear(c.pending)
		return nil
	}
	if len(c.pending) == 0 {
		return nil
	}
	err := c.persist(ctx)
	var failed *flushFailedError
	if errors.As(err, &failed) {
		c.backendFailed = failed
		clear(c.pending)
		log.Error().Err(c.failure()).Msg("The state backend failed to write state, so the sync saves no more state, and fails once it finishes")
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to persist state: %w", err)
	}
	for key, row := range c.pending {
		c.persisted[key] = row.version
	}
	clear(c.pending)
	c.lastCommit = time.Now()
	return nil
}

// flushDestinations flushes every destination that has not failed yet, and records the ones that failed to commit
// records.
func (c *checkpointer) flushDestinations() error {
	results := make([]error, len(c.destinations))
	var wg gosync.WaitGroup
	for i, destination := range c.destinations {
		if !c.failedDestination(destination.name) {
			wg.Go(func() { results[i] = destination.Flush() })
		}
	}
	wg.Wait()

	var (
		errs        []error
		newFailures bool
	)
	for i, err := range results {
		name := c.destinations[i].name
		var failed *flushFailedError
		switch {
		case errors.As(err, &failed):
			c.failed[name] = failed
			newFailures = true
		case err != nil:
			errs = append(errs, fmt.Errorf("destination %s: %w", name, err))
		}
	}
	if len(errs) > 0 {
		// failures to commit records are part of why the sync fails
		return errors.Join(append(errs, c.failure())...)
	}
	if newFailures {
		log.Error().Err(c.failure()).Msg("Destinations failed to write records, so the sync saves no more state, and fails once it finishes")
	}
	return nil
}

func (c *checkpointer) persist(ctx context.Context) error {
	var (
		send   func(*plugin.Write_Request) error
		finish func() error
	)
	if c.backendFlushes {
		if c.backendStream == nil {
			stream, err := c.backend.WriteStream(ctx)
			if err != nil {
				return err
			}
			c.backendStream = newFlushingWriteClient("state backend", stream)
		}
		send, finish = c.backendStream.Send, c.backendStream.Flush
	} else {
		stream, err := c.backend.Write(ctx)
		if err != nil {
			return err
		}
		send = stream.Send
		finish = func() error {
			_, err := stream.CloseAndRecv()
			return err
		}
	}

	// io.EOF means the stream ended, and finish returns the reason
	err := c.eachPendingRecord(func(recordBytes []byte) error {
		return send(&plugin.Write_Request{Message: &plugin.Write_Request_Insert{Insert: &plugin.Write_MessageInsert{Record: recordBytes}}})
	})
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return finish()
}

// eachPendingRecord passes the pending state to fn in records of about maxRecordBytes each, building each record only
// when the previous one is handled.
func (c *checkpointer) eachPendingRecord(fn func([]byte) error) error {
	keys := make([]string, 0, len(c.pending))
	for key := range c.pending {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	for start := 0; start < len(keys); {
		end, size := start, 0
		for end < len(keys) && (end == start || size+c.rowBytes(keys[end]) <= c.maxRecordBytes) {
			size += c.rowBytes(keys[end])
			end++
		}
		record, err := c.stateRecord(keys[start:end])
		if err != nil {
			return err
		}
		if err := fn(record); err != nil {
			return err
		}
		start = end
	}
	return nil
}

// rowBytes approximates the size of a state row in a record, with a null in every column that is not a state column.
func (c *checkpointer) rowBytes(key string) int {
	const versionAndOffsets, perColumn = 16, 8
	return len(key) + len(c.pending[key].value) + versionAndOffsets + perColumn*len(c.schema.Fields())
}

func (c *checkpointer) stateRecord(keys []string) ([]byte, error) {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, c.schema)
	defer bldr.Release()
	for i, field := range c.schema.Fields() {
		for _, key := range keys {
			row := c.pending[key]
			switch field.Name {
			case stateKeyColumn:
				bldr.Field(i).(*array.StringBuilder).Append(key)
			case stateValueColumn:
				bldr.Field(i).(*array.StringBuilder).Append(row.value)
			case stateVersionColumn:
				bldr.Field(i).(*array.Uint64Builder).Append(row.version)
			default:
				bldr.Field(i).AppendNull()
			}
		}
	}
	record := bldr.NewRecordBatch()
	defer record.Release()
	return plugin.RecordToBytes(record)
}

// close ends the stream to the backend.
func (c *checkpointer) close() error {
	if c.backendStream == nil {
		return nil
	}
	_, err := c.backendStream.CloseAndRecv()
	return err
}

// checkpointOptions decide whether a sync uses state checkpoints, and where the state is written.
type checkpointOptions struct {
	// backend_options.connection, before variables are replaced
	backendConnection       string
	destinationNames        []string
	destinations            []plugin.PluginClient
	destinationsFlush       bool
	destinationTransformers bool
	// backend plugin that is not a destination of the sync, if any
	backend      plugin.PluginClient
	backendFlush bool
}

// stateWriter returns the plugin that persists the state, and whether it supports flush, or why checkpoints are off.
func (o checkpointOptions) stateWriter() (writer plugin.PluginClient, flushes bool, disabledReason string) {
	switch {
	case !o.destinationsFlush:
		return nil, false, "not all destinations support flushing"
	case o.destinationTransformers:
		return nil, false, "destinations use transformers"
	}
	for i, name := range o.destinationNames {
		if strings.Contains(o.backendConnection, "@@plugins."+name+".") {
			return o.destinations[i], true, ""
		}
	}
	if o.backend != nil {
		return o.backend, o.backendFlush, ""
	}
	return nil, false, "the state backend is not a plugin of this sync"
}

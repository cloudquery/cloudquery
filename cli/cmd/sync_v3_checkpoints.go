package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
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
	"golang.org/x/sync/errgroup"
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
	// what the destination failed to commit so far. The sync fails because of it, so if the destination also ends the
	// stream, as writers that stop on a failed write do, the client drops what the sync still sends.
	failed *flushFailedError
	ended  bool
}

func newFlushingWriteClient(name string, client grpc.BidiStreamingClient[plugin.Write_Request, plugin.Write_Response]) *flushingWriteClient {
	return &flushingWriteClient{name: name, client: client}
}

func (s *flushingWriteClient) Send(req *plugin.Write_Request) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.ended {
		return nil
	}
	err := s.client.Send(req)
	if err != nil && s.failed != nil {
		s.endAfterFailure(err)
		return nil
	}
	return err
}

// endAfterFailure handles the end of the stream after the destination reported a failure. Records of any table may
// then have been lost, which it adds to the failure the sync already fails with.
func (s *flushingWriteClient) endAfterFailure(err error) {
	if errors.Is(err, io.EOF) {
		// io.EOF from Send means the stream ended, and Recv returns the reason
		if _, recvErr := s.client.Recv(); recvErr != nil {
			err = recvErr
		}
	}
	s.ended = true
	s.failed.add(fmt.Sprintf("write stream ended: %v", err), nil)
	log.Warn().Err(err).Str("destination", s.name).Msg("Write stream of the destination ended after it reported a failure, so it gets no more records in this sync")
}

// flushFailedError is what a destination reported it could not commit.
type flushFailedError struct {
	reasons []string
	// the tables whose messages may have been lost, unless anyTable is set
	tables   []string
	anyTable bool
}

func newFlushFailedError(reason string, tables []string) *flushFailedError {
	return &flushFailedError{reasons: []string{reason}, tables: slices.Clone(tables), anyTable: len(tables) == 0}
}

func (e *flushFailedError) Error() string { return strings.Join(e.reasons, "; ") }

// add records a later failure. Its tables add to the earlier ones, unless it names none, as messages of any table may
// then have been lost.
func (e *flushFailedError) add(reason string, tables []string) {
	if !slices.Contains(e.reasons, reason) {
		e.reasons = append(e.reasons, reason)
	}
	if len(tables) == 0 {
		e.anyTable = true
	}
	for _, table := range tables {
		if !slices.Contains(e.tables, table) {
			e.tables = append(e.tables, table)
		}
	}
}

func (e *flushFailedError) clone() *flushFailedError {
	return &flushFailedError{reasons: slices.Clone(e.reasons), tables: slices.Clone(e.tables), anyTable: e.anyTable}
}

// Flush returns once the destination has tried to commit every message sent before it. If the destination failed to
// commit some messages, at this flush or an earlier one, it returns a *flushFailedError with every failure so far.
func (s *flushingWriteClient) Flush() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.ended {
		return s.failed.clone()
	}
	s.lastFlushID++
	err := s.client.Send(&plugin.Write_Request{
		Message: &plugin.Write_Request_Flush{Flush: &plugin.Write_MessageFlush{Id: s.lastFlushID}},
	})
	// io.EOF means the stream ended, and Recv returns the reason
	if err != nil && !errors.Is(err, io.EOF) {
		return s.streamError(err)
	}
	res, err := s.client.Recv()
	if err != nil {
		return s.streamError(err)
	}
	if res.FlushId != s.lastFlushID {
		return fmt.Errorf("destination acknowledged flush %d, expected %d", res.FlushId, s.lastFlushID)
	}
	if res.FlushError != "" {
		if s.failed == nil {
			s.failed = newFlushFailedError(res.FlushError, res.FailedTables)
		} else {
			s.failed.add(res.FlushError, res.FailedTables)
		}
	}
	if s.failed == nil {
		return nil
	}
	return s.failed.clone()
}

// streamError returns err, or the failure the destination reported before it, as the stream may have ended because of
// that failure.
func (s *flushingWriteClient) streamError(err error) error {
	if s.failed == nil {
		return err
	}
	s.endAfterFailure(err)
	return s.failed.clone()
}

// failure returns what the destination failed to commit so far, or nil.
func (s *flushingWriteClient) failure() *flushFailedError {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.failed == nil {
		return nil
	}
	return s.failed.clone()
}

func (s *flushingWriteClient) CloseAndRecv() (*plugin.Write_Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.ended {
		return &plugin.Write_Response{}, nil
	}
	err := s.client.CloseSend()
	for err == nil {
		_, err = s.client.Recv()
	}
	switch {
	case errors.Is(err, io.EOF):
		return &plugin.Write_Response{}, nil
	case s.failed != nil:
		s.endAfterFailure(err)
		return &plugin.Write_Response{}, nil
	default:
		return nil, err
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
	// names of the destinations that failed to commit records. The state then summarizes lost records, so the sync saves
	// no more of it.
	failed map[string]bool
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
		failed:         make(map[string]bool),
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
	for _, destination := range c.destinations {
		if failed := destination.failure(); failed != nil {
			errs = append(errs, fmt.Errorf("destination %s: %w", destination.name, failed))
		}
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
	return fmt.Sprintf("sync finished, but failed to commit, and the source sent no state checkpoints: if it saves state itself, that state may be past the lost records: %v", e.err)
}

// waitPreferringGroup runs receive in eg, and returns the first error of eg's other goroutines before receive's: a
// failing pipeline cancels the sync's context, which then fails receive too, and the pipeline's error explains why.
func waitPreferringGroup(eg *errgroup.Group, receive func() error) error {
	var receiveErr error
	eg.Go(func() error {
		receiveErr = receive()
		return nil
	})
	if err := eg.Wait(); err != nil {
		return err
	}
	return receiveErr
}

func (e *stateNotSavedError) Unwrap() error { return e.err }

// failedDestination reports whether the destination failed to commit records.
func (c *checkpointer) failedDestination(name string) bool {
	return c.failed[name]
}

// withoutFailedTables returns the tables of the relation trees the destination committed every record of, out of
// tables, so that delete-stale neither deletes the rows of earlier syncs whose replacements failed to write, nor leaves
// the rows of a tree inconsistent. parents maps each relation to its parent table.
func (c *checkpointer) withoutFailedTables(name string, tables map[string]bool, parents map[string]string) map[string]bool {
	var failed *flushFailedError
	for _, destination := range c.destinations {
		if destination.name == name {
			failed = destination.failure()
		}
	}
	if failed == nil {
		return tables
	}
	if failed.anyTable {
		log.Warn().Str("destination", name).Msg("Skipping delete-stale, as the destination failed to write records")
		return nil
	}
	failedTrees := make(map[string]bool, len(failed.tables))
	for _, table := range failed.tables {
		failedTrees[rootTable(table, parents)] = true
	}
	kept := make(map[string]bool, len(tables))
	var skipped []string
	for table, ok := range tables {
		if failedTrees[rootTable(table, parents)] {
			skipped = append(skipped, table)
			continue
		}
		kept[table] = ok
	}
	if len(skipped) > 0 {
		slices.Sort(skipped)
		log.Warn().Str("destination", name).Strs("tables", skipped).Msg("Skipping delete-stale for the relation trees of tables the destination failed to write records to")
	}
	return kept
}

// rootTable returns the top-level table of the relation tree table is in.
func rootTable(table string, parents map[string]string) string {
	for parents[table] != "" {
		table = parents[table]
	}
	return table
}

// withUnsavedState adds to err that the sync saved no state, if a destination or the state backend failed to commit
// before err happened. A sync that finished returns a *stateNotSavedError of its own.
func (c *checkpointer) withUnsavedState(err error) error {
	if err == nil || errors.As(err, new(*stateNotSavedError)) {
		return err
	}
	if failure := c.failure(); failure != nil {
		return errors.Join(err, &stateNotSavedError{err: failure, checkpointed: c.received})
	}
	return err
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

// finish commits at the end of the sync. It flushes every destination, even without pending state or after a failure,
// so that the sync learns about every table each destination failed to commit records of.
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

// flushDestinations flushes every destination, and records the ones that failed to commit records. A commit only
// flushes before any failure, and the final flush asks failed destinations again, to learn about their later failures.
func (c *checkpointer) flushDestinations() error {
	results := make([]error, len(c.destinations))
	var wg gosync.WaitGroup
	for i, destination := range c.destinations {
		wg.Go(func() { results[i] = destination.Flush() })
	}
	wg.Wait()

	var (
		errs        []error
		newFailures bool
	)
	for i, err := range results {
		name := c.destinations[i].name
		switch {
		case errors.As(err, new(*flushFailedError)):
			newFailures = newFailures || !c.failed[name]
			c.failed[name] = true
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

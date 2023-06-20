package client

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/util"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

const (
	defaultBatchTimeoutSeconds = 20
	defaultBatchSize           = 10000
	defaultBatchSizeBytes      = 5 * 1024 * 1024 // 5 MiB
)

type StreamingBatchWriterClient interface {
	MigrateTables(context.Context, []*message.MigrateTable) error
	DeleteStale(context.Context, []*message.DeleteStale) error

	OpenTable(ctx context.Context, sourceName string, table *schema.Table, syncTime time.Time) (any, error)
	WriteTableStream(ctx context.Context, handle any, upsert bool, msgs []*message.Insert) error
	CloseTable(ctx context.Context, handle any) error
}

type StreamingBatchWriter struct {
	client           StreamingBatchWriterClient
	workers          map[string]*worker
	workersLock      *sync.RWMutex
	workersWaitGroup *sync.WaitGroup

	migrateTableLock     *sync.Mutex
	migrateTableMessages []*message.MigrateTable
	deleteStaleLock      *sync.Mutex
	deleteStaleMessages  []*message.DeleteStale

	logger         zerolog.Logger
	batchTimeout   time.Duration
	batchSize      int
	batchSizeBytes int
}

type Option func(*StreamingBatchWriter)

func WithLogger(logger zerolog.Logger) Option {
	return func(p *StreamingBatchWriter) {
		p.logger = logger
	}
}

func WithBatchTimeout(timeout time.Duration) Option {
	return func(p *StreamingBatchWriter) {
		p.batchTimeout = timeout
	}
}

func WithBatchSize(size int) Option {
	return func(p *StreamingBatchWriter) {
		p.batchSize = size
	}
}

func WithBatchSizeBytes(size int) Option {
	return func(p *StreamingBatchWriter) {
		p.batchSizeBytes = size
	}
}

type worker struct {
	count int
	ch    chan *message.Insert
	flush chan chan bool
}

func NewStreamingBatchWriter(client StreamingBatchWriterClient, opts ...Option) (*StreamingBatchWriter, error) {
	c := &StreamingBatchWriter{
		client:           client,
		workers:          make(map[string]*worker),
		workersLock:      &sync.RWMutex{},
		workersWaitGroup: &sync.WaitGroup{},
		migrateTableLock: &sync.Mutex{},
		deleteStaleLock:  &sync.Mutex{},
		logger:           zerolog.Nop(),
		batchTimeout:     defaultBatchTimeoutSeconds * time.Second,
		batchSize:        defaultBatchSize,
		batchSizeBytes:   defaultBatchSizeBytes,
	}
	for _, opt := range opts {
		opt(c)
	}
	c.migrateTableMessages = make([]*message.MigrateTable, 0, c.batchSize)
	c.deleteStaleMessages = make([]*message.DeleteStale, 0, c.batchSize)
	return c, nil
}

func (w *StreamingBatchWriter) Flush(ctx context.Context) error {
	w.workersLock.RLock()
	for _, worker := range w.workers {
		done := make(chan bool)
		worker.flush <- done
		<-done
	}
	w.workersLock.RUnlock()

	if err := w.flushMigrateTables(ctx); err != nil {
		return err
	}

	return w.flushDeleteStaleTables(ctx)
}

func (w *StreamingBatchWriter) Close(_ context.Context) error {
	w.workersLock.Lock()
	defer w.workersLock.Unlock()
	for _, w := range w.workers {
		close(w.ch)
	}
	w.workersWaitGroup.Wait()

	return nil
}

func (w *StreamingBatchWriter) worker(ctx context.Context, sourceName, tableName string, ch <-chan *message.Insert, flush <-chan chan bool) {
	sizeBytes := int64(0)
	resources := make([]*message.Insert, 0)
	upsertBatch := false

	initDone := false
	var handle any

	doInit := func(r arrow.Record) {
		if initDone {
			return
		}
		initDone = true

		syncTime := w.getSyncTime(r)
		if syncTime.IsZero() {
			syncTime = time.Now()
		}

		table, err := schema.NewTableFromArrowSchema(r.Schema())
		if err != nil {
			panic(err)
		}

		handle, err = w.client.OpenTable(ctx, sourceName, table, syncTime)
		if err != nil {
			panic(err)
		}
	}

	for {
		select {
		case r, ok := <-ch:
			if !ok {
				if len(resources) > 0 {
					w.flush(ctx, handle, tableName, upsertBatch, resources)
				}

				if initDone {
					if err := w.client.CloseTable(ctx, handle); err != nil {
						panic(err)
					}
				}
				return
			}

			doInit(r.Record)
			if upsertBatch != r.Upsert {
				w.flush(ctx, handle, tableName, upsertBatch, resources)
				resources, upsertBatch = resources[:0], r.Upsert
				resources = append(resources, r)
				sizeBytes = util.TotalRecordSize(r.Record)
			} else {
				resources = append(resources, r)
				sizeBytes += util.TotalRecordSize(r.Record)
			}
			if len(resources) >= w.batchSize || sizeBytes+util.TotalRecordSize(r.Record) >= int64(w.batchSizeBytes) {
				w.flush(ctx, handle, tableName, upsertBatch, resources)
				resources, sizeBytes = resources[:0], 0
			}
		case <-time.After(w.batchTimeout):
			if len(resources) > 0 {
				w.flush(ctx, handle, tableName, upsertBatch, resources)
				resources, sizeBytes = resources[:0], 0
			}
		case done := <-flush:
			if len(resources) > 0 {
				w.flush(ctx, handle, tableName, upsertBatch, resources)
				resources, sizeBytes = resources[:0], 0
			}
			done <- true
		}
	}
}

func (w *StreamingBatchWriter) flush(ctx context.Context, handle any, tableName string, upsertBatch bool, resources []*message.Insert) {
	// resources = w.removeDuplicatesByPK(table, resources)
	start := time.Now()
	batchSize := len(resources)
	if err := w.client.WriteTableStream(ctx, handle, upsertBatch, resources); err != nil {
		w.logger.Err(err).Str("table", tableName).Int("len", batchSize).Dur("duration", time.Since(start)).Msg("failed to write batch")
	} else {
		w.logger.Info().Str("table", tableName).Int("len", batchSize).Dur("duration", time.Since(start)).Msg("batch written successfully")
	}
}

/*
func (*StreamingBatchWriter) removeDuplicatesByPK(table *schema.Table, resources []arrow.Record) []arrow.Record {
	pkIndices := table.PrimaryKeysIndexes()
	// special case where there's no PK at all
	if len(pkIndices) == 0 {
		return resources
	}

	pks := make(map[string]struct{}, len(resources))
	res := make([]arrow.Record, 0, len(resources))
	for _, r := range resources {
		if r.NumRows() > 1 {
			panic(fmt.Sprintf("record with more than 1 row: %d", r.NumRows()))
		}
		key := pk.String(r)
		_, ok := pks[key]
		if !ok {
			pks[key] = struct{}{}
			res = append(res, r)
			continue
		}
		// duplicate, release
		r.Release()
	}

	return res
}
*/

func (*StreamingBatchWriter) getSourceName(r arrow.Record) string {
	colIndexes := r.Schema().FieldIndices(schema.CqSourceNameColumn.Name)
	if len(colIndexes) < 1 {
		return ""
	}
	return r.Column(colIndexes[0]).(*array.String).Value(0)
}

func (*StreamingBatchWriter) getSyncTime(r arrow.Record) time.Time {
	colIndexes := r.Schema().FieldIndices(schema.CqSyncTimeColumn.Name)
	if len(colIndexes) < 1 {
		return time.Time{}
	}

	return r.Column(colIndexes[0]).(*array.Timestamp).Value(0).ToTime(arrow.Microsecond).UTC()
}

func (w *StreamingBatchWriter) flushMigrateTables(ctx context.Context) error {
	w.migrateTableLock.Lock()
	defer w.migrateTableLock.Unlock()
	if len(w.migrateTableMessages) == 0 {
		return nil
	}
	if err := w.client.MigrateTables(ctx, w.migrateTableMessages); err != nil {
		return err
	}
	w.migrateTableMessages = w.migrateTableMessages[:0]
	return nil
}

func (w *StreamingBatchWriter) flushDeleteStaleTables(ctx context.Context) error {
	w.deleteStaleLock.Lock()
	defer w.deleteStaleLock.Unlock()
	if len(w.deleteStaleMessages) == 0 {
		return nil
	}
	if err := w.client.DeleteStale(ctx, w.deleteStaleMessages); err != nil {
		return err
	}
	w.deleteStaleMessages = w.deleteStaleMessages[:0]
	return nil
}

func (w *StreamingBatchWriter) flushInsert(_ context.Context, partitionKey string) {
	w.workersLock.RLock()
	worker, ok := w.workers[partitionKey]
	if !ok {
		w.workersLock.RUnlock()
		// no tables to flush
		return
	}
	w.workersLock.RUnlock()
	ch := make(chan bool)
	worker.flush <- ch
	<-ch
}

func (w *StreamingBatchWriter) flushInsertByTableName(ctx context.Context, tableName string) {
	var keys []string

	w.workersLock.RLock()
	for k := range w.workers {
		if w.isPartitionKeyForTable(k, tableName) {
			keys = append(keys, k)
		}
	}
	w.workersLock.RUnlock()

	for _, k := range keys {
		w.flushInsert(ctx, k)
	}
}

func (w *StreamingBatchWriter) Write(ctx context.Context, msgs <-chan message.Message) error {
	for msg := range msgs {
		switch m := msg.(type) {
		case *message.DeleteStale:
			if err := w.flushMigrateTables(ctx); err != nil {
				return err
			}
			w.flushInsert(ctx, w.makePartitionKey(m.SourceName, m.Table.Name))
			w.deleteStaleLock.Lock()
			w.deleteStaleMessages = append(w.deleteStaleMessages, m)
			l := len(w.deleteStaleMessages)
			w.deleteStaleLock.Unlock()
			if l > w.batchSize {
				if err := w.flushDeleteStaleTables(ctx); err != nil {
					return err
				}
			}
		case *message.Insert:
			if err := w.flushMigrateTables(ctx); err != nil {
				return err
			}
			if err := w.flushDeleteStaleTables(ctx); err != nil {
				return err
			}
			if err := w.startWorker(ctx, m); err != nil {
				return err
			}
		case *message.MigrateTable:
			w.flushInsertByTableName(ctx, m.Table.Name)
			if err := w.flushDeleteStaleTables(ctx); err != nil {
				return err
			}
			w.migrateTableLock.Lock()
			w.migrateTableMessages = append(w.migrateTableMessages, m)
			l := len(w.migrateTableMessages)
			w.migrateTableLock.Unlock()
			if l > w.batchSize {
				if err := w.flushMigrateTables(ctx); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (w *StreamingBatchWriter) startWorker(ctx context.Context, msg *message.Insert) error {
	w.workersLock.RLock()
	md := msg.Record.Schema().Metadata()
	tableName, ok := md.GetValue(schema.MetadataTableName)
	if !ok {
		w.workersLock.RUnlock()
		return fmt.Errorf("table name not found in metadata")
	}

	sourceName := w.getSourceName(msg.Record)
	partitionKey := w.makePartitionKey(sourceName, tableName)

	wr, ok := w.workers[partitionKey]
	w.workersLock.RUnlock()
	if ok {
		wr.ch <- msg
		return nil
	}
	w.workersLock.Lock()
	ch := make(chan *message.Insert)
	flush := make(chan chan bool)
	wr = &worker{
		count: 1,
		ch:    ch,
		flush: flush,
	}
	w.workers[partitionKey] = wr
	w.workersLock.Unlock()
	w.workersWaitGroup.Add(1)
	go func() {
		defer w.workersWaitGroup.Done()
		w.worker(ctx, sourceName, tableName, ch, flush)
	}()
	ch <- msg
	return nil
}

func (*StreamingBatchWriter) makePartitionKey(sourceName, tableName string) string {
	return sourceName + ":" + tableName
}

func (*StreamingBatchWriter) isPartitionKeyForTable(key, tableName string) bool {
	return strings.HasSuffix(key, ":"+tableName)
}

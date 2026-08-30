package client

import (
	"context"
	"slices"
	"sync"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/writers"
	"github.com/rs/zerolog"
	"golang.org/x/sync/semaphore"
)

// concurrentWriter is a drop-in replacement for the SDK's MixedBatchWriter that
// applies insert batches on several connections at once.
//
// The SDK writer accumulates a batch and calls InsertBatch synchronously, so a
// destination writes one batch at a time on one connection no matter how much
// capacity the database has. This writer keeps the same batching, but hands each
// full insert batch to a bounded pool of goroutines. Blocking on a full pool is
// what propagates backpressure back up the message channel.
//
// Message types other than inserts (migrations, delete-stale, delete-record) are
// rare and order-sensitive, so they stay synchronous, and every change of message
// type drains the pool first. That keeps the migrate -> insert -> delete-stale
// ordering the SDK writer guarantees.
//
// Ordering caveat: two insert batches in flight at once can be applied in either
// order, so a primary key repeated across batches no longer resolves to the last
// one written. Concurrency is therefore opt-in per spec, and defaults to 1.
type concurrentWriter struct {
	client         *Client
	logger         zerolog.Logger
	batchSize      int64
	batchSizeBytes int64
	batchTimeout   time.Duration
	concurrency    int64
}

func newConcurrentWriter(c *Client, logger zerolog.Logger, batchSize, batchSizeBytes, concurrency int64, batchTimeout time.Duration) *concurrentWriter {
	return &concurrentWriter{
		client:         c,
		logger:         logger,
		batchSize:      batchSize,
		batchSizeBytes: batchSizeBytes,
		batchTimeout:   batchTimeout,
		concurrency:    concurrency,
	}
}

// asyncFlusher runs up to `concurrency` flushes at a time and remembers the first
// error, which is returned by the next submit or by drain.
type asyncFlusher struct {
	sem *semaphore.Weighted
	wg  sync.WaitGroup

	mu  sync.Mutex
	err error
}

func newAsyncFlusher(concurrency int64) *asyncFlusher {
	return &asyncFlusher{sem: semaphore.NewWeighted(concurrency)}
}

func (a *asyncFlusher) submit(ctx context.Context, fn func() error) error {
	if err := a.load(); err != nil {
		return err
	}
	// Blocks once the pool is full, which is what pushes backpressure upstream.
	if err := a.sem.Acquire(ctx, 1); err != nil {
		return err
	}
	a.wg.Add(1)
	go func() {
		defer a.sem.Release(1)
		defer a.wg.Done()
		if err := fn(); err != nil {
			a.store(err)
		}
	}()
	return nil
}

// drain waits for every in-flight flush. It is the barrier that keeps message
// types from interleaving.
func (a *asyncFlusher) drain() error {
	a.wg.Wait()
	return a.load()
}

func (a *asyncFlusher) load() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.err
}

func (a *asyncFlusher) store(err error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.err == nil {
		a.err = err
	}
}

func (w *concurrentWriter) Write(ctx context.Context, msgChan <-chan message.WriteMessage) error {
	flusher := newAsyncFlusher(w.concurrency)

	inserts := make(message.WriteInserts, 0, w.batchSize)
	var insertRows, insertBytes int64

	// flushInserts hands the accumulated batch to the pool. The batch slice is
	// cloned because the loop below keeps appending into the same backing array,
	// which a worker would otherwise read while it is being overwritten.
	flushInserts := func() error {
		if len(inserts) == 0 {
			return nil
		}
		owned := slices.Clone(inserts)
		rows := insertRows
		inserts, insertRows, insertBytes = inserts[:0], 0, 0
		return flusher.submit(ctx, func() error {
			start := time.Now()
			err := w.client.InsertBatch(ctx, owned)
			if err != nil {
				w.logger.Err(err).Int64("len", rows).Dur("duration", time.Since(start)).Msg("failed to write batch")
				return err
			}
			w.logger.Debug().Int64("len", rows).Dur("duration", time.Since(start)).Msg("batch written successfully")
			return nil
		})
	}

	// The remaining message types are low volume and order-sensitive, so they are
	// applied inline, after the insert pool has drained.
	var migrates message.WriteMigrateTables
	var deleteStales message.WriteDeleteStales
	var deleteRecords message.WriteDeleteRecords

	flush := func(msgType writers.MsgType) error {
		switch msgType {
		case writers.MsgTypeUnset:
			return nil
		case writers.MsgTypeInsert:
			return flushInserts()
		case writers.MsgTypeMigrateTable:
			if len(migrates) == 0 {
				return nil
			}
			batch := migrates
			migrates = migrates[:0]
			return w.client.MigrateTableBatch(ctx, batch)
		case writers.MsgTypeDeleteStale:
			if len(deleteStales) == 0 {
				return nil
			}
			batch := deleteStales
			deleteStales = deleteStales[:0]
			return w.client.DeleteStaleBatch(ctx, batch)
		case writers.MsgTypeDeleteRecord:
			if len(deleteRecords) == 0 {
				return nil
			}
			batch := deleteRecords
			deleteRecords = deleteRecords[:0]
			return w.client.DeleteRecordsBatch(ctx, batch)
		default:
			panic("unknown message type")
		}
	}

	prevMsgType := writers.MsgTypeUnset
	ticker := writers.NewTicker(w.batchTimeout)
	defer ticker.Stop()

loop:
	for {
		select {
		case msg, ok := <-msgChan:
			if !ok {
				break loop
			}
			msgType := writers.MsgID(msg)
			if prevMsgType != msgType {
				if err := flush(prevMsgType); err != nil {
					return err
				}
				// Barrier: nothing of the next type starts until every in-flight
				// insert has landed.
				if err := flusher.drain(); err != nil {
					return err
				}
				ticker.Reset(w.batchTimeout)
			}
			prevMsgType = msgType

			switch v := msg.(type) {
			case *message.WriteInsert:
				inserts = append(inserts, v)
				insertRows += v.Record.NumRows()
				insertBytes += recordSize(v.Record)
				if insertRows >= w.batchSize || insertBytes >= w.batchSizeBytes {
					if err := flushInserts(); err != nil {
						return err
					}
				}
			case *message.WriteMigrateTable:
				migrates = append(migrates, v)
			case *message.WriteDeleteStale:
				deleteStales = append(deleteStales, v)
			case *message.WriteDeleteRecord:
				deleteRecords = append(deleteRecords, v)
			default:
				panic("unknown message type")
			}
		case <-ticker.Chan():
			if err := flush(prevMsgType); err != nil {
				return err
			}
			prevMsgType = writers.MsgTypeUnset
		}
	}

	if err := flush(prevMsgType); err != nil {
		return err
	}
	return flusher.drain()
}

// recordSize approximates a record's in-memory size from its buffers, standing in
// for the SDK's batch accounting, which lives in an internal package.
func recordSize(r arrow.RecordBatch) int64 {
	var total int64
	for _, col := range r.Columns() {
		total += arrayDataSize(col.Data())
	}
	return total
}

func arrayDataSize(d arrow.ArrayData) int64 {
	if d == nil {
		return 0
	}
	var total int64
	for _, buf := range d.Buffers() {
		if buf != nil {
			total += int64(buf.Len())
		}
	}
	for _, child := range d.Children() {
		total += arrayDataSize(child)
	}
	return total
}

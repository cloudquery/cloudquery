package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"slices"
	gosync "sync"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/cloudquery/cloudquery/cli/v6/internal/transformerpipeline"
)

// Checkpoints closer together than this are persisted together, to avoid flushing destinations too often.
const checkpointInterval = 30 * time.Second

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
	client      grpc.BidiStreamingClient[plugin.Write_Request, plugin.Write_Response]
	mu          gosync.Mutex
	lastFlushID uint64
}

func newFlushingWriteClient(client grpc.BidiStreamingClient[plugin.Write_Request, plugin.Write_Response]) *flushingWriteClient {
	return &flushingWriteClient{client: client}
}

func (s *flushingWriteClient) Send(req *plugin.Write_Request) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.client.Send(req)
}

// Flush returns once the destination commits every message sent before it.
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

	schema     *arrow.Schema
	pending    map[string]stateRow
	persisted  map[string]uint64
	lastCommit time.Time
}

func newCheckpointer(destinations []*flushingWriteClient, pipelines []*transformerpipeline.TransformerPipeline, backend plugin.PluginClient, backendFlushes bool, stateTable string) *checkpointer {
	return &checkpointer{
		destinations:   destinations,
		pipelines:      pipelines,
		backend:        backend,
		backendFlushes: backendFlushes,
		stateTable:     stateTable,
		pending:        make(map[string]stateRow),
		persisted:      make(map[string]uint64),
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
	keys, ok := stateColumn[*array.String](record, stateKeyColumn)
	if !ok {
		return fmt.Errorf("checkpoint has no %s string column", stateKeyColumn)
	}
	values, ok := stateColumn[*array.String](record, stateValueColumn)
	if !ok {
		return fmt.Errorf("checkpoint has no %s string column", stateValueColumn)
	}
	versions, versioned := stateColumn[*array.Uint64](record, stateVersionColumn)

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

func stateColumn[T arrow.Array](record arrow.RecordBatch, name string) (T, bool) {
	var zero T
	indices := record.Schema().FieldIndices(name)
	if len(indices) != 1 {
		return zero, false
	}
	col, ok := record.Column(indices[0]).(T)
	return col, ok
}

// due reports whether the pending state should be persisted now.
func (c *checkpointer) due() bool {
	return len(c.pending) > 0 && time.Since(c.lastCommit) >= checkpointInterval
}

// commit persists the pending state after every destination commits the records sent so far.
func (c *checkpointer) commit(ctx context.Context) error {
	if len(c.pending) == 0 {
		return nil
	}
	for _, pipeline := range c.pipelines {
		if err := pipeline.WaitDelivered(ctx); err != nil {
			return fmt.Errorf("failed to wait for records to reach destinations: %w", err)
		}
	}
	eg := errgroup.Group{}
	for _, destination := range c.destinations {
		eg.Go(destination.Flush)
	}
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("failed to flush destinations: %w", err)
	}
	if err := c.persist(ctx); err != nil {
		return fmt.Errorf("failed to persist state: %w", err)
	}
	for key, row := range c.pending {
		c.persisted[key] = row.version
	}
	clear(c.pending)
	c.lastCommit = time.Now()
	return nil
}

func (c *checkpointer) persist(ctx context.Context) error {
	recordBytes, err := c.pendingRecord()
	if err != nil {
		return err
	}
	insert := &plugin.Write_Request{Message: &plugin.Write_Request_Insert{Insert: &plugin.Write_MessageInsert{Record: recordBytes}}}

	if !c.backendFlushes {
		stream, err := c.backend.Write(ctx)
		if err != nil {
			return err
		}
		if err := stream.Send(insert); err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		_, err = stream.CloseAndRecv()
		return err
	}

	if c.backendStream == nil {
		stream, err := c.backend.WriteStream(ctx)
		if err != nil {
			return err
		}
		c.backendStream = newFlushingWriteClient(stream)
	}
	if err := c.backendStream.Send(insert); err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return c.backendStream.Flush()
}

func (c *checkpointer) pendingRecord() ([]byte, error) {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, c.schema)
	defer bldr.Release()
	keys := make([]string, 0, len(c.pending))
	for key := range c.pending {
		keys = append(keys, key)
	}
	slices.Sort(keys)
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

package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"slices"
	"strings"
	gosync "sync"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// eventLog records what the fake plugins received, in order.
type eventLog struct {
	mu     gosync.Mutex
	events []string
}

func (l *eventLog) add(event string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.events = append(l.events, event)
}

func (l *eventLog) get() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return slices.Clone(l.events)
}

type fakeWritePlugin struct {
	plugin.UnimplementedPluginServer
	name     string
	log      *eventLog
	flushErr error
}

func (p *fakeWritePlugin) record(req *plugin.Write_Request) error {
	switch m := req.Message.(type) {
	case *plugin.Write_Request_Insert:
		record, err := plugin.NewRecordFromBytes(m.Insert.Record)
		if err != nil {
			return err
		}
		tableName, _ := record.Schema().Metadata().GetValue(schema.MetadataTableName)
		event := fmt.Sprintf("%s: insert %s", p.name, tableName)
		if keys, ok := stateColumn[*array.String](record, stateKeyColumn); ok {
			values, _ := stateColumn[*array.String](record, stateValueColumn)
			for i := range keys.Len() {
				event += fmt.Sprintf(" %s=%s", keys.Value(i), values.Value(i))
			}
		}
		p.log.add(event)
	case *plugin.Write_Request_Flush:
		p.log.add(fmt.Sprintf("%s: flush %d", p.name, m.Flush.Id))
	}
	return nil
}

func (p *fakeWritePlugin) Write(stream plugin.Plugin_WriteServer) error {
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			p.log.add(p.name + ": close")
			return stream.SendAndClose(&plugin.Write_Response{})
		}
		if err != nil {
			return err
		}
		if err := p.record(req); err != nil {
			return err
		}
	}
}

func (p *fakeWritePlugin) WriteStream(stream plugin.Plugin_WriteStreamServer) error {
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			p.log.add(p.name + ": close")
			return nil
		}
		if err != nil {
			return err
		}
		if err := p.record(req); err != nil {
			return err
		}
		if flush := req.GetFlush(); flush != nil {
			if p.flushErr != nil {
				return p.flushErr
			}
			if err := stream.Send(&plugin.Write_Response{FlushId: flush.Id}); err != nil {
				return err
			}
		}
	}
}

func serveFakePlugin(t *testing.T, p *fakeWritePlugin) plugin.PluginClient {
	t.Helper()
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	plugin.RegisterPluginServer(server, p)
	go func() { _ = server.Serve(listener) }()
	t.Cleanup(server.Stop)

	conn, err := grpc.NewClient("passthrough:///bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return listener.Dial() }),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	t.Cleanup(func() { _ = conn.Close() })
	return plugin.NewPluginClient(conn)
}

func openFlushingClient(t *testing.T, client plugin.PluginClient) *flushingWriteClient {
	t.Helper()
	stream, err := client.WriteStream(context.Background())
	require.NoError(t, err)
	return newFlushingWriteClient(stream)
}

type stateEntry struct {
	key, value string
	version    uint64
}

func stateRecord(t *testing.T, tableName string, entries ...stateEntry) []byte {
	t.Helper()
	table := &schema.Table{
		Name: tableName,
		Columns: schema.ColumnList{
			{Name: stateKeyColumn, Type: arrow.BinaryTypes.String, PrimaryKey: true},
			{Name: stateValueColumn, Type: arrow.BinaryTypes.String},
			{Name: stateVersionColumn, Type: arrow.PrimitiveTypes.Uint64},
		},
	}
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	defer bldr.Release()
	for _, e := range entries {
		bldr.Field(0).(*array.StringBuilder).Append(e.key)
		bldr.Field(1).(*array.StringBuilder).Append(e.value)
		bldr.Field(2).(*array.Uint64Builder).Append(e.version)
	}
	recordBytes, err := plugin.RecordToBytes(bldr.NewRecordBatch())
	require.NoError(t, err)
	return recordBytes
}

func TestCheckpointerPersistsAfterDestinationsFlush(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	dest1 := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest1", log: log}))
	dest2 := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest2", log: log}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log})

	c := newCheckpointer([]*flushingWriteClient{dest1, dest2}, nil, backend, true, "cq_state")
	require.NoError(t, c.commit(ctx), "nothing to commit")
	require.Empty(t, log.get())

	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"a", "1", 1}, stateEntry{"b", "1", 1})))
	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"a", "2", 2})))
	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"b", "0", 0})))
	require.NoError(t, c.commit(ctx))
	events := log.get()
	require.ElementsMatch(t, []string{"dest1: flush 1", "dest2: flush 1"}, events[:2])
	require.Equal(t, []string{"backend: insert cq_state a=2 b=1", "backend: flush 1"}, events[2:])

	// versions that were persisted already are skipped
	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"a", "1", 1})))
	require.NoError(t, c.commit(ctx))
	require.Len(t, log.get(), 4)

	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"a", "3", 3})))
	require.NoError(t, c.commit(ctx))
	require.Equal(t, []string{"backend: insert cq_state a=3", "backend: flush 2"}, log.get()[6:])

	require.NoError(t, c.close())
	require.Eventually(t, func() bool { return slices.Contains(log.get(), "backend: close") }, 5*time.Second, 10*time.Millisecond)
}

func TestCheckpointerWritesToBackendWithoutFlushSupport(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log})

	c := newCheckpointer([]*flushingWriteClient{dest}, nil, backend, false, "cq_state")
	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"a", "1", 1})))
	require.NoError(t, c.commit(ctx))
	require.Equal(t, []string{"dest: flush 1", "backend: insert cq_state a=1", "backend: close"}, log.get())
	require.NoError(t, c.close())
}

func TestCheckpointerKeepsStateWhenDestinationFlushFails(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log, flushErr: errors.New("disk full")}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log})

	c := newCheckpointer([]*flushingWriteClient{dest}, nil, backend, true, "cq_state")
	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"a", "1", 1})))
	err := c.commit(ctx)
	require.ErrorContains(t, err, "disk full")
	for _, event := range log.get() {
		require.False(t, strings.HasPrefix(event, "backend:"), "state must not be persisted: %s", event)
	}
}

func TestCheckpointerRejectsOtherTables(t *testing.T) {
	c := newCheckpointer(nil, nil, nil, false, "cq_state")
	require.ErrorContains(t, c.add(stateRecord(t, "other", stateEntry{"a", "1", 1})), `checkpoint is for table "other", expected the state table "cq_state"`)
}

func TestCheckpointerIsDueAfterInterval(t *testing.T) {
	c := newCheckpointer(nil, nil, nil, false, "cq_state")
	require.False(t, c.due(), "nothing pending")
	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"a", "1", 1})))
	require.False(t, c.due(), "interval not elapsed")
	c.lastCommit = time.Now().Add(-checkpointInterval)
	require.True(t, c.due())
}

func TestCheckpointerCommitsWhileSourceIsIdle(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log})
	c := newCheckpointer([]*flushingWriteClient{dest}, nil, backend, true, "cq_state")

	responses := make(chan syncResult)
	recv := c.receiver(ctx, func() (*plugin.Sync_Response, error) {
		r := <-responses
		return r.response, r.err
	})

	checkpoint := &plugin.Sync_Response{Message: &plugin.Sync_Response_Checkpoint{Checkpoint: &plugin.Sync_MessageCheckpoint{Record: stateRecord(t, "cq_state", stateEntry{"a", "1", 1})}}}
	responses <- syncResult{response: checkpoint}
	r, err := recv()
	require.NoError(t, err)
	require.NoError(t, c.add(r.GetCheckpoint().Record))
	c.lastCommit = time.Now().Add(-checkpointInterval + 50*time.Millisecond)

	received := make(chan error)
	go func() {
		_, err := recv()
		received <- err
	}()
	require.Eventually(t, func() bool {
		return slices.Equal(log.get(), []string{"dest: flush 1", "backend: insert cq_state a=1", "backend: flush 1"})
	}, 5*time.Second, 10*time.Millisecond, "pending state is persisted while the source sends nothing")

	responses <- syncResult{err: io.EOF}
	require.ErrorIs(t, <-received, io.EOF)
	require.NoError(t, c.close())
}

func TestCheckpointerSplitsLargeState(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log})
	c := newCheckpointer([]*flushingWriteClient{dest}, nil, backend, true, "cq_state")
	c.maxRecordBytes = 60

	require.NoError(t, c.add(stateRecord(t, "cq_state",
		stateEntry{"a", "0123456789", 1},
		stateEntry{"b", "0123456789", 1},
		stateEntry{"c", "0123456789", 1},
		stateEntry{"d", strings.Repeat("x", 100), 1},
		stateEntry{"e", "0123456789", 1},
	)))
	require.NoError(t, c.commit(ctx))
	require.Equal(t, []string{
		"dest: flush 1",
		"backend: insert cq_state a=0123456789 b=0123456789",
		"backend: insert cq_state c=0123456789",
		"backend: insert cq_state d=" + strings.Repeat("x", 100),
		"backend: insert cq_state e=0123456789",
		"backend: flush 1",
	}, log.get())
	require.NoError(t, c.close())
}

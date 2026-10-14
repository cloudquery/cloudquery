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
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	discoveryPb "github.com/cloudquery/plugin-pb-go/pb/discovery/v1"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/v6/internal/transformerpipeline"
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
	name          string
	log           *eventLog
	flushErr      error
	supportsFlush bool
	// acknowledges flushes with this id instead of the requested one
	flushID uint64
	// returned by WriteStream on the first insert
	insertErr error
}

func (p *fakeWritePlugin) Init(context.Context, *plugin.Init_Request) (*plugin.Init_Response, error) {
	return &plugin.Init_Response{SupportsFlush: p.supportsFlush}, nil
}

func (*fakeWritePlugin) Close(context.Context, *plugin.Close_Request) (*plugin.Close_Response, error) {
	return &plugin.Close_Response{}, nil
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
		if p.insertErr != nil && req.GetInsert() != nil {
			return p.insertErr
		}
		if err := p.record(req); err != nil {
			return err
		}
		if flush := req.GetFlush(); flush != nil {
			if p.flushErr != nil {
				return p.flushErr
			}
			id := flush.Id
			if p.flushID != 0 {
				id = p.flushID
			}
			if err := stream.Send(&plugin.Write_Response{FlushId: id}); err != nil {
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
	return newFlushingWriteClient("dest", stream)
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
	require.ErrorContains(t, err, "destination dest")
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

	// the source sends nothing, so the receiver reports the pending state as due
	_, err = recv()
	require.ErrorIs(t, err, errCheckpointDue)
	require.NoError(t, c.commit(ctx))
	require.Equal(t, []string{"dest: flush 1", "backend: insert cq_state a=1", "backend: flush 1"}, log.get())

	go func() { responses <- syncResult{err: io.EOF} }()
	_, err = recv()
	require.ErrorIs(t, err, io.EOF)
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

func TestCheckpointerRejectsMismatchedFlushAcknowledgement(t *testing.T) {
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log, flushID: 7}))
	require.ErrorContains(t, dest.Flush(), "destination acknowledged flush 7, expected 1")
}

func TestCheckpointerValidatesColumns(t *testing.T) {
	stateTable := func(columns ...schema.Column) *schema.Table {
		return &schema.Table{Name: "cq_state", Columns: columns}
	}
	key := schema.Column{Name: stateKeyColumn, Type: arrow.BinaryTypes.String}
	value := schema.Column{Name: stateValueColumn, Type: arrow.BinaryTypes.String}
	cases := []struct {
		name  string
		table *schema.Table
		err   string
	}{
		{name: "no key", table: stateTable(value), err: `checkpoint has no "key" column`},
		{name: "large string key", table: stateTable(schema.Column{Name: stateKeyColumn, Type: arrow.BinaryTypes.LargeString}, value), err: `checkpoint column "key" is large_utf8, expected utf8`},
		{name: "int64 version", table: stateTable(key, value, schema.Column{Name: stateVersionColumn, Type: arrow.PrimitiveTypes.Int64}), err: `checkpoint column "version" is int64, expected uint64`},
		{name: "no version", table: stateTable(key, value)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bldr := array.NewRecordBuilder(memory.DefaultAllocator, tc.table.ToArrowSchema())
			defer bldr.Release()
			for _, field := range bldr.Fields() {
				field.AppendEmptyValue()
			}
			recordBytes, err := plugin.RecordToBytes(bldr.NewRecordBatch())
			require.NoError(t, err)

			c := newCheckpointer(nil, nil, nil, false, "cq_state")
			err = c.add(recordBytes)
			if tc.err == "" {
				require.NoError(t, err)
				return
			}
			require.EqualError(t, err, tc.err)
		})
	}
}

func TestUnlessFailedElsewhere(t *testing.T) {
	failed := errors.New("flush failed")
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	require.NoError(t, unlessFailedElsewhere(context.Background(), context.Background(), transformerpipeline.ErrPipelineClosed), "the pipeline reports its own error")
	require.NoError(t, unlessFailedElsewhere(context.Background(), cancelled, context.Canceled), "another goroutine of the sync failed first")
	require.ErrorIs(t, unlessFailedElsewhere(cancelled, cancelled, context.Canceled), context.Canceled, "the sync was cancelled")
	require.ErrorIs(t, unlessFailedElsewhere(context.Background(), context.Background(), failed), failed)
}

func TestCheckpointOptionsStateWriter(t *testing.T) {
	pg, pg2, backend := plugin.NewPluginClient(nil), plugin.NewPluginClient(nil), plugin.NewPluginClient(nil)
	base := checkpointOptions{
		backendConnection: "@@plugins.pg.connection",
		destinationNames:  []string{"pg2", "pg"},
		destinations:      []plugin.PluginClient{pg2, pg},
		destinationsFlush: true,
	}
	cases := []struct {
		name        string
		change      func(*checkpointOptions)
		wantWriter  plugin.PluginClient
		wantFlushes bool
		wantReason  string
	}{
		{name: "backend is a destination", wantWriter: pg, wantFlushes: true},
		{name: "destination without flush", change: func(o *checkpointOptions) { o.destinationsFlush = false }, wantReason: "not all destinations support flushing"},
		{name: "transformers", change: func(o *checkpointOptions) { o.destinationTransformers = true }, wantReason: "destinations use transformers"},
		{
			name: "separate backend",
			change: func(o *checkpointOptions) {
				o.backendConnection = "@@plugins.state.connection"
				o.backend, o.backendFlush = backend, true
			},
			wantWriter:  backend,
			wantFlushes: true,
		},
		{
			name: "separate backend without flush",
			change: func(o *checkpointOptions) {
				o.backendConnection = "@@plugins.state.connection"
				o.backend = backend
			},
			wantWriter: backend,
		},
		{name: "backend address", change: func(o *checkpointOptions) { o.backendConnection = "localhost:7777" }, wantReason: "the state backend is not a plugin of this sync"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			o := base
			if tc.change != nil {
				tc.change(&o)
			}
			writer, flushes, reason := o.stateWriter()
			require.Equal(t, tc.wantReason, reason)
			require.Equal(t, tc.wantFlushes, flushes)
			require.True(t, writer == tc.wantWriter, "unexpected state writer")
		})
	}
}

type fakeCheckpointSource struct {
	plugin.UnimplementedPluginServer
	discoveryPb.UnimplementedDiscoveryServer
	responses []*plugin.Sync_Response
	// keeps the stream open without sending more, until the client goes away
	idle bool

	mu      gosync.Mutex
	request *plugin.Sync_Request
}

func (*fakeCheckpointSource) Init(context.Context, *plugin.Init_Request) (*plugin.Init_Response, error) {
	return &plugin.Init_Response{}, nil
}

func (s *fakeCheckpointSource) Sync(req *plugin.Sync_Request, stream plugin.Plugin_SyncServer) error {
	s.mu.Lock()
	s.request = req
	s.mu.Unlock()
	for _, res := range s.responses {
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	if s.idle {
		<-stream.Context().Done()
	}
	return nil
}

func startFakeServer(t *testing.T, register func(*grpc.Server)) string {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	server := grpc.NewServer()
	register(server)
	go func() { _ = server.Serve(listener) }()
	t.Cleanup(server.Stop)
	return listener.Addr().String()
}

func grpcPluginClient(t *testing.T, typ managedplugin.PluginType, name, address string) *managedplugin.Client {
	t.Helper()
	client, err := managedplugin.NewClient(context.Background(), typ, managedplugin.Config{Name: name, Registry: managedplugin.RegistryGrpc, Path: address})
	require.NoError(t, err)
	t.Cleanup(func() { _ = client.Terminate() })
	return client
}

func itemsInsert(t *testing.T, ids ...int64) *plugin.Sync_Response {
	t.Helper()
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, checkpointItemsTable.ToArrowSchema())
	defer bldr.Release()
	bldr.Field(0).(*array.Int64Builder).AppendValues(ids, nil)
	recordBytes, err := plugin.RecordToBytes(bldr.NewRecordBatch())
	require.NoError(t, err)
	return &plugin.Sync_Response{Message: &plugin.Sync_Response_Insert{Insert: &plugin.Sync_MessageInsert{Record: recordBytes}}}
}

func checkpointResponse(t *testing.T, entries ...stateEntry) *plugin.Sync_Response {
	t.Helper()
	return &plugin.Sync_Response{Message: &plugin.Sync_Response_Checkpoint{Checkpoint: &plugin.Sync_MessageCheckpoint{Record: stateRecord(t, "cq_state", entries...)}}}
}

var checkpointItemsTable = &schema.Table{Name: "items", Columns: schema.ColumnList{{Name: "id", Type: arrow.PrimitiveTypes.Int64}}}

// syncWithFakePlugins runs a sync from a fake source to a fake destination that is also the state backend.
func syncWithFakePlugins(t *testing.T, destination *fakeWritePlugin, responses ...*plugin.Sync_Response) (*plugin.Sync_Request, error) {
	t.Helper()
	return syncWithFakeSource(t, &fakeCheckpointSource{}, destination, responses...)
}

func syncWithFakeSource(t *testing.T, source *fakeCheckpointSource, destination *fakeWritePlugin, responses ...*plugin.Sync_Response) (*plugin.Sync_Request, error) {
	t.Helper()
	tableBytes, err := plugin.SchemaToBytes(checkpointItemsTable.ToArrowSchema())
	require.NoError(t, err)
	source.responses = append([]*plugin.Sync_Response{
		{Message: &plugin.Sync_Response_MigrateTable{MigrateTable: &plugin.Sync_MessageMigrateTable{Table: tableBytes}}},
	}, responses...)
	sourceAddress := startFakeServer(t, func(s *grpc.Server) {
		plugin.RegisterPluginServer(s, source)
		discoveryPb.RegisterDiscoveryServer(s, source)
	})
	destinationAddress := startFakeServer(t, func(s *grpc.Server) { plugin.RegisterPluginServer(s, destination) })

	skipDependentTables := true
	err = syncConnectionV3(context.Background(), syncV3Options{
		source: v3source{
			client: grpcPluginClient(t, managedplugin.PluginSource, "source", sourceAddress),
			spec: specs.Source{
				Metadata:            specs.Metadata{Name: "source", Path: sourceAddress},
				Tables:              []string{"*"},
				Destinations:        []string{"dest"},
				SkipDependentTables: &skipDependentTables,
				BackendOptions:      &specs.BackendOptions{TableName: "cq_state", Connection: "@@plugins.dest.connection"},
			},
		},
		destinations: []v3destination{{
			client: grpcPluginClient(t, managedplugin.PluginDestination, "dest", destinationAddress),
			spec:   specs.Destination{Metadata: specs.Metadata{Name: "dest", Path: destinationAddress}, WriteMode: specs.WriteModeAppend},
		}},
		uid: "test",
	})
	source.mu.Lock()
	defer source.mu.Unlock()
	return source.request, err
}

func TestSyncPersistsCheckpointsAfterDestinationsFlush(t *testing.T) {
	responses := func(t *testing.T) []*plugin.Sync_Response {
		return []*plugin.Sync_Response{
			itemsInsert(t, 1, 2, 3),
			checkpointResponse(t, stateEntry{"cursor", "1", 1}),
			itemsInsert(t, 4, 5),
			checkpointResponse(t, stateEntry{"cursor", "2", 2}),
		}
	}

	t.Run("each checkpoint", func(t *testing.T) {
		interval := checkpointInterval
		checkpointInterval = 0
		t.Cleanup(func() { checkpointInterval = interval })
		// repeat, as a flush that overtakes an insert on its way to the destination would show up only sometimes
		for range 10 {
			log := &eventLog{}
			req, err := syncWithFakePlugins(t, &fakeWritePlugin{name: "dest", log: log, supportsFlush: true}, responses(t)...)
			require.NoError(t, err)
			require.True(t, req.SupportsCheckpoints)
			require.Equal(t, []string{
				"dest: insert items",
				"dest: flush 1",
				"dest: insert cq_state cursor=1",
				"dest: flush 1",
				"dest: insert items",
				"dest: flush 2",
				"dest: insert cq_state cursor=2",
				"dest: flush 2",
			}, log.get()[:8])
		}
	})

	t.Run("coalesced", func(t *testing.T) {
		log := &eventLog{}
		_, err := syncWithFakePlugins(t, &fakeWritePlugin{name: "dest", log: log, supportsFlush: true}, responses(t)...)
		require.NoError(t, err)
		require.Equal(t, []string{
			"dest: insert items",
			"dest: insert items",
			"dest: flush 1",
			"dest: insert cq_state cursor=2",
			"dest: flush 1",
		}, log.get()[:5])
	})

	t.Run("destination without flush support", func(t *testing.T) {
		log := &eventLog{}
		req, err := syncWithFakePlugins(t, &fakeWritePlugin{name: "dest", log: log}, itemsInsert(t, 1), checkpointResponse(t, stateEntry{"cursor", "1", 1}))
		require.ErrorContains(t, err, "source sent a state checkpoint, but checkpoints are not enabled for this sync")
		require.False(t, req.SupportsCheckpoints)
		for _, event := range log.get() {
			require.NotContains(t, event, "flush")
			require.NotContains(t, event, "cq_state")
		}
	})
}

func TestSyncReportsDestinationErrorWhileSourceIsIdle(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	destination := &fakeWritePlugin{name: "dest", log: &eventLog{}, supportsFlush: true, insertErr: errors.New("disk full")}
	_, err := syncWithFakeSource(t, &fakeCheckpointSource{idle: true}, destination,
		itemsInsert(t, 1, 2),
		itemsInsert(t, 3),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
	)
	require.ErrorContains(t, err, "disk full")
	require.NotContains(t, err.Error(), "context canceled")
}

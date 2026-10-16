package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path"
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
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
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
	// returned by Write and WriteStream on the first insert
	insertErr error
	// reported in flush responses, as a destination does when it could not commit some records
	flushFailure string
	// reports flushFailure only from the flush with this id on
	flushFailureFrom uint64
	// reported with flushFailure
	failedTables []string
	// ends the stream with an error after reporting flushFailure, as writers that stop on a failed write do
	endAfterFailure bool
	// returned by WriteStream once its messages end
	closeErr error
	// reported instead of flushFailure from the flush after the first failing one on
	laterFailure      string
	laterFailedTables []string
	// returned by the Close RPC
	closeRPCErr error
}

func (p *fakeWritePlugin) Init(context.Context, *plugin.Init_Request) (*plugin.Init_Response, error) {
	return &plugin.Init_Response{SupportsFlush: p.supportsFlush}, nil
}

func (p *fakeWritePlugin) Close(context.Context, *plugin.Close_Request) (*plugin.Close_Response, error) {
	return &plugin.Close_Response{}, p.closeRPCErr
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
	case *plugin.Write_Request_Delete:
		p.log.add(fmt.Sprintf("%s: delete stale %s", p.name, m.Delete.TableName))
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
		if p.insertErr != nil && req.GetInsert() != nil {
			return p.insertErr
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
			return p.closeErr
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
			res := &plugin.Write_Response{FlushId: id}
			switch {
			case p.laterFailure != "" && flush.Id > max(p.flushFailureFrom, 1):
				res.FlushError = p.laterFailure
				res.FailedTables = p.laterFailedTables
			case flush.Id >= p.flushFailureFrom:
				res.FlushError = p.flushFailure
				res.FailedTables = p.failedTables
			}
			if err := stream.Send(res); err != nil {
				return err
			}
			if res.FlushError != "" && p.endAfterFailure {
				p.log.add(p.name + ": end")
				return errors.New(res.FlushError)
			}
		}
	}
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
	// fails the test instead of hanging it if the receiver stops returning
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log})
	c := newCheckpointer([]*flushingWriteClient{dest}, nil, backend, true, "cq_state")

	responses := make(chan syncResult)
	recv := c.receiver(ctx, func() (*plugin.Sync_Response, error) {
		select {
		case r := <-responses:
			return r.response, r.err
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	})

	checkpoint := &plugin.Sync_Response{Message: &plugin.Sync_Response_Checkpoint{Checkpoint: &plugin.Sync_MessageCheckpoint{Record: stateRecord(t, "cq_state", stateEntry{"a", "1", 1})}}}
	select {
	case responses <- syncResult{response: checkpoint}:
	case <-ctx.Done():
		t.Fatal("the receiver did not read from the source")
	}
	r, err := recv()
	require.NoError(t, err)
	require.NoError(t, c.add(r.GetCheckpoint().Record))
	c.lastCommit = time.Now().Add(-checkpointInterval + 50*time.Millisecond)

	// the source sends nothing, so the receiver reports the pending state as due
	_, err = recv()
	require.ErrorIs(t, err, errCheckpointDue)
	require.NoError(t, c.commit(ctx))
	require.Equal(t, []string{"dest: flush 1", "backend: insert cq_state a=1", "backend: flush 1"}, log.get())

	go func() {
		select {
		case responses <- syncResult{err: io.EOF}:
		case <-ctx.Done():
		}
	}()
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
	c.maxRecordBytes = 110

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
		{name: "two keys", table: &schema.Table{Name: "cq_state", Columns: schema.ColumnList{key, value, key}}, err: `checkpoint has 2 "key" columns`},
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
		{
			name: "destination name is a prefix of the backend's",
			change: func(o *checkpointOptions) {
				o.backendConnection = "@@plugins.pg2.connection"
				o.destinationNames = []string{"pg", "pg2"}
				o.destinations = []plugin.PluginClient{pg, pg2}
			},
			wantWriter:  pg2,
			wantFlushes: true,
		},
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
	// keeps the stream open without sending more, until closed
	release chan struct{}

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
	if s.release != nil {
		select {
		case <-s.release:
		case <-stream.Context().Done():
		}
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
	return syncWithFakeSourceMode(t, source, destination, specs.WriteModeAppend, responses...)
}

func syncWithFakeSourceMode(t *testing.T, source *fakeCheckpointSource, destination *fakeWritePlugin, writeMode specs.WriteMode, responses ...*plugin.Sync_Response) (*plugin.Sync_Request, error) {
	t.Helper()
	err := syncWithFakeDestinations(t, source, []*fakeWritePlugin{destination}, writeMode, responses...)
	source.mu.Lock()
	defer source.mu.Unlock()
	return source.request, err
}

// syncWithFakeDestinations runs a sync from a fake source to fake destinations, the first of which is also the state
// backend.
func syncWithFakeDestinations(t *testing.T, source *fakeCheckpointSource, destinations []*fakeWritePlugin, writeMode specs.WriteMode, responses ...*plugin.Sync_Response) error {
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
	names := make([]string, 0, len(destinations))
	v3destinations := make([]v3destination, 0, len(destinations))
	for _, destination := range destinations {
		address := startFakeServer(t, func(s *grpc.Server) { plugin.RegisterPluginServer(s, destination) })
		names = append(names, destination.name)
		v3destinations = append(v3destinations, v3destination{
			client: grpcPluginClient(t, managedplugin.PluginDestination, destination.name, address),
			spec:   specs.Destination{Metadata: specs.Metadata{Name: destination.name, Path: address}, WriteMode: writeMode},
		})
	}

	skipDependentTables := true
	return syncConnectionV3(context.Background(), syncV3Options{
		source: v3source{
			client: grpcPluginClient(t, managedplugin.PluginSource, "source", sourceAddress),
			spec: specs.Source{
				Metadata:            specs.Metadata{Name: "source", Path: sourceAddress},
				Tables:              []string{"*"},
				Destinations:        names,
				SkipDependentTables: &skipDependentTables,
				BackendOptions:      &specs.BackendOptions{TableName: "cq_state", Connection: "@@plugins." + names[0] + ".connection"},
			},
		},
		destinations: v3destinations,
		uid:          "test",
	})
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

// TestSyncReportsDestinationErrorFromCommit checks that a destination write error surfaces when the commit fails.
func TestSyncReportsDestinationErrorFromCommit(t *testing.T) {
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

func TestCommitWaitsForPipelineDelivery(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log})

	pipeline, _, err := transformerpipeline.New(ctx, nil)
	require.NoError(t, err)
	release := make(chan struct{})
	require.NoError(t, pipeline.OnOutput(func(record []byte) error {
		<-release
		return dest.Send(&plugin.Write_Request{Message: &plugin.Write_Request_Insert{Insert: &plugin.Write_MessageInsert{Record: record}}})
	}))
	go func() { _ = pipeline.RunBlocking() }()
	t.Cleanup(pipeline.Close)

	c := newCheckpointer([]*flushingWriteClient{dest}, []*transformerpipeline.TransformerPipeline{pipeline}, backend, true, "cq_state")
	require.NoError(t, pipeline.Send(stateRecord(t, "items", stateEntry{"x", "1", 1})))
	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"cursor", "1", 1})))

	committed := make(chan error, 1)
	go func() { committed <- c.commit(ctx) }()
	select {
	case err := <-committed:
		t.Fatalf("commit returned before the pipeline delivered its record: %v", err)
	case <-time.After(100 * time.Millisecond):
	}
	require.NotContains(t, log.get(), "dest: flush 1", "the flush must not overtake an undelivered insert")
	close(release)
	select {
	case err := <-committed:
		require.NoError(t, err)
	case <-time.After(5 * time.Second):
		t.Fatal("commit did not return after the record was delivered")
	}
	require.Equal(t, []string{"dest: insert items x=1", "dest: flush 1", "backend: insert cq_state cursor=1", "backend: flush 1"}, log.get())
	require.NoError(t, c.close())
}

func TestSyncPersistsCheckpointsWhileSourceIsIdle(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 100 * time.Millisecond
	t.Cleanup(func() { checkpointInterval = interval })

	log := &eventLog{}
	source := &fakeCheckpointSource{release: make(chan struct{})}
	synced := make(chan error, 1)
	go func() {
		_, err := syncWithFakeSource(t, source, &fakeWritePlugin{name: "dest", log: log, supportsFlush: true},
			itemsInsert(t, 1, 2),
			checkpointResponse(t, stateEntry{"cursor", "1", 1}),
		)
		synced <- err
	}()
	require.Eventually(t, func() bool {
		return slices.Contains(log.get(), "dest: insert cq_state cursor=1")
	}, 5*time.Second, 10*time.Millisecond, "the state is persisted while the source sends nothing")
	close(source.release)
	select {
	case err := <-synced:
		require.NoError(t, err)
	case <-time.After(10 * time.Second):
		t.Fatal("sync did not finish")
	}
}

func TestCheckpointerReportsBackendErrors(t *testing.T) {
	for _, flushes := range []bool{true, false} {
		t.Run(fmt.Sprintf("backend_flushes=%t", flushes), func(t *testing.T) {
			ctx := context.Background()
			log := &eventLog{}
			dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log}))
			backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log, insertErr: errors.New("disk full")})
			c := newCheckpointer([]*flushingWriteClient{dest}, nil, backend, flushes, "cq_state")
			// one record per key, so sending the later records hits the failed stream
			c.maxRecordBytes = 1
			entries := make([]stateEntry, 0, 20)
			for i := range 20 {
				entries = append(entries, stateEntry{fmt.Sprintf("key%02d", i), "value", 1})
			}
			require.NoError(t, c.add(stateRecord(t, "cq_state", entries...)))

			err := c.commit(ctx)
			require.ErrorContains(t, err, "failed to persist state")
			require.ErrorContains(t, err, "disk full")
		})
	}
}

func TestCheckpointerStopsSavingStateAfterFailedFlush(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log, flushFailure: "table items: disk full"}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log})
	c := newCheckpointer([]*flushingWriteClient{dest}, nil, backend, true, "cq_state")

	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"cursor", "1", 1})))
	require.NoError(t, c.commit(ctx), "the sync goes on after a failed flush")
	require.Empty(t, c.withoutFailedTables("dest", map[string]bool{"items": true}, syncedTables{names: map[string]bool{"items": true}}), "delete-stale is skipped")
	require.EqualError(t, c.failure(), "destination dest: table items: disk full")

	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"cursor", "2", 2})))
	require.Empty(t, c.pending, "checkpoints after the failure are not kept")
	c.lastCommit = time.Time{}
	require.False(t, c.due())
	require.NoError(t, c.commit(ctx))
	require.Equal(t, []string{"dest: flush 1"}, log.get(), "no more commits")
	require.NoError(t, c.finish(ctx))
	require.Equal(t, []string{"dest: flush 1", "dest: flush 2"}, log.get(), "the final flush asks about later failures, and no state is saved")
}

func TestSyncFinishesWithoutStateAfterFailedFlush(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	log := &eventLog{}
	destination := &fakeWritePlugin{name: "dest", log: log, supportsFlush: true, flushFailure: "table items: disk full"}
	_, err := syncWithFakeSourceMode(t, &fakeCheckpointSource{}, destination, specs.WriteModeOverwriteDeleteStale,
		itemsInsert(t, 1, 2),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
		itemsInsert(t, 3),
		checkpointResponse(t, stateEntry{"cursor", "2", 2}),
	)
	require.ErrorContains(t, err, "sync finished without saving its state")
	require.ErrorContains(t, err, "destination dest: table items: disk full")

	var inserts int
	for _, event := range log.get() {
		require.NotContains(t, event, "cq_state", "no state is saved")
		require.NotContains(t, event, "delete stale", "delete-stale is skipped on the failed destination")
		if event == "dest: insert items" {
			inserts++
		}
	}
	require.Equal(t, 2, inserts, "the sync goes on after the failed flush")
}

func TestSyncClosesTheStateStream(t *testing.T) {
	log := &eventLog{}
	_, err := syncWithFakePlugins(t, &fakeWritePlugin{name: "dest", log: log, supportsFlush: true},
		itemsInsert(t, 1),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
	)
	require.NoError(t, err)
	var closes int
	for _, event := range log.get() {
		if event == "dest: close" {
			closes++
		}
	}
	require.Equal(t, 2, closes, "both the records stream and the state stream end")
}

func TestSyncCatchesFailureAfterLastCheckpoint(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	log := &eventLog{}
	destination := &fakeWritePlugin{name: "dest", log: log, supportsFlush: true, flushFailure: "table items: disk full", flushFailureFrom: 2}
	_, err := syncWithFakeSourceMode(t, &fakeCheckpointSource{}, destination, specs.WriteModeOverwriteDeleteStale,
		itemsInsert(t, 1, 2),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
		itemsInsert(t, 3),
	)
	require.ErrorAs(t, err, new(*stateNotSavedError))
	require.ErrorContains(t, err, "destination dest: table items: disk full")
	require.Contains(t, log.get(), "dest: insert cq_state cursor=1", "the state before the failure is saved")
	for _, event := range log.get() {
		require.NotContains(t, event, "delete stale")
	}
}

func TestSyncFlushesOtherDestinationsAfterAFailure(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	log := &eventLog{}
	failing := &fakeWritePlugin{name: "dest", log: log, supportsFlush: true, flushFailure: "table items: disk full"}
	failingLater := &fakeWritePlugin{name: "dest2", log: log, supportsFlush: true, flushFailure: "table items: quota exceeded", flushFailureFrom: 2}
	healthy := &fakeWritePlugin{name: "dest3", log: log, supportsFlush: true}
	err := syncWithFakeDestinations(t, &fakeCheckpointSource{}, []*fakeWritePlugin{failing, failingLater, healthy}, specs.WriteModeOverwriteDeleteStale,
		itemsInsert(t, 1, 2),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
		itemsInsert(t, 3),
		checkpointResponse(t, stateEntry{"cursor", "2", 2}),
	)
	require.ErrorAs(t, err, new(*stateNotSavedError))
	require.ErrorContains(t, err, "destination dest: table items: disk full")
	require.ErrorContains(t, err, "destination dest2: table items: quota exceeded")
	require.NotContains(t, err.Error(), "dest3")

	events := log.get()
	require.Contains(t, events, "dest2: flush 2", "the final flush reaches destinations that did not fail yet")
	require.Contains(t, events, "dest: flush 2", "the final flush asks a failed destination about later failures")
	require.Contains(t, events, "dest3: delete stale items", "delete-stale runs on the destination that committed every record")
	for _, event := range events {
		require.NotContains(t, event, "cq_state", "no state is saved")
		require.False(t, strings.HasPrefix(event, "dest: delete stale") || strings.HasPrefix(event, "dest2: delete stale"), "delete-stale is skipped on failed destinations: %s", event)
	}
}

func TestSyncWithoutCheckpointsFromTheSourceWarnsThatItsStateMayBeAhead(t *testing.T) {
	destination := &fakeWritePlugin{name: "dest", log: &eventLog{}, supportsFlush: true, flushFailure: "table items: disk full"}
	req, err := syncWithFakePlugins(t, destination, itemsInsert(t, 1, 2))
	require.True(t, req.SupportsCheckpoints)
	require.ErrorAs(t, err, new(*stateNotSavedError))
	require.ErrorContains(t, err, "the source sent no state checkpoints: if it saves state itself, that state may be past the lost records")
	require.ErrorContains(t, err, "destination dest: table items: disk full")
}

func TestCheckpointerStopsSavingStateAfterBackendFlushFails(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log, flushFailure: "table cq_state: disk full"})
	c := newCheckpointer([]*flushingWriteClient{dest}, nil, backend, true, "cq_state")

	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"cursor", "1", 1})))
	require.NoError(t, c.commit(ctx), "the sync goes on, as it does when a destination fails")
	require.Equal(t, map[string]bool{"items": true}, c.withoutFailedTables("dest", map[string]bool{"items": true}, syncedTables{names: map[string]bool{"items": true}}), "delete-stale still runs on the destinations")
	require.EqualError(t, c.failure(), "state backend: table cq_state: disk full")

	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"cursor", "2", 2})))
	require.Empty(t, c.pending)
	require.NoError(t, c.finish(ctx))
	require.Equal(t, []string{"dest: flush 1", "backend: insert cq_state cursor=1", "backend: flush 1", "dest: flush 2"}, log.get())
	require.NoError(t, c.close())
}

func TestCheckpointerReportsCommitFailuresWithHardErrors(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	failing := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log, flushFailure: "table items: disk full"}))
	broken := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest2", log: log, flushErr: errors.New("stream broke")}))
	broken.name = "dest2"
	c := newCheckpointer([]*flushingWriteClient{failing, broken}, nil, serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log}), true, "cq_state")

	err := c.finish(ctx)
	require.ErrorContains(t, err, "destination dest2")
	require.ErrorContains(t, err, "stream broke")
	require.EqualError(t, c.failure(), "destination dest: table items: disk full", "the sync adds the failure to the error")
}

func TestCheckpointerWritesRecordsToBackendWithoutFlushSupport(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log})
	c := newCheckpointer([]*flushingWriteClient{dest}, nil, backend, false, "cq_state")
	c.maxRecordBytes = 1

	require.NoError(t, c.add(stateRecord(t, "cq_state", stateEntry{"a", "1", 1}, stateEntry{"b", "1", 1}, stateEntry{"c", "1", 1})))
	require.NoError(t, c.commit(ctx))
	require.Equal(t, []string{
		"dest: flush 1",
		"backend: insert cq_state a=1",
		"backend: insert cq_state b=1",
		"backend: insert cq_state c=1",
		"backend: close",
	}, log.get(), "the records go in one write, which commits once it closes")
	require.NoError(t, c.close())
}

func TestCheckpointerCountsEveryColumnWhenSplittingState(t *testing.T) {
	ctx := context.Background()
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log}))
	backend := serveFakePlugin(t, &fakeWritePlugin{name: "backend", log: log})
	c := newCheckpointer([]*flushingWriteClient{dest}, nil, backend, true, "cq_state")
	// a row takes 1 + 10 + 16 bytes, plus 8 for each of the 10 columns: two rows fit, but not three
	c.maxRecordBytes = 220

	columns := make(schema.ColumnList, 0, 10)
	columns = append(columns,
		schema.Column{Name: stateKeyColumn, Type: arrow.BinaryTypes.String},
		schema.Column{Name: stateValueColumn, Type: arrow.BinaryTypes.String},
		schema.Column{Name: stateVersionColumn, Type: arrow.PrimitiveTypes.Uint64},
	)
	for i := range 7 {
		columns = append(columns, schema.Column{Name: fmt.Sprintf("extra%d", i), Type: arrow.BinaryTypes.String})
	}
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, (&schema.Table{Name: "cq_state", Columns: columns}).ToArrowSchema())
	defer bldr.Release()
	for _, key := range []string{"a", "b", "c"} {
		bldr.Field(0).(*array.StringBuilder).Append(key)
		bldr.Field(1).(*array.StringBuilder).Append("0123456789")
		bldr.Field(2).(*array.Uint64Builder).Append(1)
		for i := 3; i < len(columns); i++ {
			bldr.Field(i).AppendNull()
		}
	}
	recordBytes, err := plugin.RecordToBytes(bldr.NewRecordBatch())
	require.NoError(t, err)

	require.NoError(t, c.add(recordBytes))
	require.NoError(t, c.commit(ctx))
	require.Equal(t, []string{
		"dest: flush 1",
		"backend: insert cq_state a=0123456789 b=0123456789",
		"backend: insert cq_state c=0123456789",
		"backend: flush 1",
	}, log.get())
	require.NoError(t, c.close())
}

type v3Discovery struct {
	discoveryPb.UnimplementedDiscoveryServer
}

func (v3Discovery) GetVersions(context.Context, *discoveryPb.GetVersions_Request) (*discoveryPb.GetVersions_Response, error) {
	return &discoveryPb.GetVersions_Response{Versions: []int32{3}}, nil
}

func startFakeSource(t *testing.T, source *fakeCheckpointSource) string {
	t.Helper()
	tableBytes, err := plugin.SchemaToBytes(checkpointItemsTable.ToArrowSchema())
	require.NoError(t, err)
	source.responses = append([]*plugin.Sync_Response{
		{Message: &plugin.Sync_Response_MigrateTable{MigrateTable: &plugin.Sync_MessageMigrateTable{Table: tableBytes}}},
	}, source.responses...)
	return startFakeServer(t, func(s *grpc.Server) {
		plugin.RegisterPluginServer(s, source)
		discoveryPb.RegisterDiscoveryServer(s, v3Discovery{})
	})
}

func TestSyncGoesOnWithOtherSourcesAfterFailedFlush(t *testing.T) {
	log := &eventLog{}
	first := &fakeCheckpointSource{responses: []*plugin.Sync_Response{itemsInsert(t, 1, 2), checkpointResponse(t, stateEntry{"cursor", "1", 1})}}
	second := &fakeCheckpointSource{responses: []*plugin.Sync_Response{itemsInsert(t, 3)}}
	destination := &fakeWritePlugin{name: "dest", log: log, supportsFlush: true, flushFailure: "table items: disk full"}
	destinationAddress := startFakeServer(t, func(s *grpc.Server) {
		plugin.RegisterPluginServer(s, destination)
		discoveryPb.RegisterDiscoveryServer(s, v3Discovery{})
	})
	spec := fmt.Sprintf(`kind: source
spec:
  name: first
  registry: grpc
  path: %s
  tables: ["*"]
  destinations: ["dest"]
  backend_options:
    table_name: cq_state
    connection: "@@plugins.dest.connection"
---
kind: source
spec:
  name: second
  registry: grpc
  path: %s
  tables: ["*"]
  destinations: ["dest"]
---
kind: destination
spec:
  name: dest
  registry: grpc
  path: %s
`, startFakeSource(t, first), startFakeSource(t, second), destinationAddress)
	specFile := path.Join(t.TempDir(), "spec.yml")
	require.NoError(t, os.WriteFile(specFile, []byte(spec), 0o600))

	cmd := NewCmdRoot()
	cmd.SetArgs(append([]string{"sync", specFile}, testCommandArgs(t)...))
	err := cmd.Execute()
	require.ErrorAs(t, err, new(*stateNotSavedError))
	require.ErrorContains(t, err, "source first: ")
	require.ErrorContains(t, err, "destination dest: table items: disk full")

	second.mu.Lock()
	defer second.mu.Unlock()
	require.NotNil(t, second.request, "the second source syncs, as the first one finished")
}

func migrateResponse(t *testing.T, table *schema.Table) *plugin.Sync_Response {
	t.Helper()
	tableBytes, err := plugin.SchemaToBytes(table.ToArrowSchema())
	require.NoError(t, err)
	return &plugin.Sync_Response{Message: &plugin.Sync_Response_MigrateTable{MigrateTable: &plugin.Sync_MessageMigrateTable{Table: tableBytes}}}
}

func insertResponse(t *testing.T, table *schema.Table, ids ...int64) *plugin.Sync_Response {
	t.Helper()
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	defer bldr.Release()
	bldr.Field(0).(*array.Int64Builder).AppendValues(ids, nil)
	recordBytes, err := plugin.RecordToBytes(bldr.NewRecordBatch())
	require.NoError(t, err)
	return &plugin.Sync_Response{Message: &plugin.Sync_Response_Insert{Insert: &plugin.Sync_MessageInsert{Record: recordBytes}}}
}

func TestSyncSkipsDeleteStaleOnlyForFailedTables(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	others := &schema.Table{Name: "others", Columns: schema.ColumnList{{Name: "id", Type: arrow.PrimitiveTypes.Int64}}}
	log := &eventLog{}
	destination := &fakeWritePlugin{name: "dest", log: log, supportsFlush: true, flushFailure: "table items: disk full", failedTables: []string{"items"}}
	_, err := syncWithFakeSourceMode(t, &fakeCheckpointSource{}, destination, specs.WriteModeOverwriteDeleteStale,
		migrateResponse(t, others),
		itemsInsert(t, 1, 2),
		insertResponse(t, others, 1),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
	)
	require.ErrorAs(t, err, new(*stateNotSavedError))
	events := log.get()
	require.Contains(t, events, "dest: delete stale others", "the destination committed every record of this table")
	require.NotContains(t, events, "dest: delete stale items")
}

func TestSyncFinishesWhenADestinationEndsItsStreamAfterAFailure(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	log := &eventLog{}
	ending := &fakeWritePlugin{name: "dest", log: log, supportsFlush: true, flushFailure: "table items: disk full", failedTables: []string{"items"}, endAfterFailure: true}
	healthy := &fakeWritePlugin{name: "dest2", log: log, supportsFlush: true}
	err := syncWithFakeDestinations(t, &fakeCheckpointSource{}, []*fakeWritePlugin{ending, healthy}, specs.WriteModeOverwriteDeleteStale,
		itemsInsert(t, 1, 2),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
		itemsInsert(t, 3),
		checkpointResponse(t, stateEntry{"cursor", "2", 2}),
	)
	require.ErrorAs(t, err, new(*stateNotSavedError), "the end of the stream is not a new failure: %v", err)
	require.ErrorContains(t, err, "destination dest: table items: disk full")

	events := log.get()
	require.Contains(t, events, "dest: end")
	require.Contains(t, events, "dest2: delete stale items", "the other destination still runs delete-stale")
	require.Contains(t, events, "dest2: close")
	var inserts int
	for _, event := range events {
		if event == "dest2: insert items" {
			inserts++
		}
	}
	require.Equal(t, 2, inserts, "the sync goes on for the other destination")
}

func TestFlushingWriteClientAfterTheDestinationEndsItsStream(t *testing.T) {
	log := &eventLog{}
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: log, flushFailure: "table items: disk full", failedTables: []string{"items"}, endAfterFailure: true}))
	err := dest.Flush()
	var failed *flushFailedError
	require.ErrorAs(t, err, &failed)
	require.Equal(t, []string{"items"}, failed.tables)

	insert := &plugin.Write_Request{Message: &plugin.Write_Request_Insert{Insert: &plugin.Write_MessageInsert{Record: stateRecord(t, "items", stateEntry{"x", "1", 1})}}}
	require.Eventually(t, func() bool {
		require.NoError(t, dest.Send(insert), "records for the failed destination are dropped")
		dest.mu.Lock()
		defer dest.mu.Unlock()
		return dest.ended
	}, 5*time.Second, 10*time.Millisecond)
	require.ErrorAs(t, dest.Flush(), &failed, "the first failure is final")
	require.Equal(t, []string{"items"}, failed.tables)
	require.True(t, failed.anyTable, "the records sent since may belong to any table")
	require.ErrorContains(t, failed, "table items: disk full; write stream ended: ")
	_, err = dest.CloseAndRecv()
	require.NoError(t, err)
}

func TestFlushingWriteClientReportsStreamErrorsWithoutAFailure(t *testing.T) {
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: &eventLog{}, insertErr: errors.New("disk full")}))
	insert := &plugin.Write_Request{Message: &plugin.Write_Request_Insert{Insert: &plugin.Write_MessageInsert{Record: stateRecord(t, "items", stateEntry{"x", "1", 1})}}}
	var sendErr error
	require.Eventually(t, func() bool {
		sendErr = dest.Send(insert)
		return sendErr != nil
	}, 5*time.Second, 10*time.Millisecond, "a stream that ended without a reported failure fails the sync")
	require.ErrorIs(t, sendErr, io.EOF)
	_, err := dest.CloseAndRecv()
	require.ErrorContains(t, err, "disk full")
}

func TestWaitPreferringGroup(t *testing.T) {
	var eg errgroup.Group
	pipelineErr := errors.New("destination failed")
	received := make(chan struct{})
	eg.Go(func() error {
		// fails only after receive does, as a pipeline whose failure cancelled receive's context would
		<-received
		return pipelineErr
	})
	err := waitPreferringGroup(&eg, func() error {
		defer close(received)
		return context.Canceled
	})
	require.ErrorIs(t, err, pipelineErr)

	require.EqualError(t, waitPreferringGroup(&errgroup.Group{}, func() error { return errors.New("bad record") }), "bad record")
	require.NoError(t, waitPreferringGroup(&errgroup.Group{}, func() error { return nil }))
}

func TestSyncReportsUnsavedStateWhenALaterSourceFails(t *testing.T) {
	first := &fakeCheckpointSource{responses: []*plugin.Sync_Response{itemsInsert(t, 1, 2), checkpointResponse(t, stateEntry{"cursor", "1", 1})}}
	// without backend options, a checkpoint is a hard error
	second := &fakeCheckpointSource{responses: []*plugin.Sync_Response{itemsInsert(t, 3), checkpointResponse(t, stateEntry{"cursor", "1", 1})}}
	destination := &fakeWritePlugin{name: "dest", log: &eventLog{}, supportsFlush: true, flushFailure: "table items: disk full"}
	destinationAddress := startFakeServer(t, func(s *grpc.Server) {
		plugin.RegisterPluginServer(s, destination)
		discoveryPb.RegisterDiscoveryServer(s, v3Discovery{})
	})
	spec := fmt.Sprintf(`kind: source
spec:
  name: first
  registry: grpc
  path: %s
  tables: ["*"]
  destinations: ["dest"]
  backend_options:
    table_name: cq_state
    connection: "@@plugins.dest.connection"
---
kind: source
spec:
  name: second
  registry: grpc
  path: %s
  tables: ["*"]
  destinations: ["dest"]
---
kind: destination
spec:
  name: dest
  registry: grpc
  path: %s
`, startFakeSource(t, first), startFakeSource(t, second), destinationAddress)
	specFile := path.Join(t.TempDir(), "spec.yml")
	require.NoError(t, os.WriteFile(specFile, []byte(spec), 0o600))

	cmd := NewCmdRoot()
	cmd.SetArgs(append([]string{"sync", specFile}, testCommandArgs(t)...))
	err := cmd.Execute()
	require.ErrorContains(t, err, "checkpoints are not enabled for this sync")
	require.ErrorAs(t, err, new(*stateNotSavedError), "the first source's failure is not lost")
	require.ErrorContains(t, err, "source first: ")
}

func TestFlushingWriteClientClosesAfterTheDestinationEndsItsStream(t *testing.T) {
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: &eventLog{}, flushFailure: "table items: disk full", endAfterFailure: true}))
	require.ErrorAs(t, dest.Flush(), new(*flushFailedError))
	_, err := dest.CloseAndRecv()
	require.NoError(t, err, "the sync already fails because of the reported failure")
}

func TestSyncClosesEveryDestinationWhenOneFailsToClose(t *testing.T) {
	log := &eventLog{}
	failing := &fakeWritePlugin{name: "dest", log: log, supportsFlush: true, closeErr: errors.New("commit failed")}
	healthy := &fakeWritePlugin{name: "dest2", log: log, supportsFlush: true}
	err := syncWithFakeDestinations(t, &fakeCheckpointSource{}, []*fakeWritePlugin{failing, healthy}, specs.WriteModeAppend, itemsInsert(t, 1))
	require.ErrorContains(t, err, "commit failed")
	require.Contains(t, log.get(), "dest2: close")
}

func TestSyncSkipsDeleteStaleForTablesThatFailAfterTheFirstFailure(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	others := &schema.Table{Name: "others", Columns: schema.ColumnList{{Name: "id", Type: arrow.PrimitiveTypes.Int64}}}
	unrelated := &schema.Table{Name: "unrelated", Columns: schema.ColumnList{{Name: "id", Type: arrow.PrimitiveTypes.Int64}}}
	log := &eventLog{}
	destination := &fakeWritePlugin{
		name: "dest", log: log, supportsFlush: true,
		flushFailure: "table items: disk full", failedTables: []string{"items"},
		laterFailure: "table others: disk full", laterFailedTables: []string{"others"},
	}
	_, err := syncWithFakeSourceMode(t, &fakeCheckpointSource{}, destination, specs.WriteModeOverwriteDeleteStale,
		migrateResponse(t, others),
		migrateResponse(t, unrelated),
		itemsInsert(t, 1, 2),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
		insertResponse(t, others, 1),
		insertResponse(t, unrelated, 1),
	)
	require.ErrorAs(t, err, new(*stateNotSavedError))
	require.ErrorContains(t, err, "destination dest: table items: disk full; table others: disk full")
	events := log.get()
	require.NotContains(t, events, "dest: delete stale items")
	require.NotContains(t, events, "dest: delete stale others", "others failed after the first failure")
	require.Contains(t, events, "dest: delete stale unrelated")
}

func TestSyncSkipsDeleteStaleForTheRelationTreeOfAFailedTable(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	table := func(name string, parent *schema.Table) *schema.Table {
		return &schema.Table{Name: name, Parent: parent, Columns: schema.ColumnList{{Name: "id", Type: arrow.PrimitiveTypes.Int64}}}
	}
	parent := table("parent", nil)
	child := table("child", parent)
	grandchild := table("grandchild", child)
	unrelated := table("unrelated", nil)
	for _, failed := range []string{"parent", "grandchild"} {
		t.Run(failed, func(t *testing.T) {
			log := &eventLog{}
			destination := &fakeWritePlugin{name: "dest", log: log, supportsFlush: true, flushFailure: "table " + failed + ": disk full", failedTables: []string{failed}}
			_, err := syncWithFakeSourceMode(t, &fakeCheckpointSource{}, destination, specs.WriteModeOverwriteDeleteStale,
				migrateResponse(t, parent),
				migrateResponse(t, child),
				migrateResponse(t, grandchild),
				migrateResponse(t, unrelated),
				insertResponse(t, parent, 1),
				checkpointResponse(t, stateEntry{"cursor", "1", 1}),
			)
			require.ErrorAs(t, err, new(*stateNotSavedError))
			events := log.get()
			for _, name := range []string{"parent", "child", "grandchild"} {
				require.NotContains(t, events, "dest: delete stale "+name, "the whole relation tree keeps its rows")
			}
			require.Contains(t, events, "dest: delete stale unrelated")
		})
	}
}

func TestSyncReportsStreamErrorsAfterAFailure(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	destination := &fakeWritePlugin{
		name: "dest", log: &eventLog{}, supportsFlush: true,
		flushFailure: "table items: disk full", failedTables: []string{"items"},
		closeErr: errors.New("table others: commit failed"),
	}
	_, err := syncWithFakeSourceMode(t, &fakeCheckpointSource{}, destination, specs.WriteModeAppend,
		itemsInsert(t, 1, 2),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
	)
	require.ErrorAs(t, err, new(*stateNotSavedError))
	require.ErrorContains(t, err, "table items: disk full")
	require.ErrorContains(t, err, "table others: commit failed", "a later loss is not hidden behind the first failure")
}

func TestSyncKeepsUnsavedStateWithLaterErrors(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	failing := &fakeWritePlugin{name: "dest", log: &eventLog{}, supportsFlush: true, flushFailure: "table items: disk full"}
	broken := &fakeWritePlugin{name: "dest2", log: &eventLog{}, supportsFlush: true, closeErr: errors.New("commit failed")}
	err := syncWithFakeDestinations(t, &fakeCheckpointSource{}, []*fakeWritePlugin{failing, broken}, specs.WriteModeAppend,
		itemsInsert(t, 1),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
	)
	require.ErrorContains(t, err, "commit failed")
	require.ErrorContains(t, err, "sync stopped without saving its state")
	_, finished := err.(*stateNotSavedError) //nolint:errorlint
	require.False(t, finished, "the later error still fails the sync, and stops the next sources")
}

type namedSource struct {
	name   string
	source *fakeCheckpointSource
	// sends checkpoints, so it needs the state backend
	backend bool
}

// syncSourcesWithCommand runs the sync command for sources that all write to destination.
func syncSourcesWithCommand(t *testing.T, destination *fakeWritePlugin, sources ...namedSource) error {
	t.Helper()
	destinationAddress := startFakeServer(t, func(s *grpc.Server) {
		plugin.RegisterPluginServer(s, destination)
		discoveryPb.RegisterDiscoveryServer(s, v3Discovery{})
	})
	var spec strings.Builder
	for _, source := range sources {
		fmt.Fprintf(&spec, "kind: source\nspec:\n  name: %s\n  registry: grpc\n  path: %s\n  tables: [\"*\"]\n  destinations: [\"%s\"]\n", source.name, startFakeSource(t, source.source), destination.name)
		if source.backend {
			fmt.Fprintf(&spec, "  backend_options:\n    table_name: cq_state_%s\n    connection: \"@@plugins.%s.connection\"\n", source.name, destination.name)
		}
		spec.WriteString("---\n")
	}
	fmt.Fprintf(&spec, "kind: destination\nspec:\n  name: %s\n  registry: grpc\n  path: %s\n", destination.name, destinationAddress)
	specFile := path.Join(t.TempDir(), "spec.yml")
	require.NoError(t, os.WriteFile(specFile, []byte(spec.String()), 0o600))
	cmd := NewCmdRoot()
	cmd.SetArgs(append([]string{"sync", specFile}, testCommandArgs(t)...))
	return cmd.Execute()
}

func synced(source *fakeCheckpointSource) bool {
	source.mu.Lock()
	defer source.mu.Unlock()
	return source.request != nil
}

func TestSyncReportsEverySourceThatSavedNoState(t *testing.T) {
	failing := func(table string) *fakeCheckpointSource {
		return &fakeCheckpointSource{responses: []*plugin.Sync_Response{
			itemsInsert(t, 1),
			{Message: &plugin.Sync_Response_Checkpoint{Checkpoint: &plugin.Sync_MessageCheckpoint{Record: stateRecord(t, table, stateEntry{"cursor", "1", 1})}}},
		}}
	}
	first, second, third := failing("cq_state_first"), failing("cq_state_second"), &fakeCheckpointSource{responses: []*plugin.Sync_Response{itemsInsert(t, 2)}}
	destination := &fakeWritePlugin{name: "dest", log: &eventLog{}, supportsFlush: true, flushFailure: "table items: disk full"}
	err := syncSourcesWithCommand(t, destination,
		namedSource{name: "first", source: first, backend: true},
		namedSource{name: "second", source: second, backend: true},
		namedSource{name: "third", source: third},
	)
	require.ErrorContains(t, err, "source first: ")
	require.ErrorContains(t, err, "source second: ")
	require.True(t, synced(third), "a sync that finished does not stop the next source")
}

func TestSyncStopsAtAHardErrorAfterASourceSavedNoState(t *testing.T) {
	first := &fakeCheckpointSource{responses: []*plugin.Sync_Response{
		itemsInsert(t, 1),
		{Message: &plugin.Sync_Response_Checkpoint{Checkpoint: &plugin.Sync_MessageCheckpoint{Record: stateRecord(t, "cq_state_first", stateEntry{"cursor", "1", 1})}}},
	}}
	// without backend options, a checkpoint is a hard error
	second := &fakeCheckpointSource{responses: []*plugin.Sync_Response{itemsInsert(t, 2), checkpointResponse(t, stateEntry{"cursor", "1", 1})}}
	third := &fakeCheckpointSource{responses: []*plugin.Sync_Response{itemsInsert(t, 3)}}
	destination := &fakeWritePlugin{name: "dest", log: &eventLog{}, supportsFlush: true, flushFailure: "table items: disk full"}
	err := syncSourcesWithCommand(t, destination,
		namedSource{name: "first", source: first, backend: true},
		namedSource{name: "second", source: second},
		namedSource{name: "third", source: third},
	)
	require.ErrorContains(t, err, "checkpoints are not enabled for this sync")
	require.ErrorContains(t, err, "source first: ")
	require.False(t, synced(third), "a hard error stops the sync")
}

func TestSyncStopsAtAHardErrorJoinedWithUnsavedState(t *testing.T) {
	first := &fakeCheckpointSource{responses: []*plugin.Sync_Response{
		itemsInsert(t, 1),
		{Message: &plugin.Sync_Response_Checkpoint{Checkpoint: &plugin.Sync_MessageCheckpoint{Record: stateRecord(t, "cq_state_first", stateEntry{"cursor", "1", 1})}}},
	}}
	second := &fakeCheckpointSource{responses: []*plugin.Sync_Response{itemsInsert(t, 2)}}
	destination := &fakeWritePlugin{name: "dest", log: &eventLog{}, supportsFlush: true, flushFailure: "table items: disk full", closeRPCErr: errors.New("connection lost")}
	err := syncSourcesWithCommand(t, destination,
		namedSource{name: "first", source: first, backend: true},
		namedSource{name: "second", source: second},
	)
	require.ErrorContains(t, err, "connection lost")
	require.ErrorContains(t, err, "sync stopped without saving its state")
	require.False(t, synced(second), "a hard error stops the sync, even when it also saved no state")
}

func TestSyncWarnsAppendModeDestinationsOnlyAboutRowsSyncedAgain(t *testing.T) {
	var logs bytes.Buffer
	logger := zlog.Logger
	zlog.Logger = zerolog.New(&logs)
	t.Cleanup(func() { zlog.Logger = logger })
	const warning = "in append mode, gets them again"

	for _, tc := range []struct {
		name      string
		responses []*plugin.Sync_Response
		warns     bool
	}{
		{name: "checkpoints", responses: []*plugin.Sync_Response{itemsInsert(t, 1), checkpointResponse(t, stateEntry{"cursor", "1", 1})}, warns: true},
		// the source may have saved its state itself, past the lost rows
		{name: "no checkpoints", responses: []*plugin.Sync_Response{itemsInsert(t, 1)}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			logs.Reset()
			destination := &fakeWritePlugin{name: "dest", log: &eventLog{}, supportsFlush: true, flushFailure: "table items: disk full"}
			_, err := syncWithFakeSourceMode(t, &fakeCheckpointSource{}, destination, specs.WriteModeAppend, tc.responses...)
			require.ErrorAs(t, err, new(*stateNotSavedError))
			require.Equal(t, tc.warns, strings.Contains(logs.String(), warning), logs.String())
		})
	}
}

func TestFlushingWriteClientFlushAfterTheStreamEnded(t *testing.T) {
	dest := openFlushingClient(t, serveFakePlugin(t, &fakeWritePlugin{name: "dest", log: &eventLog{}, flushFailure: "table items: disk full", failedTables: []string{"items"}, endAfterFailure: true}))
	var failed *flushFailedError
	require.ErrorAs(t, dest.Flush(), &failed)
	require.False(t, failed.anyTable)

	// the next flush finds the stream ended
	require.ErrorAs(t, dest.Flush(), &failed)
	require.True(t, failed.anyTable, "records of any table may have been lost since")
	require.ErrorContains(t, failed, "table items: disk full; write stream ended: ")
}

func TestSyncSkipsDeleteStaleForEveryTableWhenALaterFailureNamesNone(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	unrelated := &schema.Table{Name: "unrelated", Columns: schema.ColumnList{{Name: "id", Type: arrow.PrimitiveTypes.Int64}}}
	for _, tc := range []struct {
		name              string
		failedTables      []string
		laterFailure      string
		laterFailedTables []string
	}{
		// as the streaming writer reports failed migrations and deletes
		{name: "later failure without tables", failedTables: []string{"items"}, laterFailure: "migrate failed"},
		{name: "table the source did not sync", failedTables: []string{"elsewhere"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			log := &eventLog{}
			destination := &fakeWritePlugin{
				name: "dest", log: log, supportsFlush: true,
				flushFailure: "table items: disk full", failedTables: tc.failedTables,
				laterFailure: tc.laterFailure, laterFailedTables: tc.laterFailedTables,
			}
			_, err := syncWithFakeSourceMode(t, &fakeCheckpointSource{}, destination, specs.WriteModeOverwriteDeleteStale,
				migrateResponse(t, unrelated),
				itemsInsert(t, 1, 2),
				checkpointResponse(t, stateEntry{"cursor", "1", 1}),
				insertResponse(t, unrelated, 1),
			)
			require.ErrorAs(t, err, new(*stateNotSavedError))
			for _, event := range log.get() {
				require.NotContains(t, event, "delete stale")
			}
		})
	}
}

func TestSyncReportsAFailureOnceWithAHardFlushError(t *testing.T) {
	interval := checkpointInterval
	checkpointInterval = 0
	t.Cleanup(func() { checkpointInterval = interval })

	failing := &fakeWritePlugin{name: "dest", log: &eventLog{}, supportsFlush: true, flushFailure: "table items: disk full"}
	broken := &fakeWritePlugin{name: "dest2", log: &eventLog{}, supportsFlush: true, flushErr: errors.New("stream broke")}
	err := syncWithFakeDestinations(t, &fakeCheckpointSource{}, []*fakeWritePlugin{failing, broken}, specs.WriteModeAppend,
		itemsInsert(t, 1),
		checkpointResponse(t, stateEntry{"cursor", "1", 1}),
	)
	require.ErrorContains(t, err, "stream broke")
	require.Equal(t, 1, strings.Count(err.Error(), "table items: disk full"), err.Error())
}

func TestSyncedTablesRoot(t *testing.T) {
	tables := syncedTables{parents: map[string]string{"child": "parent", "grandchild": "child", "a": "b", "b": "a"}}
	require.Equal(t, "parent", tables.root("grandchild"))
	require.Equal(t, "parent", tables.root("parent"))
	require.Equal(t, "unrelated", tables.root("unrelated"))
	require.Contains(t, []string{"a", "b"}, tables.root("a"), "a cycle of parents ends the walk")
}

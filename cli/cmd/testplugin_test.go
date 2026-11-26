package cmd

import (
	"context"
	"fmt"
	"io"
	"net"
	"slices"
	gosync "sync"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// testTable is the one table the fake source plugin serves.
var testTable = &schema.Table{
	Name: "test_table",
	Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "name", Type: arrow.BinaryTypes.String},
	},
}

// pluginLog records what the fake plugins saw, in the order they saw it, so that
// a test can assert on ordering across the source and the destination.
type pluginLog struct {
	mu      gosync.Mutex
	entries []string
}

func (l *pluginLog) add(format string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.entries = append(l.entries, fmt.Sprintf(format, args...))
}

func (l *pluginLog) snapshot() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]string(nil), l.entries...)
}

func (l *pluginLog) count(entry string) int {
	n := 0
	for _, e := range l.snapshot() {
		if e == entry {
			n++
		}
	}
	return n
}

// indexOf returns the position of the first entry equal to entry, or -1.
func (l *pluginLog) indexOf(entry string) int {
	for i, e := range l.snapshot() {
		if e == entry {
			return i
		}
	}
	return -1
}

// lastIndexOfPrefix returns the position of the last entry starting with prefix, or -1.
func (l *pluginLog) lastIndexOfPrefix(prefix string) int {
	last := -1
	for i, e := range l.snapshot() {
		if len(e) >= len(prefix) && e[:len(prefix)] == prefix {
			last = i
		}
	}
	return last
}

// testSourcePlugin is a fake source plugin: it serves a scripted Watch stream and
// answers Sync, both for a full sync and for refresh requests, so that the events
// loop can be driven end to end over real gRPC.
type testSourcePlugin struct {
	plugin.UnimplementedPluginServer

	log *pluginLog

	// ready is what the Watch handshake replies with.
	ready *plugin.Watch_Ready
	// changes are streamed once the client has sent its Position.
	changes []*plugin.Watch_Change
	// watchErr, when set, fails the Watch stream instead of draining it.
	watchErr error
	// fullSyncRows is how many rows the full sync emits.
	fullSyncRows int
	// fullSyncErr, when set, fails the full sync after the rows it did emit.
	fullSyncErr error
	// refreshErrorTable, when set, makes every refresh emit an Error for it.
	refreshErrorTable string
	// refreshFailIDs makes a refresh of these resource ids emit an Error for the
	// table and no row, while the other resources of the request still succeed.
	refreshFailIDs map[string]bool

	mu        gosync.Mutex
	starts    []*plugin.Watch_Start
	positions []*plugin.Watch_Position
	acks      []string
	syncReqs  []*plugin.Sync_Request

	acked chan struct{}
}

func (*testSourcePlugin) Init(context.Context, *plugin.Init_Request) (*plugin.Init_Response, error) {
	return &plugin.Init_Response{}, nil
}

func (*testSourcePlugin) Close(context.Context, *plugin.Close_Request) (*plugin.Close_Response, error) {
	return &plugin.Close_Response{}, nil
}

func (s *testSourcePlugin) syncRequests() []*plugin.Sync_Request {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]*plugin.Sync_Request(nil), s.syncReqs...)
}

func (s *testSourcePlugin) watchStarts() []*plugin.Watch_Start {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]*plugin.Watch_Start(nil), s.starts...)
}

func (s *testSourcePlugin) watchPositions() []*plugin.Watch_Position {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]*plugin.Watch_Position(nil), s.positions...)
}

func (s *testSourcePlugin) ackedEventIDs() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.acks...)
}

func (s *testSourcePlugin) Sync(req *plugin.Sync_Request, stream grpc.ServerStreamingServer[plugin.Sync_Response]) error {
	s.mu.Lock()
	s.syncReqs = append(s.syncReqs, req)
	s.mu.Unlock()

	if len(req.GetRefresh()) == 0 {
		s.log.add("sync:full")
		schemaBytes, err := plugin.SchemaToBytes(testTable.ToArrowSchema())
		if err != nil {
			return err
		}
		if err := stream.Send(&plugin.Sync_Response{Message: &plugin.Sync_Response_MigrateTable{
			MigrateTable: &plugin.Sync_MessageMigrateTable{Table: schemaBytes},
		}}); err != nil {
			return err
		}
		for i := range s.fullSyncRows {
			recordBytes, err := testRecordBytes(fmt.Sprintf("full-%d", i))
			if err != nil {
				return err
			}
			if err := stream.Send(&plugin.Sync_Response{Message: &plugin.Sync_Response_Insert{
				Insert: &plugin.Sync_MessageInsert{Record: recordBytes},
			}}); err != nil {
				return err
			}
		}
		return s.fullSyncErr
	}

	s.log.add("sync:refresh:%d", len(req.GetRefresh()))
	if s.refreshErrorTable != "" {
		if err := stream.Send(&plugin.Sync_Response{Message: &plugin.Sync_Response_Error{
			Error: &plugin.Sync_MessageError{TableName: s.refreshErrorTable, Error: "refresh failed"},
		}}); err != nil {
			return err
		}
	}
	for _, resource := range req.GetRefresh() {
		if resource.GetOp() == plugin.Op_OP_DELETE {
			if err := stream.Send(&plugin.Sync_Response{Message: &plugin.Sync_Response_DeleteRecord{
				DeleteRecord: &plugin.Sync_MessageDeleteRecord{TableName: resource.GetTable()},
			}}); err != nil {
				return err
			}
			continue
		}
		keys, err := plugin.NewRecordFromBytes(resource.GetKeys())
		if err != nil {
			return err
		}
		if s.refreshFailIDs[schema.KeysString(keys)] {
			if err := stream.Send(&plugin.Sync_Response{Message: &plugin.Sync_Response_Error{
				Error: &plugin.Sync_MessageError{TableName: resource.GetTable(), Error: "refresh denied"},
			}}); err != nil {
				return err
			}
			continue
		}
		recordBytes, err := testRecordBytes(schema.KeysString(keys))
		if err != nil {
			return err
		}
		if err := stream.Send(&plugin.Sync_Response{Message: &plugin.Sync_Response_Insert{
			Insert: &plugin.Sync_MessageInsert{Record: recordBytes},
		}}); err != nil {
			return err
		}
	}
	return nil
}

func (s *testSourcePlugin) Watch(stream grpc.BidiStreamingServer[plugin.Watch_Request, plugin.Watch_Response]) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	start := req.GetStart()
	if start == nil {
		return status.Errorf(codes.InvalidArgument, "expected Start, got %T", req.GetMessage())
	}
	s.mu.Lock()
	s.starts = append(s.starts, start)
	s.mu.Unlock()
	s.log.add("watch:start")

	ready := s.ready
	if ready == nil {
		ready = &plugin.Watch_Ready{}
	}
	if err := stream.Send(&plugin.Watch_Response{Message: &plugin.Watch_Response_Ready{Ready: ready}}); err != nil {
		return err
	}

	req, err = stream.Recv()
	if err != nil {
		return err
	}
	position := req.GetPosition()
	if position == nil {
		return status.Errorf(codes.InvalidArgument, "expected Position, got %T", req.GetMessage())
	}
	s.mu.Lock()
	s.positions = append(s.positions, position)
	s.mu.Unlock()
	s.log.add("watch:position")

	go func() {
		for {
			req, err := stream.Recv()
			if err != nil {
				return
			}
			if ack := req.GetAck(); ack != nil {
				s.mu.Lock()
				s.acks = append(s.acks, ack.GetEventId())
				s.mu.Unlock()
				s.log.add("watch:ack:%s", ack.GetEventId())
				select {
				case s.acked <- struct{}{}:
				default:
				}
			}
		}
	}()

	for _, change := range s.changes {
		if err := stream.Send(&plugin.Watch_Response{Message: &plugin.Watch_Response_Change{Change: change}}); err != nil {
			return err
		}
	}

	if s.watchErr != nil {
		return s.watchErr
	}

	if start.GetDrain() {
		if err := stream.Send(&plugin.Watch_Response{Message: &plugin.Watch_Response_Drained{Drained: &plugin.Watch_Drained{}}}); err != nil {
			return err
		}
		// Stay on the stream until the client has acknowledged the last batch, so
		// that the acknowledgement is recorded before the handler returns.
		if len(s.changes) > 0 {
			last := s.changes[len(s.changes)-1].GetEventId()
			deadline := time.After(30 * time.Second)
		wait:
			for !slices.Contains(s.ackedEventIDs(), last) {
				select {
				case <-s.acked:
				case <-stream.Context().Done():
					break wait
				case <-deadline:
					break wait
				}
			}
		}
		return nil
	}

	<-stream.Context().Done()
	return nil
}

// testUnimplementedSourcePlugin answers every RPC the events loop needs except
// Watch, exactly as a plugin built before the Watch RPC does.
type testUnimplementedSourcePlugin struct {
	plugin.UnimplementedPluginServer
}

func (*testUnimplementedSourcePlugin) Init(context.Context, *plugin.Init_Request) (*plugin.Init_Response, error) {
	return &plugin.Init_Response{}, nil
}

func (*testUnimplementedSourcePlugin) Close(context.Context, *plugin.Close_Request) (*plugin.Close_Response, error) {
	return &plugin.Close_Response{}, nil
}

// testDestinationPlugin is a fake destination plugin that records every write.
type testDestinationPlugin struct {
	plugin.UnimplementedPluginServer

	log *pluginLog

	mu       gosync.Mutex
	requests []*plugin.Write_Request
}

func (*testDestinationPlugin) Init(context.Context, *plugin.Init_Request) (*plugin.Init_Response, error) {
	return &plugin.Init_Response{}, nil
}

func (*testDestinationPlugin) Close(context.Context, *plugin.Close_Request) (*plugin.Close_Response, error) {
	return &plugin.Close_Response{}, nil
}

func (d *testDestinationPlugin) writes() []*plugin.Write_Request {
	d.mu.Lock()
	defer d.mu.Unlock()
	return append([]*plugin.Write_Request(nil), d.requests...)
}

func (d *testDestinationPlugin) Write(stream grpc.ClientStreamingServer[plugin.Write_Request, plugin.Write_Response]) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&plugin.Write_Response{})
		}
		if err != nil {
			return err
		}
		d.mu.Lock()
		d.requests = append(d.requests, req)
		d.mu.Unlock()
		switch m := req.GetMessage().(type) {
		case *plugin.Write_Request_MigrateTable:
			d.log.add("write:migrate")
		case *plugin.Write_Request_Insert:
			d.log.add("write:insert")
		case *plugin.Write_Request_DeleteRecord:
			d.log.add("write:delete-record:%s", m.DeleteRecord.GetTableName())
		case *plugin.Write_Request_Delete:
			d.log.add("write:delete-stale:%s", m.Delete.GetTableName())
		}
	}
}

// testRecordBytes builds a one-row record of testTable with the given id.
func testRecordBytes(id string) ([]byte, error) {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, testTable.ToArrowSchema())
	defer bldr.Release()
	bldr.Field(0).(*array.StringBuilder).Append(id)
	bldr.Field(1).(*array.StringBuilder).Append("name-" + id)
	record := bldr.NewRecordBatch()
	defer record.Release()
	return plugin.RecordToBytes(record)
}

// testChangeFor builds a Watch change for one row of testTable.
func testChangeFor(t *testing.T, id string, op plugin.Op, eventID string) *plugin.Watch_Change {
	t.Helper()
	keys, err := schema.NewKeysRecord(testTable, map[string]any{"id": id})
	require.NoError(t, err)
	keyBytes, err := plugin.RecordToBytes(keys)
	require.NoError(t, err)
	return &plugin.Watch_Change{
		Table:   testTable.Name,
		Keys:    keyBytes,
		Op:      op,
		EventId: eventID,
		Time:    timestamppb.Now(),
	}
}

// servePlugin starts srv on a local port and returns a managedplugin client
// connected to it, as `registry: grpc` does for an already-running plugin.
func servePlugin(t *testing.T, typ managedplugin.PluginType, name string, srv plugin.PluginServer) *managedplugin.Client {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	grpcServer := grpc.NewServer()
	plugin.RegisterPluginServer(grpcServer, srv)
	go func() {
		_ = grpcServer.Serve(listener)
	}()
	t.Cleanup(grpcServer.Stop)

	client, err := managedplugin.NewClient(t.Context(), typ, managedplugin.Config{
		Name:     name,
		Registry: managedplugin.RegistryGrpc,
		Path:     listener.Addr().String(),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = client.Terminate()
	})
	return client
}

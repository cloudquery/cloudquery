package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// eventsRunResult is everything an events run left behind for the assertions.
type eventsRunResult struct {
	err             error
	log             *pluginLog
	source          *testSourcePlugin
	destination     *testDestinationPlugin
	summaryLocation string
}

func (r eventsRunResult) summaries(t *testing.T) []syncSummary {
	t.Helper()
	raw, err := os.ReadFile(r.summaryLocation)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	require.NoError(t, err)
	var summaries []syncSummary
	for _, line := range splitLines(raw) {
		var s syncSummary
		require.NoError(t, json.Unmarshal(line, &s))
		summaries = append(summaries, s)
	}
	return summaries
}

func splitLines(b []byte) [][]byte {
	var out [][]byte
	start := 0
	for i, c := range b {
		if c != '\n' {
			continue
		}
		if i > start {
			out = append(out, b[start:i])
		}
		start = i + 1
	}
	if start < len(b) {
		out = append(out, b[start:])
	}
	return out
}

// runEventsSync wires the fake source and destination to syncEventsV3 over real
// gRPC and runs it to completion.
func runEventsSync(ctx context.Context, t *testing.T, src *testSourcePlugin, events *specs.EventsOptions) eventsRunResult {
	t.Helper()

	log := &pluginLog{}
	src.log = log
	src.acked = make(chan struct{}, 1)
	dest := &testDestinationPlugin{log: log}

	sourceClient := servePlugin(t, managedplugin.PluginSource, "test-source", src)
	destinationClient := servePlugin(t, managedplugin.PluginDestination, "test-destination", dest)

	sourceSpec := specs.Source{
		Metadata: specs.Metadata{
			Name:     "test-source",
			Path:     sourceClient.ConnectionString(),
			Registry: specs.RegistryGRPC,
		},
		Tables:       []string{"*"},
		Destinations: []string{"test-destination"},
		BackendOptions: &specs.BackendOptions{
			TableName:  "cq_state_test",
			Connection: "@@plugins.test-destination.connection",
		},
		SyncMode: specs.SyncModeEvents,
		Events:   events,
	}
	sourceSpec.SetDefaults()
	require.NoError(t, sourceSpec.Validate(false))

	destinationSpec := specs.Destination{
		Metadata: specs.Metadata{
			Name:     "test-destination",
			Path:     destinationClient.ConnectionString(),
			Registry: specs.RegistryGRPC,
		},
		WriteMode: specs.WriteModeOverwriteDeleteStale,
	}
	destinationSpec.SetDefaults()

	summaryLocation := filepath.Join(t.TempDir(), "summary.jsonl")
	err := syncEventsV3(ctx, syncV3Options{
		source:                    v3source{client: sourceClient, spec: sourceSpec},
		destinations:              []v3destination{{client: destinationClient, spec: destinationSpec}},
		transformersByDestination: map[string][]v3transformer{},
		uid:                       "test-invocation",
		summaryLocation:           summaryLocation,
	})

	return eventsRunResult{err: err, log: log, source: src, destination: dest, summaryLocation: summaryLocation}
}

func fastEvents(fullSync string) *specs.EventsOptions {
	return &specs.EventsOptions{
		FullSync:        fullSync,
		BatchMaxChanges: 100,
		BatchMaxWait:    duration(50 * time.Millisecond),
		SummaryInterval: duration(time.Hour),
	}
}

func TestSyncEventsV3_FirstRunFullSyncThenRefresh(t *testing.T) {
	src := &testSourcePlugin{
		ready:        &plugin.Watch_Ready{},
		fullSyncRows: 2,
		changes: []*plugin.Watch_Change{
			testChangeFor(t, "a", plugin.Op_OP_UPSERT, "1"),
			testChangeFor(t, "b", plugin.Op_OP_UPSERT, "2"),
		},
	}
	res := runEventsSync(t.Context(), t, src, fastEvents(specs.FullSyncOnFirstRun))
	require.NoError(t, res.err)

	// The full sync ran and stale rows were deleted exactly once, for the one table.
	require.Equal(t, 1, res.log.count("sync:full"))
	require.Equal(t, 1, res.log.count("write:delete-stale:test_table"))

	// Watch was opened before the full sync, with drain set because follow is off.
	starts := res.source.watchStarts()
	require.Len(t, starts, 1)
	require.True(t, starts[0].GetDrain())
	require.Equal(t, []string{"*"}, starts[0].GetTables())
	require.Equal(t, "cq_state_test", starts[0].GetBackend().GetTableName())
	require.Less(t, res.log.indexOf("watch:start"), res.log.indexOf("sync:full"))

	// After a full sync the position resumes from the run start.
	positions := res.source.watchPositions()
	require.Len(t, positions, 1)
	require.NotNil(t, positions[0].Since)

	// Both changes were refreshed in one batch and acknowledged.
	require.Equal(t, 1, res.log.count("sync:refresh:2"))
	require.Equal(t, []string{"2"}, res.source.ackedEventIDs())
}

func TestSyncEventsV3_CatchUpRunDoesNotDeleteStale(t *testing.T) {
	src := &testSourcePlugin{
		ready: &plugin.Watch_Ready{
			HasCheckpoint:  true,
			CheckpointTime: timestamppb.New(time.Now().Add(-time.Minute)),
		},
		fullSyncRows: 2,
		changes: []*plugin.Watch_Change{
			testChangeFor(t, "a", plugin.Op_OP_UPSERT, "1"),
		},
	}
	res := runEventsSync(t.Context(), t, src, fastEvents(specs.FullSyncOnFirstRun))
	require.NoError(t, res.err)

	require.Equal(t, 0, res.log.count("sync:full"))
	require.Equal(t, 0, res.log.lastIndexOfPrefix("write:delete-stale")+1)
	require.Equal(t, 1, res.log.count("sync:refresh:1"))

	// With a checkpoint the source is asked to resume from where it left off.
	positions := res.source.watchPositions()
	require.Len(t, positions, 1)
	require.Nil(t, positions[0].Since)

	// The refresh request carries no table selection, only the changed resources.
	var refreshReqs []*plugin.Sync_Request
	for _, req := range res.source.syncRequests() {
		if len(req.GetRefresh()) > 0 {
			refreshReqs = append(refreshReqs, req)
		}
	}
	require.Len(t, refreshReqs, 1)
	require.Equal(t, "cq_state_test", refreshReqs[0].GetBackend().GetTableName())
}

func TestSyncEventsV3_CoalescesChangesAndAcksTheLastEvent(t *testing.T) {
	src := &testSourcePlugin{
		ready: &plugin.Watch_Ready{HasCheckpoint: true, CheckpointTime: timestamppb.Now()},
		changes: []*plugin.Watch_Change{
			testChangeFor(t, "a", plugin.Op_OP_UPSERT, "1"),
			testChangeFor(t, "a", plugin.Op_OP_UPSERT, "2"),
			testChangeFor(t, "b", plugin.Op_OP_UPSERT, "3"),
			testChangeFor(t, "a", plugin.Op_OP_DELETE, "4"),
		},
	}
	res := runEventsSync(t.Context(), t, src, fastEvents(specs.FullSyncNever))
	require.NoError(t, res.err)

	// Four changes over two resources become two refresh resources, and `a` keeps
	// the last operation seen for it.
	require.Equal(t, 1, res.log.count("sync:refresh:2"))
	var refresh []*plugin.Sync_Resource
	for _, req := range res.source.syncRequests() {
		if len(req.GetRefresh()) > 0 {
			refresh = req.GetRefresh()
		}
	}
	require.Len(t, refresh, 2)
	require.Equal(t, plugin.Op_OP_DELETE, refresh[0].GetOp())
	require.Equal(t, plugin.Op_OP_UPSERT, refresh[1].GetOp())

	// The acknowledged event is the last one received, not the last one kept.
	require.Equal(t, []string{"4"}, res.source.ackedEventIDs())

	// The delete became a DeleteRecord on the destination, not a DeleteStale.
	require.Equal(t, 1, res.log.count("write:delete-record:test_table"))
	require.Equal(t, -1, res.log.lastIndexOfPrefix("write:delete-stale"))

	summaries := res.summaries(t)
	require.Len(t, summaries, 1)
	require.Equal(t, uint64(4), summaries[0].EventsReceived)
	require.Equal(t, uint64(1), summaries[0].EventsRefreshed)
	require.Equal(t, uint64(2), summaries[0].EventsCoalesced)
	require.Equal(t, uint64(1), summaries[0].EventsDeleted)
	require.NotNil(t, summaries[0].IntervalStart)
	require.NotNil(t, summaries[0].IntervalEnd)
	require.NotNil(t, summaries[0].CheckpointTime)
}

func TestSyncEventsV3_AcksOnlyAfterDestinationWrites(t *testing.T) {
	src := &testSourcePlugin{
		ready: &plugin.Watch_Ready{HasCheckpoint: true, CheckpointTime: timestamppb.Now()},
		changes: []*plugin.Watch_Change{
			testChangeFor(t, "a", plugin.Op_OP_UPSERT, "1"),
			testChangeFor(t, "b", plugin.Op_OP_UPSERT, "2"),
		},
	}
	res := runEventsSync(t.Context(), t, src, fastEvents(specs.FullSyncNever))
	require.NoError(t, res.err)

	ackIndex := res.log.indexOf("watch:ack:2")
	require.NotEqual(t, -1, ackIndex, "expected an acknowledgement, got %v", res.log.snapshot())
	lastWrite := res.log.lastIndexOfPrefix("write:insert")
	require.NotEqual(t, -1, lastWrite, "expected destination writes, got %v", res.log.snapshot())
	require.Greater(t, ackIndex, lastWrite, "the batch was acknowledged before it was written: %v", res.log.snapshot())
}

func TestSyncEventsV3_DrainedWithoutChangesWritesASummary(t *testing.T) {
	src := &testSourcePlugin{ready: &plugin.Watch_Ready{HasCheckpoint: true, CheckpointTime: timestamppb.Now()}}
	res := runEventsSync(t.Context(), t, src, fastEvents(specs.FullSyncNever))
	require.NoError(t, res.err)

	require.Empty(t, res.source.ackedEventIDs())
	require.Equal(t, 0, res.log.count("sync:full"))
	summaries := res.summaries(t)
	require.Len(t, summaries, 1)
	require.Equal(t, uint64(0), summaries[0].EventsReceived)
}

func TestSyncEventsV3_WatchErrorFailsWithoutDeleteStale(t *testing.T) {
	src := &testSourcePlugin{
		ready:    &plugin.Watch_Ready{HasCheckpoint: true, CheckpointTime: timestamppb.Now()},
		watchErr: status.Error(codes.Internal, "the change stream broke"),
	}
	res := runEventsSync(t.Context(), t, src, fastEvents(specs.FullSyncNever))
	require.Error(t, res.err)
	require.ErrorContains(t, res.err, "failed to watch source plugin test-source")
	require.ErrorContains(t, res.err, "the change stream broke")
	require.Equal(t, -1, res.log.lastIndexOfPrefix("write:delete-stale"))
	require.Empty(t, res.source.ackedEventIDs())
}

func TestSyncEventsV3_FullSyncFailureSkipsDeleteStale(t *testing.T) {
	src := &testSourcePlugin{
		ready:        &plugin.Watch_Ready{},
		fullSyncRows: 1,
		fullSyncErr:  status.Error(codes.Internal, "the full sync broke"),
	}
	res := runEventsSync(t.Context(), t, src, fastEvents(specs.FullSyncAlways))
	require.Error(t, res.err)
	require.ErrorContains(t, res.err, "the full sync broke")
	require.Equal(t, -1, res.log.lastIndexOfPrefix("write:delete-stale"))
	require.Empty(t, res.source.ackedEventIDs())
}

func TestSyncEventsV3_DeleteStaleRunsOnceBeforeTheStreamFails(t *testing.T) {
	src := &testSourcePlugin{
		ready:        &plugin.Watch_Ready{},
		fullSyncRows: 1,
		watchErr:     status.Error(codes.Internal, "the change stream broke"),
	}
	res := runEventsSync(t.Context(), t, src, fastEvents(specs.FullSyncAlways))
	require.Error(t, res.err)
	// The full sync completed, so its stale rows were deleted once and only once,
	// even though the change stream failed afterwards.
	require.Equal(t, 1, res.log.count("write:delete-stale:test_table"))
}

func TestSyncEventsV3_FollowWritesASummaryPerInterval(t *testing.T) {
	changes := make([]*plugin.Watch_Change, 0, 4)
	for i, id := range []string{"a", "b", "c", "d"} {
		changes = append(changes, testChangeFor(t, id, plugin.Op_OP_UPSERT, string(rune('1'+i))))
	}
	src := &testSourcePlugin{
		ready:             &plugin.Watch_Ready{HasCheckpoint: true, CheckpointTime: timestamppb.Now()},
		changes:           changes,
		refreshErrorTable: testTable.Name,
	}
	events := &specs.EventsOptions{
		FullSync:        specs.FullSyncNever,
		Follow:          true,
		BatchMaxChanges: 1,
		BatchMaxWait:    duration(10 * time.Millisecond),
		SummaryInterval: duration(20 * time.Millisecond),
	}

	ctx, cancel := context.WithCancel(t.Context())
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	res := runEventsSync(ctx, t, src, events)
	require.NoError(t, res.err)

	summaries := res.summaries(t)
	require.Greater(t, len(summaries), 1, "a follow run writes a summary per interval and one on exit")

	// Every change is accounted for exactly once across the intervals, and the
	// intervals are contiguous.
	var received, resources, rowsPerTable, errorsPerTable, sourceErrors uint64
	for i, s := range summaries {
		received += s.EventsReceived
		resources += s.Resources
		rowsPerTable += s.ResourcesPerTable[testTable.Name]
		errorsPerTable += s.ErrorsPerTable[testTable.Name]
		sourceErrors += s.SourceErrors
		require.NotNil(t, s.IntervalStart)
		require.NotNil(t, s.IntervalEnd)
		if i > 0 {
			require.Equal(t, *summaries[i-1].IntervalEnd, *s.IntervalStart)
		}
	}
	require.Equal(t, uint64(len(changes)), received)

	// Every count on a row covers that row's interval only, so the rows add up to
	// the run instead of each repeating everything that came before it.
	require.Equal(t, uint64(len(changes)), resources)
	require.Equal(t, uint64(len(changes)), rowsPerTable)
	require.Equal(t, uint64(len(changes)), errorsPerTable)
	require.Equal(t, uint64(len(changes)), sourceErrors)

	// Concretely: once an interval has reported a table's rows, a later interval
	// with no changes of its own reports none.
	firstWithRows := -1
	for i, s := range summaries {
		if s.ResourcesPerTable[testTable.Name] > 0 {
			firstWithRows = i
			break
		}
	}
	require.NotEqual(t, -1, firstWithRows, "expected an interval carrying the refreshed rows")
	require.Less(t, firstWithRows, len(summaries)-1, "expected at least one interval after the rows landed")
	for _, s := range summaries[firstWithRows+1:] {
		if s.EventsReceived == 0 {
			require.Zero(t, s.ResourcesPerTable[testTable.Name], "an interval with no changes must not repeat an earlier interval's rows")
			require.Zero(t, s.ErrorsPerTable[testTable.Name])
			require.Zero(t, s.SourceErrors)
		}
	}
}

func TestSyncEventsV3_PluginWithoutWatchIsRejected(t *testing.T) {
	log := &pluginLog{}
	sourceClient := servePlugin(t, managedplugin.PluginSource, "legacy-source", &testUnimplementedSourcePlugin{})
	destinationClient := servePlugin(t, managedplugin.PluginDestination, "test-destination", &testDestinationPlugin{log: log})

	sourceSpec := specs.Source{
		Metadata: specs.Metadata{
			Name:     "legacy-source",
			Path:     sourceClient.ConnectionString(),
			Registry: specs.RegistryGRPC,
		},
		Tables:       []string{"*"},
		Destinations: []string{"test-destination"},
		BackendOptions: &specs.BackendOptions{
			TableName:  "cq_state_test",
			Connection: "@@plugins.test-destination.connection",
		},
		SyncMode: specs.SyncModeEvents,
	}
	sourceSpec.SetDefaults()
	destinationSpec := specs.Destination{
		Metadata: specs.Metadata{
			Name:     "test-destination",
			Path:     destinationClient.ConnectionString(),
			Registry: specs.RegistryGRPC,
		},
	}
	destinationSpec.SetDefaults()

	err := syncEventsV3(t.Context(), syncV3Options{
		source:                    v3source{client: sourceClient, spec: sourceSpec},
		destinations:              []v3destination{{client: destinationClient, spec: destinationSpec}},
		transformersByDestination: map[string][]v3transformer{},
		uid:                       "test-invocation",
	})
	require.EqualError(t, err, "source plugin legacy-source does not support sync_mode: events")
}

func TestSyncEventsV3_FollowStopsOnContextCancel(t *testing.T) {
	src := &testSourcePlugin{
		ready: &plugin.Watch_Ready{HasCheckpoint: true, CheckpointTime: timestamppb.Now()},
		changes: []*plugin.Watch_Change{
			testChangeFor(t, "a", plugin.Op_OP_UPSERT, "1"),
		},
	}
	events := fastEvents(specs.FullSyncNever)
	events.Follow = true

	ctx, cancel := context.WithCancel(t.Context())
	go func() {
		// Give the run enough time to refresh and acknowledge the one change.
		time.Sleep(500 * time.Millisecond)
		cancel()
	}()
	res := runEventsSync(ctx, t, src, events)
	require.NoError(t, res.err)

	// A follow run asks the source not to drain.
	starts := res.source.watchStarts()
	require.Len(t, starts, 1)
	require.False(t, starts[0].GetDrain())

	require.Equal(t, []string{"1"}, res.source.ackedEventIDs())
	require.Equal(t, -1, res.log.lastIndexOfPrefix("write:delete-stale"))
	require.Len(t, res.summaries(t), 1)
}

func TestSyncEventsV3_RefreshErrorsAreCounted(t *testing.T) {
	src := &testSourcePlugin{
		ready:             &plugin.Watch_Ready{HasCheckpoint: true, CheckpointTime: timestamppb.Now()},
		refreshErrorTable: testTable.Name,
		changes: []*plugin.Watch_Change{
			testChangeFor(t, "a", plugin.Op_OP_UPSERT, "1"),
		},
	}
	res := runEventsSync(t.Context(), t, src, fastEvents(specs.FullSyncNever))
	require.NoError(t, res.err)

	summaries := res.summaries(t)
	require.Len(t, summaries, 1)
	require.Equal(t, uint64(1), summaries[0].RefreshErrors)
	require.Equal(t, uint64(1), summaries[0].SourceErrors)
}

func TestSyncEventsV3_BatchRecordsCarryTheBatchSyncTime(t *testing.T) {
	src := &testSourcePlugin{
		ready:        &plugin.Watch_Ready{},
		fullSyncRows: 1,
		changes: []*plugin.Watch_Change{
			testChangeFor(t, "a", plugin.Op_OP_UPSERT, "1"),
		},
	}
	res := runEventsSync(t.Context(), t, src, fastEvents(specs.FullSyncAlways))
	require.NoError(t, res.err)

	var syncTimes []time.Time
	for _, req := range res.destination.writes() {
		insert, ok := req.GetMessage().(*plugin.Write_Request_Insert)
		if !ok {
			continue
		}
		record, err := plugin.NewRecordFromBytes(insert.Insert.GetRecord())
		require.NoError(t, err)
		idx := record.Schema().FieldIndices("_cq_sync_time")
		require.Len(t, idx, 1)
		ts, ok := record.Column(idx[0]).(*array.Timestamp)
		require.True(t, ok)
		syncTimes = append(syncTimes, ts.Value(0).ToTime(arrow.Microsecond).UTC())
	}
	require.Len(t, syncTimes, 2, "expected the full sync row and the refreshed row")
	require.True(t, syncTimes[1].After(syncTimes[0]), "the refresh batch must carry its own sync time: %v", syncTimes)
}

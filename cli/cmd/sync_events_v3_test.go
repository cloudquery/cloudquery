package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func duration(d time.Duration) *configtype.Duration {
	v := configtype.NewDuration(d)
	return &v
}

func TestNeedsFullSync(t *testing.T) {
	now := time.Date(2026, 9, 18, 12, 0, 0, 0, time.UTC)
	ready := func(hasCheckpoint bool, age time.Duration) *plugin.Watch_Ready {
		return &plugin.Watch_Ready{
			HasCheckpoint:  hasCheckpoint,
			CheckpointTime: timestamppb.New(now.Add(-age)),
		}
	}

	cases := []struct {
		name   string
		events specs.EventsOptions
		ready  *plugin.Watch_Ready
		want   bool
	}{
		{
			name:   "always without a checkpoint",
			events: specs.EventsOptions{FullSync: specs.FullSyncAlways},
			ready:  ready(false, 0),
			want:   true,
		},
		{
			name:   "always with a fresh checkpoint",
			events: specs.EventsOptions{FullSync: specs.FullSyncAlways, MaxCheckpointAge: duration(20 * time.Hour)},
			ready:  ready(true, time.Minute),
			want:   true,
		},
		{
			name:   "on_first_run without a checkpoint",
			events: specs.EventsOptions{FullSync: specs.FullSyncOnFirstRun},
			ready:  ready(false, 0),
			want:   true,
		},
		{
			name:   "on_first_run with a checkpoint and no max age",
			events: specs.EventsOptions{FullSync: specs.FullSyncOnFirstRun},
			ready:  ready(true, 100*time.Hour),
			want:   false,
		},
		{
			name:   "on_first_run with a checkpoint younger than max age",
			events: specs.EventsOptions{FullSync: specs.FullSyncOnFirstRun, MaxCheckpointAge: duration(20 * time.Hour)},
			ready:  ready(true, 19*time.Hour),
			want:   false,
		},
		{
			name:   "on_first_run with a checkpoint exactly at max age",
			events: specs.EventsOptions{FullSync: specs.FullSyncOnFirstRun, MaxCheckpointAge: duration(20 * time.Hour)},
			ready:  ready(true, 20*time.Hour),
			want:   false,
		},
		{
			name:   "on_first_run with a checkpoint older than max age",
			events: specs.EventsOptions{FullSync: specs.FullSyncOnFirstRun, MaxCheckpointAge: duration(20 * time.Hour)},
			ready:  ready(true, 21*time.Hour),
			want:   true,
		},
		{
			name:   "never without a checkpoint",
			events: specs.EventsOptions{FullSync: specs.FullSyncNever},
			ready:  ready(false, 0),
			want:   false,
		},
		{
			name:   "never without a checkpoint and a max age",
			events: specs.EventsOptions{FullSync: specs.FullSyncNever, MaxCheckpointAge: duration(time.Hour)},
			ready:  ready(false, 0),
			want:   false,
		},
		{
			name:   "never with a checkpoint younger than max age",
			events: specs.EventsOptions{FullSync: specs.FullSyncNever, MaxCheckpointAge: duration(20 * time.Hour)},
			ready:  ready(true, time.Hour),
			want:   false,
		},
		{
			name:   "never with a checkpoint older than max age",
			events: specs.EventsOptions{FullSync: specs.FullSyncNever, MaxCheckpointAge: duration(20 * time.Hour)},
			ready:  ready(true, 21*time.Hour),
			want:   true,
		},
		{
			name:   "never with a checkpoint and no max age",
			events: specs.EventsOptions{FullSync: specs.FullSyncNever},
			ready:  ready(true, 1000*time.Hour),
			want:   false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, needsFullSync(tc.events, tc.ready, now))
		})
	}
}

func TestPositionRequest(t *testing.T) {
	runStart := time.Date(2026, 9, 18, 10, 0, 0, 0, time.UTC)
	now := time.Date(2026, 9, 18, 12, 0, 0, 0, time.UTC)
	withCheckpoint := &plugin.Watch_Ready{HasCheckpoint: true, CheckpointTime: timestamppb.New(runStart)}
	withoutCheckpoint := &plugin.Watch_Ready{}

	t.Run("after a full sync it resumes from the run start", func(t *testing.T) {
		req := positionRequest(true, withoutCheckpoint, specs.EventsOptions{FullSync: specs.FullSyncAlways}, runStart, now)
		require.Equal(t, runStart, req.GetPosition().GetSince().AsTime())
	})

	t.Run("a first run that never syncs fully starts from now", func(t *testing.T) {
		req := positionRequest(false, withoutCheckpoint, specs.EventsOptions{FullSync: specs.FullSyncNever}, runStart, now)
		require.Equal(t, now, req.GetPosition().GetSince().AsTime())
	})

	t.Run("with a checkpoint it resumes from the checkpoint", func(t *testing.T) {
		req := positionRequest(false, withCheckpoint, specs.EventsOptions{FullSync: specs.FullSyncOnFirstRun}, runStart, now)
		require.Nil(t, req.GetPosition().Since)
	})

	t.Run("a first run waiting for events resumes from the checkpoint", func(t *testing.T) {
		req := positionRequest(false, withoutCheckpoint, specs.EventsOptions{FullSync: specs.FullSyncOnFirstRun}, runStart, now)
		require.Nil(t, req.GetPosition().Since)
	})
}

// keysTable is a minimal table with a single `id` primary key, used to build the
// keys records that changes carry.
var keysTable = &schema.Table{
	Name: "test_table",
	Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	},
}

func testChange(t *testing.T, table, id string, op plugin.Op, eventID string) *plugin.Watch_Change {
	t.Helper()
	keys, err := schema.NewKeysRecord(keysTable, map[string]any{"id": id})
	require.NoError(t, err)
	keyBytes, err := plugin.RecordToBytes(keys)
	require.NoError(t, err)
	return &plugin.Watch_Change{Table: table, Keys: keyBytes, Op: op, EventId: eventID}
}

func TestCoalesceChanges(t *testing.T) {
	t.Run("the last operation for a resource wins", func(t *testing.T) {
		changes := []*plugin.Watch_Change{
			testChange(t, "test_table", "a", plugin.Op_OP_UPSERT, "1"),
			testChange(t, "test_table", "b", plugin.Op_OP_UPSERT, "2"),
			testChange(t, "test_table", "a", plugin.Op_OP_DELETE, "3"),
			testChange(t, "test_table", "b", plugin.Op_OP_DELETE, "4"),
			testChange(t, "test_table", "b", plugin.Op_OP_UPSERT, "5"),
		}
		resources, err := coalesceChanges(changes)
		require.NoError(t, err)
		require.Len(t, resources, 2)
		require.Equal(t, "test_table", resources[0].GetTable())
		require.Equal(t, plugin.Op_OP_DELETE, resources[0].GetOp())
		require.Equal(t, plugin.Op_OP_UPSERT, resources[1].GetOp())
		// The acknowledged event is the last one received, not the last one kept.
		require.Equal(t, "5", changes[len(changes)-1].GetEventId())
	})

	t.Run("the same keys on different tables are distinct resources", func(t *testing.T) {
		changes := []*plugin.Watch_Change{
			testChange(t, "table_a", "a", plugin.Op_OP_UPSERT, "1"),
			testChange(t, "table_b", "a", plugin.Op_OP_DELETE, "2"),
		}
		resources, err := coalesceChanges(changes)
		require.NoError(t, err)
		require.Len(t, resources, 2)
		require.Equal(t, "table_a", resources[0].GetTable())
		require.Equal(t, "table_b", resources[1].GetTable())
	})

	t.Run("undecodable keys are an error", func(t *testing.T) {
		_, err := coalesceChanges([]*plugin.Watch_Change{{Table: "test_table", Keys: []byte("not arrow"), EventId: "1"}})
		require.ErrorContains(t, err, "failed to decode the keys of event 1 on table test_table")
	})

	t.Run("an empty batch coalesces to nothing", func(t *testing.T) {
		resources, err := coalesceChanges(nil)
		require.NoError(t, err)
		require.Empty(t, resources)
	})
}

func TestCollectBatch(t *testing.T) {
	change := func(id string) watchMessage {
		return watchMessage{change: &plugin.Watch_Change{EventId: id}}
	}

	t.Run("batch_max_changes closes the batch", func(t *testing.T) {
		in := make(chan watchMessage, 4)
		in <- change("1")
		in <- change("2")
		in <- change("3")
		batch, end, err := collectBatch(context.Background(), in, 2, time.Hour)
		require.NoError(t, err)
		require.Equal(t, batchReady, end)
		require.Len(t, batch, 2)
		require.Equal(t, "1", batch[0].GetEventId())
		require.Equal(t, "2", batch[1].GetEventId())
	})

	t.Run("batch_max_wait closes the batch", func(t *testing.T) {
		in := make(chan watchMessage, 1)
		in <- change("1")
		start := time.Now()
		batch, end, err := collectBatch(context.Background(), in, 100, 20*time.Millisecond)
		require.NoError(t, err)
		require.Equal(t, batchReady, end)
		require.Len(t, batch, 1)
		require.GreaterOrEqual(t, time.Since(start), 20*time.Millisecond)
	})

	t.Run("drained ends the run and keeps the changes collected so far", func(t *testing.T) {
		in := make(chan watchMessage, 2)
		in <- change("1")
		in <- watchMessage{drained: true}
		batch, end, err := collectBatch(context.Background(), in, 100, time.Hour)
		require.NoError(t, err)
		require.Equal(t, batchDrained, end)
		require.Len(t, batch, 1)
	})

	t.Run("a stream error ends the run", func(t *testing.T) {
		in := make(chan watchMessage, 1)
		in <- watchMessage{err: context.DeadlineExceeded}
		batch, end, err := collectBatch(context.Background(), in, 100, time.Hour)
		require.ErrorIs(t, err, context.DeadlineExceeded)
		require.Equal(t, batchFailed, end)
		require.Empty(t, batch)
	})

	t.Run("a closed stream ends the run", func(t *testing.T) {
		in := make(chan watchMessage)
		close(in)
		_, end, err := collectBatch(context.Background(), in, 100, time.Hour)
		require.NoError(t, err)
		require.Equal(t, batchFailed, end)
	})

	t.Run("a cancelled run keeps the in-flight batch", func(t *testing.T) {
		in := make(chan watchMessage, 1)
		in <- change("1")
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(10 * time.Millisecond)
			cancel()
		}()
		batch, end, err := collectBatch(ctx, in, 100, time.Hour)
		require.NoError(t, err)
		require.Equal(t, batchStopped, end)
		require.Len(t, batch, 1)
	})

	t.Run("an empty batch waits for the first change", func(t *testing.T) {
		in := make(chan watchMessage)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
		defer cancel()
		batch, end, err := collectBatch(ctx, in, 100, time.Millisecond)
		require.NoError(t, err)
		require.Equal(t, batchStopped, end)
		require.Empty(t, batch)
	})
}

func TestEventCountersCoalesced(t *testing.T) {
	require.Equal(t, uint64(7), eventCounters{received: 10, refreshed: 3}.coalesced())
	require.Equal(t, uint64(0), eventCounters{received: 0, refreshed: 0}.coalesced())
	// Refreshing more resources than were received cannot happen, but must not underflow.
	require.Equal(t, uint64(0), eventCounters{received: 1, refreshed: 2}.coalesced())
}

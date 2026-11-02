package client

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimestampsStoredInUTCRegardlessOfTimeZone(t *testing.T) {
	ctx := context.Background()
	specBytes, err := json.Marshal(Spec{ConnectionString: "?threads=1"})
	require.NoError(t, err)

	pluginClient, err := New(ctx, zerolog.New(zerolog.NewTestWriter(t)), specBytes, plugin.NewClientOptions{})
	require.NoError(t, err)
	c := pluginClient.(*Client)
	t.Cleanup(func() { require.NoError(t, c.Close(ctx)) })

	require.NoError(t, c.exec(ctx, "SET GLOBAL TimeZone = 'America/New_York'"))

	const sourceName = "test_source"
	syncTime := time.Date(2026, time.July, 1, 12, 0, 0, 0, time.UTC)
	table := &schema.Table{
		Name: "test_timestamps_utc",
		Columns: []schema.Column{
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		},
	}
	require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: table}}))

	bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	defer bldr.Release()
	bldr.Field(0).(*array.StringBuilder).Append(sourceName)
	bldr.Field(1).(*array.TimestampBuilder).Append(arrow.Timestamp(syncTime.UnixMicro()))
	bldr.Field(2).(*array.Int64Builder).Append(1)
	record := bldr.NewRecordBatch()
	defer record.Release()
	require.NoError(t, c.WriteTableBatch(ctx, table.Name, message.WriteInserts{{Record: record}}))

	var storedSyncTime string
	require.NoError(t, c.db.QueryRowContext(ctx, "SELECT strftime(_cq_sync_time, '%Y-%m-%d %H:%M:%S') FROM test_timestamps_utc").Scan(&storedSyncTime))
	assert.Equal(t, syncTime.Format(time.DateTime), storedSyncTime)

	require.NoError(t, c.DeleteStale(ctx, message.WriteDeleteStales{{TableName: table.Name, SourceName: sourceName, SyncTime: syncTime}}))

	var count int64
	require.NoError(t, c.db.QueryRowContext(ctx, "SELECT count(*) FROM test_timestamps_utc").Scan(&count))
	require.Equal(t, int64(1), count)
}

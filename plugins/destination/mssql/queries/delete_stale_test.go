package queries

import (
	"database/sql"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/stretchr/testify/require"
)

func TestDeleteStale(t *testing.T) {
	now := time.Now()
	const (
		schemaName = "cq"
		sourceName = "cq_source"
		expected   = `DELETE FROM [cq].[table_name] WHERE [_cq_source_name] = @sourceName AND [_cq_sync_time] < @syncTime;`
	)

	query, params := DeleteStale(schemaName, &message.WriteDeleteStale{
		TableName:  "table_name",
		SourceName: sourceName,
		SyncTime:   now,
	})

	require.Equal(t, expected, query)
	require.Equal(t, 2, len(params))

	named, ok := params[0].(sql.NamedArg)
	require.True(t, ok)
	require.Equal(t, "sourceName", named.Name)
	require.Equal(t, sourceName, named.Value)

	named, ok = params[1].(sql.NamedArg)
	require.True(t, ok)
	require.Equal(t, "syncTime", named.Name)
	require.Equal(t, now, named.Value)
}

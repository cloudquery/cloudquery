package queries

import (
	"database/sql"
	"testing"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	const (
		schemaName = "cq"
		sourceName = "cq_source"
		expected   = `SELECT
  [_cq_id],
  [_cq_parent_id],
  [_cq_source_name],
  [_cq_sync_time],
  [extra_col]
FROM [cq].[table_name]
WHERE [_cq_source_name] = @sourceName
ORDER BY [_cq_sync_time] ASC;`
	)

	query, params := Read(schemaName, sourceName, &schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{Name: "extra_col", Type: arrow.PrimitiveTypes.Float64},
		},
	})

	require.Equal(t, expected, query)
	require.Equal(t, 1, len(params))

	named, ok := params[0].(sql.NamedArg)
	require.True(t, ok)
	require.Equal(t, "sourceName", named.Name)
	require.Equal(t, sourceName, named.Value)
}

package queries

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `SELECT
  [_cq_id],
  [_cq_parent_id],
  [_cq_source_name],
  [_cq_sync_time],
  [extra_col]
FROM [cq].[table_name];`
	)

	query := Read(schemaName, &schema.Table{
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
}

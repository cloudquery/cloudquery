package queries

import (
	"testing"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	const (
		sourceName = "cq_source"
	)

	query, params := Read(sourceName, &schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{Name: "extra_col", Type: schema.TypeFloat},
		},
	})

	ensureContents(t, query, "read.sql")
	require.Equal(t, 1, len(params))

	named, ok := params[0].(driver.NamedValue)
	require.True(t, ok)
	require.Equal(t, "sourceName", named.Name)
	require.Equal(t, sourceName, named.Value)
}

package queries

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/stretchr/testify/require"
)

func TestAddPK(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `ALTER TABLE [cq].[table_name] ADD CONSTRAINT [table_name_cqpk]
  PRIMARY KEY (
  [_cq_id]
  );`
	)

	query := AddPK(schemaName, &schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{Name: "extra_col", Type: schema.TypeFloat},
		},
	})

	require.Equal(t, expected, query)
}

func TestDropPK(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `ALTER TABLE [cq].[table_name] DROP CONSTRAINT [table_name_cqpk];`
	)

	query := DropPK(schemaName, &schema.Table{Name: "table_name"})

	require.Equal(t, expected, query)
}

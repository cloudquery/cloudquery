package queries

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/stretchr/testify/require"
)

func TestAddColumn(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `ALTER TABLE [cq].[table_name] ADD [my_col] bigint NOT NULL;`
	)

	query := AddColumn(schemaName, schema.CQSchemaToArrow(&schema.Table{Name: "table_name"}), &Definition{
		Name:    "my_col",
		typ:     "bigint",
		notNull: true,
	})

	require.Equal(t, expected, query)
}

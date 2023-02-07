package queries

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/stretchr/testify/require"
)

func TestAddColumn(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `ALTER TABLE [cq].[table_name] ADD [my_col] bigint NOT NULL;`
	)

	query := AddColumn(schemaName, &schema.Table{Name: "table_name"}, &Definition{
		Name:    "my_col",
		typ:     "bigint",
		notNull: true,
	})

	require.Equal(t, expected, query)
}

func TestDropColumn(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `ALTER TABLE [cq].[table_name] DROP COLUMN [my_col];`
	)

	query := DropColumn(schemaName, &schema.Table{Name: "table_name"}, &Definition{Name: "my_col"})

	require.Equal(t, expected, query)
}

func TestAlterColumn(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `ALTER TABLE [cq].[table_name] ALTER COLUMN [my_col] bigint NOT NULL;`
	)

	query := AlterColumn(schemaName,
		&schema.Table{Name: "table_name"},
		&Definition{
			Name:    "my_col",
			typ:     "bigint",
			notNull: true,
		},
	)

	require.Equal(t, expected, query)
}

package queries

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/stretchr/testify/require"
)

func TestAddColumn(t *testing.T) {
	query, err := AddColumn("table_name", "", schema.CQColumnToArrowField(&schema.Column{
		Name:            "my_col",
		Type:            schema.TypeInt,
		CreationOptions: schema.ColumnCreationOptions{NotNull: true},
	}))
	require.NoError(t, err)
	ensureContents(t, query, "col_add.sql")
}

func TestAddColumnCluster(t *testing.T) {
	query, err := AddColumn("table_name", "my_cluster", schema.CQColumnToArrowField(&schema.Column{
		Name:            "my_col",
		Type:            schema.TypeInt,
		CreationOptions: schema.ColumnCreationOptions{NotNull: true},
	}))
	require.NoError(t, err)
	ensureContents(t, query, "col_add_cluster.sql")
}

func TestDropColumn(t *testing.T) {
	query := DropColumn("table_name", "", schema.CQColumnToArrowField(&schema.Column{Name: "my_col", Type: schema.TypeInt}))

	ensureContents(t, query, "col_drop.sql")
}

func TestDropColumnCluster(t *testing.T) {
	query := DropColumn("table_name", "my_cluster", schema.CQColumnToArrowField(&schema.Column{Name: "my_col", Type: schema.TypeInt}))

	ensureContents(t, query, "col_drop_cluster.sql")
}

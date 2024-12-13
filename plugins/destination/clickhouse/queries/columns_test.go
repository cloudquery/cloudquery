package queries

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestAddColumn(t *testing.T) {
	query, err := AddColumn("table_name", "",
		schema.Column{Name: "my_col", Type: arrow.PrimitiveTypes.Int64, NotNull: true})
	require.NoError(t, err)
	ensureContents(t, query, "col_add.sql")
}

func TestAddColumnCluster(t *testing.T) {
	query, err := AddColumn("table_name", "my_cluster",
		schema.Column{Name: "my_col", Type: arrow.PrimitiveTypes.Int64, NotNull: true})
	require.NoError(t, err)
	ensureContents(t, query, "col_add_cluster.sql")
}

func TestDropColumn(t *testing.T) {
	query := DropColumn("table_name", "", schema.Column{Name: "my_col"})

	ensureContents(t, query, "col_drop.sql")
}

func TestDropColumnCluster(t *testing.T) {
	query := DropColumn("table_name", "my_cluster", schema.Column{Name: "my_col"})

	ensureContents(t, query, "col_drop_cluster.sql")
}

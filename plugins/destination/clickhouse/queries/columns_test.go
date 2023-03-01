package queries

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
)

func TestAddColumn(t *testing.T) {
	query := AddColumn("table_name", "", &schema.Column{
		Name:            "my_col",
		Type:            schema.TypeInt,
		CreationOptions: schema.ColumnCreationOptions{NotNull: true},
	})

	ensureContents(t, query, "col_add.sql")
}

func TestAddColumnCluster(t *testing.T) {
	query := AddColumn("table_name", "my_cluster", &schema.Column{
		Name:            "my_col",
		Type:            schema.TypeInt,
		CreationOptions: schema.ColumnCreationOptions{NotNull: true},
	})

	ensureContents(t, query, "col_add_cluster.sql")
}

func TestDropColumn(t *testing.T) {
	query := DropColumn("table_name", "", &schema.Column{Name: "my_col"})

	ensureContents(t, query, "col_drop.sql")
}

func TestDropColumnCluster(t *testing.T) {
	query := DropColumn("table_name", "my_cluster", &schema.Column{Name: "my_col"})

	ensureContents(t, query, "col_drop_cluster.sql")
}

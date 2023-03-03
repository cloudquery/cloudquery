package queries

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
)

func TestCreateTable(t *testing.T) {
	query := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:            "extra_col",
				Type:            schema.TypeFloat,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true},
			},
			schema.Column{Name: "extra_inet_col", Type: schema.TypeInet},
			schema.Column{Name: "extra_inet_arr_col", Type: schema.TypeInetArray},
		},
	}, "")

	ensureContents(t, query, "create_table.sql")
}

func TestCreateTableCluster(t *testing.T) {
	query := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:            "extra_col",
				Type:            schema.TypeFloat,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true},
			},
			schema.Column{Name: "extra_inet_col", Type: schema.TypeInet},
			schema.Column{Name: "extra_inet_arr_col", Type: schema.TypeInetArray},
		},
	}, "my_cluster")

	ensureContents(t, query, "create_table_cluster.sql")
}

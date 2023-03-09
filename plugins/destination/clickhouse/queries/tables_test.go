package queries

import (
	"encoding/json"
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
	}, "", DefaultEngine())

	ensureContents(t, query, "create_table.sql")
}

func TestCreateTableOnCluster(t *testing.T) {
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
	}, "my_cluster", DefaultEngine())

	ensureContents(t, query, "create_table_cluster.sql")
}

func TestCreateTableWithEngine(t *testing.T) {
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
	}, "", &Engine{
		Name:       "ReplicatedMergeTree",
		Parameters: []any{"a", "b", 1, int32(2), int64(3), float32(1.2), float64(3.4), json.Number("327"), false, true},
	})

	ensureContents(t, query, "create_table_engine.sql")
}

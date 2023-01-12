package queries

import (
	"testing"
)

func TestAddColumn(t *testing.T) {
	query := AddColumn("table_name", &ColumnDefinition{
		Name: "my_col",
		Type: "Int64",
	})

	ensureContents(t, query, "col_add.sql")
}

func TestModifyColumn(t *testing.T) {
	query := ModifyColumn("table_name", &ColumnDefinition{
		Name: "my_col",
		Type: "Int64",
	})

	ensureContents(t, query, "col_mod.sql")
}

package queries

import (
	"slices"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

type colQueryBuilder struct {
	Schema string
	Table  string
	Column *schema.Column
}

func AddColumn(schemaName string, table *schema.Table, column *schema.Column) string {
	return execTemplate("col_add.sql.tpl", &colQueryBuilder{
		Schema: schemaName,
		Table:  table.Name,
		Column: column,
	})
}

func UpdateColumnType(schemaName string, table *schema.Table, column *schema.Column) string {
	return execTemplate("col_update_type.sql.tpl", &colQueryBuilder{
		Schema: schemaName,
		Table:  table.Name,
		Column: column,
	})
}

func GetValueColumns(table *schema.Table) []string {
	columns := make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		if !col.PrimaryKey {
			columns = append(columns, col.Name)
		}
	}

	return slices.Clip(columns)
}

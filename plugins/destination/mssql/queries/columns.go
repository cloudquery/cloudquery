package queries

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

type colQueryBuilder struct {
	Table      string
	Definition *Definition
}

func AddColumn(schemaName string, table *schema.Table, definition *Definition) string {
	return execTemplate("col_add.sql.tpl", &colQueryBuilder{
		Table:      SanitizedTableName(schemaName, table),
		Definition: definition.sanitized(),
	})
}

func GetPKColumns(table *schema.Table) []string {
	return sanitized(table.PrimaryKeys()...)
}

func GetValueColumns(columns schema.ColumnList) []string {
	var cols []string

	for _, col := range columns {
		if !col.CreationOptions.PrimaryKey {
			cols = append(cols, col.Name)
		}
	}

	return sanitized(cols...)
}

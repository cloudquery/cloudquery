package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
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

func DropColumn(schemaName string, table *schema.Table, definition *Definition) string {
	return execTemplate("col_drop.sql.tpl", &colQueryBuilder{
		Table:      SanitizedTableName(schemaName, table),
		Definition: definition.sanitized(),
	})
}

func AlterColumn(schemaName string, table *schema.Table, definition *Definition) string {
	return execTemplate("col_alter.sql.tpl", &colQueryBuilder{
		Table:      SanitizedTableName(schemaName, table),
		Definition: definition.sanitized(),
	})
}

func GetPKColumns(table *schema.Table, enabled bool) []string {
	pk := table.PrimaryKeys()

	if !enabled || len(pk) == 0 {
		return sanitized(schema.CqIDColumn.Name)
	}

	return sanitized(pk...)
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

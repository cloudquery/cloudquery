package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type createTableQueryBuilder struct {
	Table       string
	Definitions Definitions
	PrimaryKey  []string
}

func CreateTable(schemaName string, pkEnabled bool, table *schema.Table) string {
	return execTemplate("create_table.sql.tpl", &createTableQueryBuilder{
		Table:       SanitizeID(schemaName, table.Name),
		Definitions: GetDefinitions(table.Columns, pkEnabled),
		PrimaryKey:  GetPKColumns(table, pkEnabled),
	})
}

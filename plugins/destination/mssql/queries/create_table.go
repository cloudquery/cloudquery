package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type (
	createTableQueryBuilder struct {
		Table       string
		Definitions Definitions
		PrimaryKey  *pkQueryBuilder
	}
)

func CreateTable(schemaName string, table *schema.Table, pkEnabled bool) string {
	return execTemplate("create_table.sql.tpl", &createTableQueryBuilder{
		Table:       SanitizedTableName(schemaName, table),
		Definitions: GetDefinitions(table.Columns, pkEnabled),
		PrimaryKey: &pkQueryBuilder{
			Name:    pkConstraint(table),
			Columns: GetPKColumns(table, pkEnabled),
		},
	})
}

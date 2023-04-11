package queries

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
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
			Columns: GetPKColumns(table),
		},
	})
}

func DropTable(schemaName string, table *schema.Table) string {
	return execTemplate("drop_table.sql.tpl", &createTableQueryBuilder{
		Table: SanitizedTableName(schemaName, table),
	})
}

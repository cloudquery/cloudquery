package queries

import (
	"github.com/apache/arrow/go/v12/arrow"
)

type (
	createTableQueryBuilder struct {
		Table       string
		Definitions Definitions
		PrimaryKey  *pkQueryBuilder
	}
)

func CreateTable(schemaName string, sc *arrow.Schema, pkEnabled bool) string {
	return execTemplate("create_table.sql.tpl", &createTableQueryBuilder{
		Table:       SanitizedTableName(schemaName, sc),
		Definitions: GetDefinitions(sc, pkEnabled),
		PrimaryKey: &pkQueryBuilder{
			Name:    pkConstraint(sc),
			Columns: GetPKColumns(sc),
		},
	})
}

func DropTable(schemaName string, sc *arrow.Schema) string {
	return execTemplate("drop_table.sql.tpl", &createTableQueryBuilder{
		Table: SanitizedTableName(schemaName, sc),
	})
}

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
	builder := &createTableQueryBuilder{
		Table:       SanitizedTableName(schemaName, sc),
		Definitions: GetDefinitions(sc, pkEnabled),
		PrimaryKey: &pkQueryBuilder{
			Table:   SanitizedTableName(schemaName, sc),
			Name:    pkConstraint(sc),
			Columns: GetPKColumns(sc),
		},
	}

	if len(builder.PrimaryKey.Columns) == 0 {
		builder.PrimaryKey = nil
	}

	return execTemplate("create_table.sql.tpl", builder)
}

func DropTable(schemaName string, sc *arrow.Schema) string {
	return execTemplate("drop_table.sql.tpl", &createTableQueryBuilder{
		Table: SanitizedTableName(schemaName, sc),
	})
}

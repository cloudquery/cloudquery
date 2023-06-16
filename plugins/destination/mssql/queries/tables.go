package queries

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

type (
	createTableQueryBuilder struct {
		Schema     string
		Table      string
		Columns    schema.ColumnList
		PrimaryKey *pkQueryBuilder
	}
)

func CreateTable(schemaName string, table *schema.Table, pkEnabled bool) string {
	builder := &createTableQueryBuilder{
		Schema:  schemaName,
		Table:   table.Name,
		Columns: table.Columns,
		PrimaryKey: &pkQueryBuilder{
			Schema:  schemaName,
			Table:   table.Name,
			Name:    pkConstraint(table),
			Columns: table.PrimaryKeys(),
		},
	}

	if len(builder.PrimaryKey.Columns) == 0 {
		builder.PrimaryKey = nil
	}

	return execTemplate("create_table.sql.tpl", builder)
}

func DropTable(schemaName string, table *schema.Table) string {
	return execTemplate("drop_table.sql.tpl", &createTableQueryBuilder{
		Schema: schemaName,
		Table:  table.Name,
	})
}

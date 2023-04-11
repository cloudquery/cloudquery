package queries

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

type pkQueryBuilder struct {
	Table   string
	Name    string // constraint name
	Columns []string
}

func pkConstraint(table *schema.Table) string {
	const pkSuffix = "_cqpk"
	return sanitizeID(table.Name + pkSuffix)
}

// AddPK should be called only for mode with PK enabled.
func AddPK(schemaName string, table *schema.Table) string {
	return execTemplate("pk_add.sql.tpl", &pkQueryBuilder{
		Table:   sanitizeID(schemaName, table.Name),
		Name:    pkConstraint(table),
		Columns: GetPKColumns(table),
	})
}

// DropPK should be called only for mode with PK enabled.
func DropPK(schemaName string, table *schema.Table) string {
	return execTemplate("pk_drop.sql.tpl", &pkQueryBuilder{
		Table: sanitizeID(schemaName, table.Name),
		Name:  pkConstraint(table),
	})
}

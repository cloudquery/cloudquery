package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type pkQueryBuilder struct {
	Table   string
	Name    string // constraint name
	Columns []string
}

const (
	pkSuffix = "_cqpk"
)

// AddPK should be called only for mode with PK enabled.
func AddPK(schemaName string, table *schema.Table) string {
	return execTemplate("pk_add.sql.tpl", &pkQueryBuilder{
		Table:   SanitizeID(schemaName, table.Name),
		Name:    SanitizeID(table.Name + pkSuffix),
		Columns: GetPKColumns(table, true), // we call AddPK only for enabled
	})
}

// DropPK should be called only for mode with PK enabled.
func DropPK(schemaName string, table *schema.Table) string {
	return execTemplate("pk_drop.sql.tpl", &pkQueryBuilder{
		Table: SanitizeID(schemaName, table.Name),
		Name:  SanitizeID(table.Name + pkSuffix),
	})
}

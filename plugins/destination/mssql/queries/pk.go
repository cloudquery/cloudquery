package queries

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

type pkQueryBuilder struct {
	Table   string
	Name    string // constraint name
	Columns []string
}

func pkConstraint(sc *arrow.Schema) string {
	const pkSuffix = "_cqpk"
	return sanitizeID(schema.TableName(sc) + pkSuffix)
}

// AddPK should be called only for mode with PK enabled.
func AddPK(schemaName string, sc *arrow.Schema) string {
	return execTemplate("pk_add.sql.tpl", &pkQueryBuilder{
		Table:   SanitizedTableName(schemaName, sc),
		Name:    pkConstraint(sc),
		Columns: GetPKColumns(sc),
	})
}

// DropPK should be called only for mode with PK enabled.
func DropPK(schemaName string, sc *arrow.Schema) string {
	return execTemplate("pk_drop.sql.tpl", &pkQueryBuilder{
		Table: SanitizedTableName(schemaName, sc),
		Name:  pkConstraint(sc),
	})
}

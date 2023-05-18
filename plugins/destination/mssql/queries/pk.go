package queries

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

type pkQueryBuilder struct {
	Schema  string
	Table   string
	Name    string // constraint name
	Columns []string
}

func pkConstraint(table *schema.Table) string {
	if len(table.PkConstraintName) > 0 {
		return sanitizeID(table.PkConstraintName)
	}

	const pkSuffix = "_cqpk"
	return sanitizeID(table.Name + pkSuffix)
}

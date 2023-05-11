package queries

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

type pkQueryBuilder struct {
	Table   string
	Name    string // constraint name
	Columns []string
}

func pkConstraint(sc *arrow.Schema) string {
	if constraintName, ok := sc.Metadata().GetValue(schema.MetadataConstraintName); ok {
		return sanitizeID(constraintName)
	}

	const pkSuffix = "_cqpk"
	return sanitizeID(schema.TableName(sc) + pkSuffix)
}

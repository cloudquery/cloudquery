package types

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
)

// https://clickhouse.com/docs/en/sql-reference/data-types/tuple
func structType(_struct *arrow.StructType) string {
	return "Tuple(" + strings.Join(definitions(_struct.Fields()), ", ") + ")"
}

func definitions(fields []arrow.Field) []string {
	res := make([]string, len(fields))
	for i, field := range fields {
		res[i] = FieldDefinition(field)
	}
	return res
}

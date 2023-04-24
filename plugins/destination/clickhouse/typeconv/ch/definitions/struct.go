package definitions

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
)

func definitions(fields ...arrow.Field) []string {
	res := make([]string, len(fields))
	for i, field := range fields {
		res[i] = FieldDefinition(field)
	}
	return res
}

func structType(_struct *arrow.StructType) string {
	// https://clickhouse.com/docs/en/sql-reference/data-types/tuple
	return "Tuple(" + strings.Join(definitions(_struct.Fields()...), ", ") + ")"
}

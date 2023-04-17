package queries

import (
	"github.com/apache/arrow/go/v12/arrow"
)

func fieldsDefinitions(fields []arrow.Field) []string {
	res := make([]string, len(fields))
	for i, field := range fields {
		res[i] = sanitizeID(field.Name) + " " + chFieldType(field)
	}
	return res
}

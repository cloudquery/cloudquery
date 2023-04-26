package queries

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func SanitizedTableName(schemaName string, sc *arrow.Schema) string {
	return sanitizeID(schemaName, schema.TableName(sc))
}

func sanitizeID(parts ...string) string {
	for i, pt := range parts {
		parts[i] = `[` + strings.ReplaceAll(pt, string([]byte{0}), ``) + `]`
	}
	return strings.Join(parts, ".")
}

func sanitized(elems ...string) []string {
	result := make([]string, len(elems))
	for i, column := range elems {
		result[i] = sanitizeID(column)
	}
	return result
}

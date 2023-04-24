package queries

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func SanitizedTableName(schemaName string, table *schema.Table) string {
	return sanitizeID(schemaName, table.Name)
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

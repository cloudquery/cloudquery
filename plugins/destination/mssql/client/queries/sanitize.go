package queries

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func SanitizeID(parts ...string) string {
	for i, pt := range parts {
		parts[i] = `[` + strings.ReplaceAll(pt, string([]byte{0}), ``) + `]`
	}
	return strings.Join(parts, ".")
}

func SanitizedTableName(schemaName string, table *schema.Table) string {
	return SanitizeID(schemaName, table.Name)
}

func Sanitized(elems ...string) []string {
	result := make([]string, len(elems))
	for i, column := range elems {
		result[i] = SanitizeID(column)
	}
	return result
}

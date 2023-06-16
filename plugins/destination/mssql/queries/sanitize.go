package queries

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v3/schema"
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

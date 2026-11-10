package queries

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func SanitizedTableName(schemaName string, table *schema.Table) string {
	return sanitizeID(schemaName, table.Name)
}

func sanitizeID(parts ...string) string {
	for i, pt := range parts {
		withoutNUL := strings.ReplaceAll(pt, string([]byte{0}), ``)
		parts[i] = `[` + strings.ReplaceAll(withoutNUL, `]`, `]]`) + `]`
	}
	return strings.Join(parts, ".")
}

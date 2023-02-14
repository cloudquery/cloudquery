package queries

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func Insert(table *schema.Table) string {
	return `INSERT INTO ` + sanitizeID(table.Name) +
		`(` + strings.Join(sanitized(table.Columns.Names()...), `, `) + `)`
}

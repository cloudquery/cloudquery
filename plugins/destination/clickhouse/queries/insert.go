package queries

import (
	"slices"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/util"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func Insert(table *schema.Table) string {
	insert := `INSERT INTO ` + util.SanitizeID(table.Name)
	columns := table.Columns.Names()
	if slices.ContainsFunc(columns, driverCannotParseColumnName) {
		return insert
	}
	return insert + ` (` + strings.Join(util.Sanitized(columns...), `, `) + `)`
}

func driverCannotParseColumnName(name string) bool {
	return strings.ContainsAny(name, "\\`()")
}

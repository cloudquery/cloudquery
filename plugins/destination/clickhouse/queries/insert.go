package queries

import (
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func Insert(table *schema.Table) string {
	return `INSERT INTO ` + util.SanitizeID(table.Name) +
		`(` + strings.Join(util.Sanitized(table.Columns.Names()...), `, `) + `)`
}

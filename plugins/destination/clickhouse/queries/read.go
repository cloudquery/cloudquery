package queries

import (
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/util"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func Read(table *schema.Table) string {
	return "SELECT " + strings.Join(util.Sanitized(table.Columns.Names()...), ", ") +
		` FROM ` + util.SanitizeID(table.Name)
}

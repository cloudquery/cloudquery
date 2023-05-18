package queries

import (
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func Read(sourceName string, table *schema.Table) (query string, params []any) {
	return "SELECT " + strings.Join(util.Sanitized(table.Columns.Names()...), ", ") + `
FROM ` + util.SanitizeID(table.Name) + `
WHERE ` + util.SanitizeID(schema.CqSourceNameColumn.Name) + ` = @sourceName
ORDER BY ` + util.SanitizeID(schema.CqSyncTimeColumn.Name),
		[]any{driver.NamedValue{Name: "sourceName", Value: sourceName}}
}

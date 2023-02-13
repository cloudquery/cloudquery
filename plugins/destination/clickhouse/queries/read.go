package queries

import (
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Read(sourceName string, table *schema.Table) (query string, params []any) {
	return "SELECT " + strings.Join(sanitized(table.Columns.Names()...), ", ") + `
FROM ` + sanitizeID(table.Name) + `
WHERE ` + sanitizeID(schema.CqSourceNameColumn.Name) + ` = @sourceName
ORDER BY ` + sanitizeID(schema.CqSyncTimeColumn.Name),
		[]any{driver.NamedValue{Name: "sourceName", Value: sourceName}}
}

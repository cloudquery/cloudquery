package queries

import (
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func Read(sourceName string, table *arrow.Schema) (query string, params []any) {
	return "SELECT " + strings.Join(sanitized(ColumnNames(table.Fields())...), ", ") + `
FROM ` + sanitizeID(schema.TableName(table)) + `
WHERE ` + sanitizeID(schema.CqSourceNameColumn.Name) + ` = @sourceName
ORDER BY ` + sanitizeID(schema.CqSyncTimeColumn.Name),
		[]any{driver.NamedValue{Name: "sourceName", Value: sourceName}}
}

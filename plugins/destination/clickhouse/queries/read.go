package queries

import (
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func Read(sourceName string, sc *arrow.Schema) (query string, params []any) {
	return "SELECT " + strings.Join(util.Sanitized(ColumnNames(sc.Fields())...), ", ") + `
FROM ` + util.SanitizeID(schema.TableName(sc)) + `
WHERE ` + util.SanitizeID(schema.CqSourceNameColumn.Name) + ` = @sourceName
ORDER BY ` + util.SanitizeID(schema.CqSyncTimeColumn.Name),
		[]any{driver.NamedValue{Name: "sourceName", Value: sourceName}}
}

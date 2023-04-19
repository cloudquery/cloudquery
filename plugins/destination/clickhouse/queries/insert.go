package queries

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func Insert(table *arrow.Schema) string {
	return `INSERT INTO ` + util.SanitizeID(schema.TableName(table)) +
		`(` + strings.Join(util.Sanitized(ColumnNames(table.Fields())...), `, `) + `)`
}

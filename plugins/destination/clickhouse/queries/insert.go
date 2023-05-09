package queries

import (
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func Insert(sc *arrow.Schema) string {
	return `INSERT INTO ` + util.SanitizeID(schema.TableName(sc)) +
		`(` + strings.Join(util.Sanitized(ColumnNames(sc.Fields())...), `, `) + `)`
}

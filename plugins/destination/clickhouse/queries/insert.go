package queries

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func Insert(table *arrow.Schema) string {
	return `INSERT INTO ` + sanitizeID(schema.TableName(table)) +
		`(` + strings.Join(sanitized(ColumnNames(table.Fields())...), `, `) + `)`
}

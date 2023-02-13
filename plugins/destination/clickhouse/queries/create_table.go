package queries

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func CreateTable(table *schema.Table) string {
	normalized := normalizeTable(table)
	return "CREATE TABLE " + sanitizeID(normalized.Name) + ` (
  ` + strings.Join(definitions(normalized.Columns), `,
  `) + `
) ENGINE = MergeTree ORDER BY (` + sanitizeID(schema.CqIDColumn.Name) + `)`
}

func definitions(columns schema.ColumnList) []string {
	res := make([]string, len(columns))

	for i, col := range columns {
		res[i] = sanitizeID(col.Name) + " " + chType(&col)
	}

	return res
}

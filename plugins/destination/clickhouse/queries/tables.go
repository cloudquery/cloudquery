package queries

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func sortKeys(table *schema.Table) []string {
	keys := make([]string, 0)
	for _, col := range table.Columns {
		if col.CreationOptions.NotNull {
			keys = append(keys, col.Name)
		}
	}
	return keys
}

func CreateTable(table *schema.Table, cluster string) string {
	normalized := normalizeTable(table)
	strBuilder := strings.Builder{}
	strBuilder.WriteString("CREATE TABLE ")
	strBuilder.WriteString(tableNamePart(normalized.Name, cluster))
	strBuilder.WriteString(" (\n")
	for i, col := range normalized.Columns {
		strBuilder.WriteString("  ")
		strBuilder.WriteString(sanitizeID(col.Name))
		strBuilder.WriteString(" ")
		strBuilder.WriteString(chType(&col))
		if i < len(normalized.Columns)-1 {
			strBuilder.WriteString(",")
		}
		strBuilder.WriteString("\n")
	}
	strBuilder.WriteString(") ENGINE = MergeTree ORDER BY (")
	sortingKeys := sanitized(sortKeys(normalized)...)
	strBuilder.WriteString(strings.Join(sortingKeys, ", "))
	strBuilder.WriteString(")")

	return strBuilder.String()
}

func DropTable(table *schema.Table) string {
	return "DROP TABLE IF EXISTS " + sanitizeID(table.Name)
}

package queries

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func sortKeys(table *arrow.Schema) []string {
	keys := make([]string, 0)
	for _, field := range table.Fields() {
		if !field.Nullable {
			keys = append(keys, field.Name)
		}
	}
	return keys
}

func CreateTable(table *arrow.Schema, cluster string, engine *Engine) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("CREATE TABLE ")
	strBuilder.WriteString(tableNamePart(schema.TableName(table), cluster))
	strBuilder.WriteString(" (\n")
	strBuilder.WriteString("  ")
	strBuilder.WriteString(strings.Join(fieldsDefinitions(table.Fields()), "\n  "))
	strBuilder.WriteString("\n) ENGINE = ")
	strBuilder.WriteString(engine.String())
	strBuilder.WriteString(" ORDER BY (")
	sortingKeys := sanitized(sortKeys(table)...)
	strBuilder.WriteString(strings.Join(sortingKeys, ", "))
	strBuilder.WriteString(")")

	return strBuilder.String()
}

func DropTable(table *arrow.Schema, cluster string) string {
	return "DROP TABLE IF EXISTS " + tableNamePart(schema.TableName(table), cluster)
}

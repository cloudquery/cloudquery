package queries

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func sortKeys(sc *arrow.Schema) []string {
	keys := make([]string, 0)
	for _, field := range sc.Fields() {
		if !field.Nullable {
			keys = append(keys, field.Name)
		}
	}
	return keys
}

func CreateTable(sc *arrow.Schema, cluster string, engine *Engine) (string, error) {
	definitions, err := typeconv.ClickHouseDefinitions(sc.Fields()...)
	if err != nil {
		return "", err
	}
	strBuilder := strings.Builder{}
	strBuilder.WriteString("CREATE TABLE ")
	strBuilder.WriteString(tableNamePart(schema.TableName(sc), cluster))
	strBuilder.WriteString(" (\n")
	strBuilder.WriteString("  ")
	strBuilder.WriteString(strings.Join(definitions, ",\n  "))
	strBuilder.WriteString("\n) ENGINE = ")
	strBuilder.WriteString(engine.String())
	strBuilder.WriteString(" ORDER BY (")
	sortingKeys := util.Sanitized(sortKeys(sc)...)
	strBuilder.WriteString(strings.Join(sortingKeys, ", "))
	strBuilder.WriteString(")")

	return strBuilder.String(), nil
}

func DropTable(sc *arrow.Schema, cluster string) string {
	return "DROP TABLE IF EXISTS " + tableNamePart(schema.TableName(sc), cluster)
}

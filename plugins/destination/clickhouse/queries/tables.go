package queries

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	_clickhouse "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"golang.org/x/exp/slices"
)

func sortKeys(sc *arrow.Schema) []string {
	keys := make([]string, 0, len(sc.Fields()))
	for _, field := range sc.Fields() {
		if !field.Nullable {
			keys = append(keys, field.Name)
		}
	}

	// check for _cq_id
	if idx := slices.Index(keys, schema.CqIDColumn.Name); idx >= 0 {
		// move to back, as _cq_id is better suited there
		keys = append(slices.Delete(keys, idx, idx+1), schema.CqIDColumn.Name)
	}

	return slices.Clip(keys)
}

func CreateTable(sc *arrow.Schema, cluster string, engine *Engine) (string, error) {
	definitions, err := _clickhouse.Definitions(sc.Fields()...)
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
	if orderBy := sortKeys(sc); len(orderBy) > 0 {
		strBuilder.WriteString(" ORDER BY (")
		strBuilder.WriteString(strings.Join(util.Sanitized(orderBy...), ", "))
		strBuilder.WriteString(")")
	}

	return strBuilder.String(), nil
}

func DropTable(sc *arrow.Schema, cluster string) string {
	return "DROP TABLE IF EXISTS " + tableNamePart(schema.TableName(sc), cluster)
}

package queries

import (
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/types"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v3/schema"
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
	builder := strings.Builder{}
	builder.WriteString("CREATE TABLE ")
	builder.WriteString(tableNamePart(schema.TableName(sc), cluster))
	builder.WriteString(" (\n")
	builder.WriteString("  ")
	fields := sc.Fields()
	for i, field := range fields {
		definition, err := types.FieldDefinition(field)
		if err != nil {
			return "", err
		}
		builder.WriteString(definition)
		if i < len(fields)-1 {
			builder.WriteString(",\n  ")
		}
	}
	builder.WriteString("\n) ENGINE = ")
	builder.WriteString(engine.String())
	if orderBy := sortKeys(sc); len(orderBy) > 0 {
		builder.WriteString(" ORDER BY (")
		builder.WriteString(strings.Join(util.Sanitized(orderBy...), ", "))
		builder.WriteString(")")
	}
	builder.WriteString(" SETTINGS allow_nullable_key=1") // allows nullable keys

	return builder.String(), nil
}

func DropTable(sc *arrow.Schema, cluster string) string {
	return "DROP TABLE IF EXISTS " + tableNamePart(schema.TableName(sc), cluster)
}

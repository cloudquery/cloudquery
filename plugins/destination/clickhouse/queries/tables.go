package queries

import (
	"slices"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/client/spec"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/types"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const cqSyncGroupID = "_cq_sync_group_id"

func sortKeys(table *schema.Table) []string {
	keys := make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		if col.NotNull || col.PrimaryKey {
			keys = append(keys, col.Name)
		}
	}

	// check for _cq_id
	if idx := slices.Index(keys, schema.CqIDColumn.Name); idx >= 0 {
		// move to back, as _cq_id is better suited there
		keys = append(slices.Delete(keys, idx, idx+1), schema.CqIDColumn.Name)
	}

	return slices.Clip(keys)
}

func CreateTable(table *schema.Table, cluster string, engine *spec.Engine, partitionBySyncGroupID bool) (string, error) {
	builder := strings.Builder{}
	builder.WriteString("CREATE TABLE ")
	builder.WriteString(tableNamePart(table.Name, cluster))
	builder.WriteString(" (\n")
	builder.WriteString("  ")
	for i, col := range table.Columns {
		definition, err := types.FieldDefinition(col.ToArrowField())
		if err != nil {
			return "", err
		}
		builder.WriteString(definition)
		if i < len(table.Columns)-1 {
			builder.WriteString(",\n  ")
		}
	}
	builder.WriteString("\n) ENGINE = ")
	builder.WriteString(engine.String())
	if partitionBySyncGroupID && hasColumn(table, cqSyncGroupID) {
		builder.WriteString(" PARTITION BY ")
		builder.WriteString(cqSyncGroupID)
	}
	builder.WriteString(" ORDER BY ")
	if orderBy := sortKeys(table); len(orderBy) > 0 {
		builder.WriteString("(")
		builder.WriteString(strings.Join(util.Sanitized(orderBy...), ", "))
		builder.WriteString(")")
	} else {
		builder.WriteString("tuple()")
	}
	builder.WriteString(" SETTINGS allow_nullable_key=1") // allows nullable keys

	return builder.String(), nil
}

func DropTable(table *schema.Table, cluster string) string {
	return "DROP TABLE IF EXISTS " + tableNamePart(table.Name, cluster)
}

func hasColumn(table *schema.Table, columnName string) bool {
	for _, col := range table.Columns {
		if col.Name == columnName {
			return true
		}
	}
	return false
}

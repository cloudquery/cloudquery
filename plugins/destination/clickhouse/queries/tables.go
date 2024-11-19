package queries

import (
	"fmt"
	"slices"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/client/spec"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/types"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v4/glob"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

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

func CreateTable(table *schema.Table, cluster string, engine *spec.Engine, partition []spec.PartitionStrategy, orderBy []spec.OrderByStrategy) (string, error) {
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
	partitionBy, err := resolvePartitionBy(table.Name, partition)
	if err != nil {
		return "", err
	}
	if partitionBy != "" {
		builder.WriteString(" PARTITION BY ")
		builder.WriteString(partitionBy)
	}
	builder.WriteString(" ORDER BY ")
	resolvedOrderBy, err := resolveOrderBy(table.Name, orderBy)
	if err != nil {
		return "", err
	}
	if len(resolvedOrderBy) > 0 {
		builder.WriteString("(")
		builder.WriteString(strings.Join(resolvedOrderBy, ", "))
		builder.WriteString(")")
	} else if sortKeys := sortKeys(table); len(sortKeys) > 0 {
		builder.WriteString("(")
		builder.WriteString(strings.Join(util.Sanitized(sortKeys...), ", "))
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

func resolvePartitionBy(table string, partition []spec.PartitionStrategy) (string, error) {
	hasMatchedAlready := false
	partitionBy := ""
	for _, p := range partition {
		if !tableMatchesAnyGlobPatterns(table, p.SkipTables) && tableMatchesAnyGlobPatterns(table, p.Tables) {
			if hasMatchedAlready {
				return "", fmt.Errorf("table %q matched multiple partition strategies", table)
			}
			hasMatchedAlready = true
			partitionBy = p.PartitionBy
		}
	}
	if !hasMatchedAlready {
		return "", nil
	}
	return partitionBy, nil
}

func resolveOrderBy(table string, orderBy []spec.OrderByStrategy) ([]string, error) {
	hasMatchedAlready := false
	resolvedOrderBy := []string{}
	for _, o := range orderBy {
		if !tableMatchesAnyGlobPatterns(table, o.SkipTables) && tableMatchesAnyGlobPatterns(table, o.Tables) {
			if hasMatchedAlready {
				return nil, fmt.Errorf("table %q matched multiple order by strategies", table)
			}
			hasMatchedAlready = true
			resolvedOrderBy = o.OrderBy
		}
	}
	if !hasMatchedAlready {
		return nil, nil
	}
	return resolvedOrderBy, nil
}

func tableMatchesAnyGlobPatterns(table string, patterns []string) bool {
	for _, pattern := range patterns {
		if glob.Glob(pattern, table) {
			return true
		}
	}
	return false
}

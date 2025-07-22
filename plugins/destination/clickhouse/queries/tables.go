package queries

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/client/spec"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/typeconv/ch/types"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/util"
	"github.com/cloudquery/plugin-sdk/v4/glob"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func SortKeys(table *schema.Table) []string {
	keys := make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		if (col.NotNull || col.PrimaryKey) && !IsCompoundType(col) {
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

func IsCompoundType(col schema.Column) bool {
	switch col.Type.(type) {
	case *arrow.StructType:
		return true
	case *arrow.MapType:
		return true
	case *arrow.ListType:
		return true
	default:
		return false
	}
}

type CreateTableOptions struct {
	TTL *time.Duration
}

type CreateTableOption func(*CreateTableOptions) *CreateTableOptions

func WithTTL(ttl time.Duration) CreateTableOption {
	return func(o *CreateTableOptions) *CreateTableOptions {
		o.TTL = &ttl
		return o
	}
}

func CreateTable(table *schema.Table, cluster string, engine *spec.Engine, partition []spec.PartitionStrategy, orderBy []spec.OrderByStrategy, ttl []spec.TTLStrategy) (string, error) {
	builder := strings.Builder{}
	builder.WriteString("CREATE TABLE IF NOT EXISTS ")
	builder.WriteString(tableNamePart(table.Name, cluster))
	builder.WriteString(" (\n")
	builder.WriteString("  ")
	isCqSyncTimeNotNull := false
	for i, col := range table.Columns {
		definition, err := types.FieldDefinition(col.ToArrowField())
		if err != nil {
			return "", err
		}
		builder.WriteString(definition)
		if i < len(table.Columns)-1 {
			builder.WriteString(",\n  ")
		}
		if col.Name == schema.CqSyncTimeColumn.Name && col.NotNull {
			isCqSyncTimeNotNull = true
		}
	}
	builder.WriteString("\n) ENGINE = ")
	builder.WriteString(engine.String())
	partitionBy, err := ResolvePartitionBy(table, partition)
	if err != nil {
		return "", err
	}
	if partitionBy != "" {
		builder.WriteString(" PARTITION BY ")
		builder.WriteString(partitionBy)
	}
	builder.WriteString(" ORDER BY ")
	resolvedOrderBy, err := ResolveOrderBy(table, orderBy)
	if err != nil {
		return "", err
	}
	if len(resolvedOrderBy) > 0 {
		builder.WriteString("(")
		builder.WriteString(strings.Join(resolvedOrderBy, ", "))
		builder.WriteString(")")
	} else {
		builder.WriteString("tuple()")
	}

	resolvedTTL, err := ResolveTTL(table, ttl)
	if err != nil {
		return "", err
	}
	if len(resolvedTTL) > 0 {
		builder.WriteString(" TTL " + GetTTLString(resolvedTTL, isCqSyncTimeNotNull))
	}

	builder.WriteString(" SETTINGS allow_nullable_key=1") // allows nullable keys

	return builder.String(), nil
}

func DropTable(table *schema.Table, cluster string) string {
	return "DROP TABLE IF EXISTS " + tableNamePart(table.Name, cluster)
}

func ResolvePartitionBy(table *schema.Table, partition []spec.PartitionStrategy) (string, error) {
	hasMatchedAlready := false
	partitionBy := ""
	for _, p := range partition {
		if p.SkipIncrementalTables && table.IsIncremental {
			continue
		}
		if !tableMatchesAnyGlobPatterns(table.Name, p.SkipTables) && tableMatchesAnyGlobPatterns(table.Name, p.Tables) {
			if hasMatchedAlready {
				return "", fmt.Errorf("table %q matched multiple partition strategies", table.Name)
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

func ResolveOrderBy(table *schema.Table, orderBy []spec.OrderByStrategy) ([]string, error) {
	hasMatchedAlready := false
	resolvedOrderBy := []string{}
	for _, o := range orderBy {
		if !tableMatchesAnyGlobPatterns(table.Name, o.SkipTables) && tableMatchesAnyGlobPatterns(table.Name, o.Tables) {
			if hasMatchedAlready {
				return nil, fmt.Errorf("table %q matched multiple order by strategies", table.Name)
			}
			hasMatchedAlready = true
			resolvedOrderBy = o.OrderBy
		}
	}
	if !hasMatchedAlready {
		return util.Sanitized(SortKeys(table)...), nil
	}
	return resolvedOrderBy, nil
}

func ResolveTTL(table *schema.Table, ttl []spec.TTLStrategy) (string, error) {
	hasMatchedAlready := false
	resolvedTTL := ""
	for _, t := range ttl {
		if !tableMatchesAnyGlobPatterns(table.Name, t.SkipTables) && tableMatchesAnyGlobPatterns(table.Name, t.Tables) {
			if hasMatchedAlready {
				return "", fmt.Errorf("table %q matched multiple TTL strategies", table.Name)
			}
			hasMatchedAlready = true
			resolvedTTL = t.TTL
		}
	}
	if !hasMatchedAlready {
		return "", nil
	}
	return resolvedTTL, nil
}

func GetTTLString(resolvedTTL string, isCqSyncTimeNotNull bool) string {
	if resolvedTTL == "" {
		return ""
	}

	// At the moment, _cq_sync_time is nullable in most instances of the CloudQuery CLI,
	// but the --cq-columns-not-null flag does allow users to control this. As ClickHouse TTLs
	// don't allow nullable columns to be used, we stay on the safe side and use a coalesce with fallback
	// to 1970, but for performance reasons only do this if _cq_sync_time is not guaranteed to be not null.
	if isCqSyncTimeNotNull {
		return "_cq_sync_time + (" + resolvedTTL + ")"
	} else {
		return "toDateTime(coalesce(_cq_sync_time, makeDate(1970, 1, 1))) + (" + resolvedTTL + ")"
	}
}

func tableMatchesAnyGlobPatterns(table string, patterns []string) bool {
	for _, pattern := range patterns {
		if glob.Glob(pattern, table) {
			return true
		}
	}
	return false
}

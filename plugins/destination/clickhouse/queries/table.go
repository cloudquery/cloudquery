package queries

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/arrow/types"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"golang.org/x/exp/slices"
)

func GetTablesSchema(database string) (query string, params []any) {
	const tableSchemaQuery = "SELECT `table`, `name`, `type` FROM system.columns WHERE `database` = @databaseName ORDER BY `table`, `position`"
	return tableSchemaQuery, []any{driver.NamedValue{Name: "databaseName", Value: database}}
}

// ScanTableSchemas doesn't close rows, so that's on caller.
func ScanTableSchemas(rows driver.Rows, need schema.Schemas) (schema.Schemas, error) {
	defs := make(map[string][]arrow.Field)

	var table, name, typ string
	for rows.Next() {
		if err := rows.Scan(&table, &name, &typ); err != nil {
			return nil, err
		}

		field, err := types.Field(name, typ)
		if err != nil {
			return nil, err
		}
		defs[table] = append(defs[table], *field)
	}

	res := make(schema.Schemas, 0, len(defs))
	for _, needed := range need {
		tableName := schema.TableName(needed)
		if def, ok := defs[tableName]; ok {
			metadata := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{tableName})
			res = append(res, arrow.NewSchema(def, &metadata))
		}
	}

	return slices.Clip(res), nil
}

func tableNamePart(table, cluster string) string {
	if len(cluster) > 0 {
		return util.SanitizeID(table) + " ON CLUSTER " + util.SanitizeID(cluster)
	}
	return util.SanitizeID(table)
}

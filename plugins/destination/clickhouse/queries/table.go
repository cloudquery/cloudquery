package queries

import (
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/typeconv/arrow/types"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/util"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func GetTablesSchema(database string) (query string, params []any) {
	const tableSchemaQuery = "SELECT `table`, `name`, `type` FROM system.columns WHERE `database` = @databaseName ORDER BY `table`, `position`"
	return tableSchemaQuery, []any{driver.NamedValue{Name: "databaseName", Value: database}}
}

// ScanTableSchemas doesn't close rows, so that's on caller.
func ScanTableSchemas(rows driver.Rows, messages message.WriteMigrateTables) (schema.Tables, error) {
	defs := make(map[string]schema.ColumnList, len(messages))

	var table, name, typ string
	for rows.Next() {
		if err := rows.Scan(&table, &name, &typ); err != nil {
			return nil, err
		}

		if !messages.Exists(table) {
			// only save the info about required tables
			continue
		}

		field, err := types.Field(name, typ)
		if err != nil {
			return nil, err
		}

		defs[table] = append(defs[table], schema.NewColumnFromArrowField(*field))
	}

	res := make(schema.Tables, 0, len(defs))
	for name, columns := range defs {
		res = append(res, &schema.Table{Name: name, Columns: columns})
	}

	return res, nil
}

func tableNamePart(table, cluster string) string {
	if len(cluster) > 0 {
		return util.SanitizeID(table) + " ON CLUSTER " + util.SanitizeID(cluster)
	}
	return util.SanitizeID(table)
}

func GetPartitionKeyAndSortingKeyQuery(database, table string) string {
	return fmt.Sprintf(`SELECT partition_key, sorting_key FROM system.tables WHERE database = '%s' AND name = '%s'`, database, table)
}

func GetTTLQuery(database, table string) string {
	return fmt.Sprintf(`SHOW CREATE TABLE "%s"."%s"`, database, table)
}

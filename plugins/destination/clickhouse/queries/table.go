package queries

import (
	"fmt"
	"strings"

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

func EqualTTLsQuery(table *schema.Table, ttl1, ttl2 string) string {
	// ClickHouse allows different syntaxes for the same TTL expression,
	// so we use a query to compare two given TTLs.
	// However, to evaluate the expression, we replace referenced columns with a fixed date.
	for i, col := range table.Columns {
		// Replace column names with a fixed date to ensure the comparison works
		ttl1 = strings.ReplaceAll(ttl1, col.Name, fmt.Sprintf("makeDate(1970, 1, 1) + INTERVAL %d SECOND", i))
		ttl2 = strings.ReplaceAll(ttl2, col.Name, fmt.Sprintf("makeDate(1970, 1, 1) + INTERVAL %d SECOND", i))
	}
	return fmt.Sprintf(`SELECT %s == %s AS equal`, ttl1, ttl2)
}

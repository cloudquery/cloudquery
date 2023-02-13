package queries

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/exp/maps"
)

func GetTablesSchema(database string) (query string, params []any) {
	const tableSchemaQuery = "SELECT `table`, `name`, `type` FROM system.columns WHERE `database` = @databaseName ORDER BY `table`, `position`"
	return tableSchemaQuery, []any{driver.NamedValue{Name: "databaseName", Value: database}}
}

// ScanTableSchemas doesn't close rows, so that's on caller.
func ScanTableSchemas(rows driver.Rows) (schema.Tables, error) {
	defs := make(map[string]*schema.Table)

	var table, name, typ string
	for rows.Next() {
		if err := rows.Scan(&table, &name, &typ); err != nil {
			return nil, err
		}

		def := defs[table]
		if def == nil {
			def = &schema.Table{Name: table}
			defs[table] = def
		}

		def.Columns = append(def.Columns, cqCol(name, typ))
	}

	return maps.Values(defs), nil
}

// NormalizedTables returns flattened normalized table definitions
func NormalizedTables(tables schema.Tables) schema.Tables {
	flattened := tables.FlattenTables()
	defs := make(schema.Tables, len(flattened))

	for i, tbl := range flattened {
		defs[i] = normalizeTable(tbl)
	}

	return defs
}

func normalizeTable(table *schema.Table) *schema.Table {
	columns := make(schema.ColumnList, len(table.Columns))

	for i, col := range table.Columns {
		columns[i] = normalizeColumn(col)
	}

	return &schema.Table{Name: table.Name, Columns: columns}
}

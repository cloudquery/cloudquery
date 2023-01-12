package queries

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func GetTablesSchema(database string) (query string, params []any) {
	const tableSchemaQuery = "SELECT `table`, `name`, `type` FROM system.columns WHERE `database` = @databaseName ORDER BY `table`, `position`"
	return tableSchemaQuery, []any{driver.NamedValue{Name: "databaseName", Value: database}}
}

// ScanTableSchemas doesn't close rows, so that's on caller.
func ScanTableSchemas(rows driver.Rows) (TableDefinitions, error) {
	defs := make(TableDefinitions)

	var table, name, typ string
	for rows.Next() {
		if err := rows.Scan(&table, &name, &typ); err != nil {
			return nil, err
		}

		def := defs[table]
		if def == nil {
			def = &TableDefinition{Name: table}
			defs[table] = def
		}

		def.Columns = append(def.Columns, &ColumnDefinition{Name: name, Type: typ})
	}

	return defs, nil
}

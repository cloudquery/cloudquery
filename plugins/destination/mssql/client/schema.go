package client

import (
	"context"
	"database/sql"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) schemaTables(ctx context.Context, tables schema.Tables) (schema.Tables, error) {
	query, params := queries.AllTables(c.schemaName)
	rows, err := c.db.QueryContext(ctx, query, params...)
	if err != nil {
		c.logErr(err)
		return nil, err
	}

	names := make([]string, 0)
	if err := processRows(rows, func(row *sql.Rows) error {
		var tableCatalog string
		var tableName string
		var tableType string
		var schemaType string

		if err := row.Scan(&tableCatalog, &tableType, &tableName, &schemaType); err != nil {
			return err
		}
		if tables.Get(tableName) == nil {
			return nil
		}
		names = append(names, tableName)
		return nil
	}); err != nil {
		c.logErr(err)
		return nil, err
	}

	result := make(schema.Tables, len(names))
	for i, tableName := range names {
		pks, err := c.getTablePK(ctx, tableName)
		if err != nil {
			return nil, err
		}

		columns, err := c.getTableColumns(ctx, tableName, pks)
		if err != nil {
			return nil, err
		}

		result[i] = &schema.Table{Name: tableName, Columns: columns}
	}

	return result, nil
}

func (c *Client) normalizedTables(tables schema.Tables) schema.Tables {
	normalized := make(schema.Tables, len(tables))
	for i, table := range tables {
		normalized[i] = c.normalizeTable(table)
	}
	return normalized
}

func (c *Client) normalizeTable(table *schema.Table) *schema.Table {
	columns := make(schema.ColumnList, len(table.Columns))

	for i, col := range table.Columns {
		// Since multiple schema types can map to the same MSSQL type
		// we need to normalize them to avoid false positives when detecting schema changes.
		// This should never return an error
		col.Type = queries.SchemaType(queries.SQLType(col.Type))
		if c.pkEnabled() && col.PrimaryKey {
			col.NotNull = true
		}
		columns[i] = col
	}

	return &schema.Table{Name: table.Name, Columns: columns}
}

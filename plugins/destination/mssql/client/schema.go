package client

import (
	"context"
	"database/sql"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/v5/queries"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) schemaTables(ctx context.Context, messages message.WriteMigrateTables) (schema.Tables, error) {
	query, params := queries.AllTables(c.spec.Schema)
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
		if !messages.Exists(tableName) {
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

func normalizedTables(messages message.WriteMigrateTables) schema.Tables {
	normalized := make(schema.Tables, len(messages))
	for i, m := range messages {
		normalized[i] = normalizeTable(m.Table)
	}
	return normalized
}

func normalizeTable(table *schema.Table) *schema.Table {
	columns := make(schema.ColumnList, len(table.Columns))

	for i, col := range table.Columns {
		// Since multiple schema types can map to the same MSSQL type
		// we need to normalize them to avoid false positives when detecting schema changes.
		// This should never return an error
		col.Type = queries.SchemaType(queries.SQLType(col.Type, col.PrimaryKey))
		col.NotNull = col.NotNull || col.PrimaryKey
		columns[i] = col
	}

	return &schema.Table{Name: table.Name, Columns: columns}
}

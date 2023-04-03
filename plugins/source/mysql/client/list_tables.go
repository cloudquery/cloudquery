package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func Identifier(name string) string {
	return "`" + name + "`"
}

func (c *Client) getTableColumns(ctx context.Context, table *schema.Table) (schema.ColumnList, error) {
	query := `SELECT COLUMN_NAME, DATA_TYPE, COLUMN_TYPE, IS_NULLABLE, COLUMN_KEY FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ?;`
	var tc schema.ColumnList

	rows, err := c.db.QueryContext(ctx, query, table.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var dataType string
		// columnType has additional information, such such as the precision or length
		var columnType string
		var nullable string
		var key string

		if err := rows.Scan(&name, &dataType, &columnType, &nullable, &key); err != nil {
			return nil, err
		}

		schemaType, err := SchemaType(table.Name, name, dataType, columnType)
		if err != nil {
			return nil, err
		}
		column := schema.Column{
			Name: name, Type: schemaType,
			CreationOptions: schema.ColumnCreationOptions{NotNull: nullable == "NO", PrimaryKey: key == "PRI"},
		}
		tc = append(tc, column)
	}

	return tc, nil
}

func (c *Client) listTables(ctx context.Context) (schema.Tables, error) {
	query := `SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE' AND TABLE_SCHEMA = ?;`
	rows, err := c.db.QueryContext(ctx, query, c.tableSchema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	schemaTables := make(schema.Tables, 0)
	for rows.Next() {
		var tableName string

		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		schemaTables = append(schemaTables, &schema.Table{Name: tableName})
	}

	for _, table := range schemaTables {
		columns, err := c.getTableColumns(ctx, table)
		if err != nil {
			return nil, err
		}
		table.Columns = columns
	}

	return schemaTables, nil
}

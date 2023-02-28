package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func identifier(name string) string {
	return fmt.Sprintf("`%s`", name)
}

func (c *Client) getTableColumns(ctx context.Context, table *schema.Table) (schema.ColumnList, error) {
	query := `SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, CHARACTER_MAXIMUM_LENGTH, COLUMN_KEY FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ?;`
	var tc schema.ColumnList

	rows, err := c.db.QueryContext(ctx, query, table.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var typ string
		var nullable string
		var charMaxLength *string
		var key string

		if err := rows.Scan(&name, &typ, &nullable, &charMaxLength, &key); err != nil {
			return nil, err
		}

		schemaType, err := SchemaType(table.Name, name, typ)
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

func (c *Client) schemaTables(ctx context.Context, tables schema.Tables) (schema.Tables, error) {
	query := `SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE';`
	rows, err := c.db.QueryContext(ctx, query)
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
		if tables.Get(tableName) == nil {
			continue
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

func (c *Client) addColumn(ctx context.Context, table *schema.Table, column schema.Column) error {
	builder := strings.Builder{}
	builder.WriteString("ALTER TABLE ")
	builder.WriteString(identifier(table.Name))
	builder.WriteString(" ADD COLUMN ")
	builder.WriteString(identifier(column.Name))
	builder.WriteString(" ")
	builder.WriteString(SQLType(column.Type))
	if column.CreationOptions.NotNull {
		builder.WriteString(" NOT NULL")
	}
	if column.CreationOptions.Unique {
		builder.WriteString(" UNIQUE")
	}
	builder.WriteString(";")
	_, err := c.db.ExecContext(ctx, builder.String())
	return err
}

func (c *Client) createTable(ctx context.Context, table *schema.Table) error {
	builder := strings.Builder{}
	builder.WriteString("CREATE TABLE ")
	builder.WriteString(identifier(table.Name))
	builder.WriteString(" (\n  ")
	for i, column := range table.Columns {
		builder.WriteString(identifier(column.Name))
		builder.WriteString(" ")
		builder.WriteString(SQLType(column.Type))
		if column.CreationOptions.Unique {
			builder.WriteString(" UNIQUE")
		}
		if column.CreationOptions.NotNull {
			builder.WriteString(" NOT NULL")
		}
		if i < len(table.Columns)-1 {
			builder.WriteString(",\n  ")
		}
	}
	pks := table.PrimaryKeys()
	if len(pks) > 0 {
		builder.WriteString(",\n  ")
		builder.WriteString(" PRIMARY KEY (")
		for i, pk := range pks {
			builder.WriteString(identifier(pk))
			if table.Columns.Get(pk).Type == schema.TypeString {
				// Since we use `text` for strings we need to specify the prefix length to use for the primary key
				builder.WriteString("(64)")
			}
			if i < len(pks)-1 {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(")\n")
	}
	builder.WriteString(") CHARACTER SET utf8mb4;")
	_, err := c.db.ExecContext(ctx, builder.String())
	return err
}

func (c *Client) dropTable(ctx context.Context, table *schema.Table) error {
	_, err := c.db.ExecContext(ctx, "DROP TABLE "+identifier(table.Name))
	return err
}

func (c *Client) recreateTable(ctx context.Context, table *schema.Table) error {
	if err := c.dropTable(ctx, table); err != nil {
		return err
	}
	return c.createTable(ctx, table)
}

package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func identifier(name string) string {
	return fmt.Sprintf("`%s`", name)
}

func (c *Client) getTableColumns(ctx context.Context, tableName string) ([]schema.Column, error) {
	query := `SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, CHARACTER_MAXIMUM_LENGTH, COLUMN_KEY FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ?;`
	var columns []schema.Column

	rows, err := c.db.QueryContext(ctx, query, tableName)
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

		schemaType, err := mySQLTypeToArrowType(tableName, name, typ)
		if err != nil {
			return nil, err
		}

		columns = append(columns, schema.Column{
			Name:       name,
			Type:       schemaType,
			PrimaryKey: key == "PRI",
			NotNull:    nullable != "YES",
		})
	}

	return columns, nil
}

func (c *Client) schemaTables(ctx context.Context, tables schema.Tables) (schema.Tables, error) {
	query := `SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE';`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	schemaTables := make(schema.Tables, 0)
	tableNames := make([]string, 0)
	for rows.Next() {
		var tableName string

		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}

		if tables.Get(tableName) == nil {
			continue
		}

		tableNames = append(tableNames, tableName)
	}

	for _, tableName := range tableNames {
		fields, err := c.getTableColumns(ctx, tableName)
		if err != nil {
			return nil, err
		}
		schemaTables = append(schemaTables, &schema.Table{Name: tableName, Columns: fields})
	}

	return schemaTables, nil
}

func (c *Client) addColumn(ctx context.Context, table *schema.Table, column *schema.Column) error {

	builder := strings.Builder{}
	builder.WriteString("ALTER TABLE ")
	builder.WriteString(identifier(table.Name))
	builder.WriteString(" ADD COLUMN ")
	builder.WriteString(identifier(column.Name))
	builder.WriteString(" ")
	builder.WriteString(arrowTypeToMySqlStr(column.Type))
	if column.NotNull {
		builder.WriteString(" NOT NULL")
	}
	if column.Unique {
		builder.WriteString(" UNIQUE")
	}
	builder.WriteString(";")
	_, err := c.db.ExecContext(ctx, builder.String())
	return err
}

func (c *Client) createTable(ctx context.Context, table *schema.Table) error {

	totalColumns := len(table.Columns)
	primaryKeysIndices := []int{}

	builder := strings.Builder{}
	builder.WriteString("CREATE TABLE ")
	builder.WriteString(identifier(table.Name))
	builder.WriteString(" (\n  ")
	for i, column := range table.Columns {
		builder.WriteString(identifier(column.Name))
		builder.WriteString(" ")
		builder.WriteString(arrowTypeToMySqlStr(column.Type))
		if column.Unique {
			builder.WriteString(" UNIQUE")
		}
		if column.NotNull {
			builder.WriteString(" NOT NULL")
		}
		if i < totalColumns-1 {
			builder.WriteString(",\n  ")
		}

		if c.pkEnabled() && column.PrimaryKey {
			primaryKeysIndices = append(primaryKeysIndices, i)
		}
	}
	if len(primaryKeysIndices) > 0 {
		builder.WriteString(",\n  ")
		builder.WriteString(" PRIMARY KEY (")
		for i, pk := range primaryKeysIndices {
			column := table.Columns[pk]
			builder.WriteString(identifier(column.Name))
			if column.Type == arrow.BinaryTypes.LargeString {
				// Since we use `text` for strings we need to specify the prefix length to use for the primary key
				builder.WriteString("(64)")
			}
			if i < len(primaryKeysIndices)-1 {
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

package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func identifier(name string) string {
	return fmt.Sprintf("`%s`", name)
}

func (c *Client) getTableColumns(ctx context.Context, tableName string) ([]arrow.Field, error) {
	query := `SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, CHARACTER_MAXIMUM_LENGTH, COLUMN_KEY FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ?;`
	var fields []arrow.Field

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
		var fieldMetadata schema.MetadataFieldOptions
		if key == "PRI" {
			fieldMetadata.PrimaryKey = true
		}
		field := arrow.Field{
			Name:     name,
			Type:     schemaType,
			Nullable: nullable == "YES",
			Metadata: schema.NewFieldMetadataFromOptions(fieldMetadata),
		}
		fields = append(fields, field)
	}

	return fields, nil
}

func (c *Client) schemaTables(ctx context.Context, tables schema.Schemas) (schema.Schemas, error) {
	query := `SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE';`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	schemaTables := make(schema.Schemas, 0)
	tableNames := make([]string, 0)
	for rows.Next() {
		var tableName string

		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}

		if tables.SchemaByName(tableName) == nil {
			continue
		}

		tableNames = append(tableNames, tableName)
	}

	for _, tableName := range tableNames {
		fields, err := c.getTableColumns(ctx, tableName)
		if err != nil {
			return nil, err
		}

		var tableMetadata schema.MetadataSchemaOptions
		tableMetadata.TableName = tableName
		m := schema.NewSchemaMetadataFromOptions(tableMetadata)
		schemaTables = append(schemaTables, arrow.NewSchema(fields, &m))
	}

	return schemaTables, nil
}

func (c *Client) addColumn(ctx context.Context, table *arrow.Schema, column arrow.Field) error {
	tableName, ok := table.Metadata().GetValue(schema.MetadataTableName)
	if !ok {
		return fmt.Errorf("schema %s has no table name", table.String())
	}
	builder := strings.Builder{}
	builder.WriteString("ALTER TABLE ")
	builder.WriteString(identifier(tableName))
	builder.WriteString(" ADD COLUMN ")
	builder.WriteString(identifier(column.Name))
	builder.WriteString(" ")
	builder.WriteString(arrowTypeToMySqlStr(column.Type))
	if !column.Nullable {
		builder.WriteString(" NOT NULL")
	}
	if schema.IsUnique(column) {
		builder.WriteString(" UNIQUE")
	}
	builder.WriteString(";")
	_, err := c.db.ExecContext(ctx, builder.String())
	return err
}

func (c *Client) createTable(ctx context.Context, table *arrow.Schema) error {
	tableName, ok := table.Metadata().GetValue(schema.MetadataTableName)
	if !ok {
		return fmt.Errorf("schema %s has no table name", table.String())
	}

	totalColumns := len(table.Fields())
	primaryKeysIndices := []int{}

	builder := strings.Builder{}
	builder.WriteString("CREATE TABLE ")
	builder.WriteString(identifier(tableName))
	builder.WriteString(" (\n  ")
	for i, column := range table.Fields() {
		builder.WriteString(identifier(column.Name))
		builder.WriteString(" ")
		builder.WriteString(arrowTypeToMySqlStr(column.Type))
		if schema.IsUnique(column) {
			builder.WriteString(" UNIQUE")
		}
		if !column.Nullable {
			builder.WriteString(" NOT NULL")
		}
		if i < totalColumns-1 {
			builder.WriteString(",\n  ")
		}

		if c.pkEnabled() && schema.IsPk(column) {
			primaryKeysIndices = append(primaryKeysIndices, i)
		}
	}
	if len(primaryKeysIndices) > 0 {
		builder.WriteString(",\n  ")
		builder.WriteString(" PRIMARY KEY (")
		for i, pk := range primaryKeysIndices {
			field := table.Field(pk)
			builder.WriteString(identifier(field.Name))
			if field.Type == arrow.BinaryTypes.LargeString {
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

func (c *Client) dropTable(ctx context.Context, table *arrow.Schema) error {
	tableName, ok := table.Metadata().GetValue(schema.MetadataTableName)
	if !ok {
		return fmt.Errorf("schema %s has no table name", table.String())
	}

	_, err := c.db.ExecContext(ctx, "DROP TABLE "+identifier(tableName))
	return err
}

func (c *Client) recreateTable(ctx context.Context, table *arrow.Schema) error {
	if err := c.dropTable(ctx, table); err != nil {
		return err
	}
	return c.createTable(ctx, table)
}

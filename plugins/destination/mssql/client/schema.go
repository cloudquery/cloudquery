package client

import (
	"context"
	"database/sql"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) schemaTables(ctx context.Context, scs schema.Schemas) (schema.Schemas, error) {
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
		if scs.SchemaByName(tableName) == nil {
			return nil
		}
		names = append(names, tableName)
		return nil
	}); err != nil {
		c.logErr(err)
		return nil, err
	}

	tables := make([]*arrow.Schema, len(names))
	for i, tableName := range names {
		pks, err := c.getTablePK(ctx, tableName)
		if err != nil {
			return nil, err
		}

		fields, err := c.getTableFields(ctx, tableName, pks)
		if err != nil {
			return nil, err
		}

		tableMD := schema.NewSchemaMetadataFromOptions(schema.MetadataSchemaOptions{TableName: tableName})
		tables[i] = arrow.NewSchema(fields, &tableMD)
	}

	return tables, nil
}

func (c *Client) normalizedSchemas(scs schema.Schemas) schema.Schemas {
	normalized := make(schema.Schemas, len(scs))
	for i, sc := range scs {
		normalized[i] = c.normalizeSchema(sc)
	}
	return normalized
}

func (c *Client) normalizeSchema(sc *arrow.Schema) *arrow.Schema {
	tableName := schema.TableName(sc)
	fields := make([]arrow.Field, len(sc.Fields()))

	for i, field := range sc.Fields() {
		// Since multiple schema types can map to the same MSSQL type we need to normalize them to avoid false positives when detecting schema changes
		// This should never return an error
		field.Type = queries.SchemaType(queries.SQLType(field.Type))
		field.Metadata = c.normalizeFieldMetadata(field.Metadata)
		if schema.IsPk(field) {
			field.Nullable = false
		}
		fields[i] = field
	}

	tableMD := schema.NewSchemaMetadataFromOptions(schema.MetadataSchemaOptions{TableName: tableName})
	return arrow.NewSchema(fields, &tableMD)
}

func (c *Client) normalizeFieldMetadata(metadata arrow.Metadata) arrow.Metadata {
	keys, values := metadata.Keys(), metadata.Values()
	md := make(map[string]string, len(keys))

	for idx, key := range keys {
		switch key {
		case schema.MetadataUnique:
			continue // we don't scan unique constraints
		case schema.MetadataPrimaryKey:
			if !c.pkEnabled() {
				continue
			}
		}

		md[key] = values[idx]
	}

	return arrow.MetadataFrom(md)
}

package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

func getInsertQueryBuild(table *arrow.Schema) *strings.Builder {
	tableName := schema.TableName(table)
	builder := strings.Builder{}
	builder.WriteString("INSERT INTO " + identifier(tableName))
	builder.WriteString(" (")
	fields := table.Fields()
	for i, col := range table.Fields() {
		builder.WriteString(identifier(col.Name))
		if i < len(fields)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(") VALUES (")
	builder.WriteString(strings.TrimSuffix(strings.Repeat("?,", len(fields)), ","))
	builder.WriteString(")")
	return &builder
}

func (c *Client) writeResources(ctx context.Context, query string, resources []arrow.Record) error {
	for _, data := range resources {
		transformedRecords := transformRecord(data)
		for _, transformedRecord := range transformedRecords {
			_, err := c.db.ExecContext(ctx, query, transformedRecord...)
			data.Release()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) appendTableBatch(ctx context.Context, table *arrow.Schema, resources []arrow.Record) error {
	builder := getInsertQueryBuild(table)
	builder.WriteString(";")
	return c.writeResources(ctx, builder.String(), resources)
}

func (c *Client) overwriteTableBatch(ctx context.Context, table *arrow.Schema, resources []arrow.Record) error {
	builder := getInsertQueryBuild(table)
	builder.WriteString(" ON DUPLICATE KEY UPDATE ")
	fields := table.Fields()
	for i, col := range table.Fields() {
		builder.WriteString(fmt.Sprintf("%s = VALUES(%s)", identifier(col.Name), identifier(col.Name)))
		if i < len(fields)-1 {
			builder.WriteString(", ")
		}
	}
	return c.writeResources(ctx, builder.String(), resources)
}

func (c *Client) WriteTableBatch(ctx context.Context, table *arrow.Schema, resources []arrow.Record) error {
	switch c.spec.WriteMode {
	case specs.WriteModeAppend:
		return c.appendTableBatch(ctx, table, resources)
	case specs.WriteModeOverwrite, specs.WriteModeOverwriteDeleteStale:
		return c.overwriteTableBatch(ctx, table, resources)
	default:
		panic("unsupported write mode " + c.spec.WriteMode.String())
	}
}

package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func getInsertQueryBuild(table *schema.Table) *strings.Builder {
	builder := strings.Builder{}
	builder.WriteString("INSERT INTO " + table.Name)
	builder.WriteString(" (")
	for i, col := range table.Columns {
		builder.WriteString(identifier(col.Name))
		if i < len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(") VALUES (")
	builder.WriteString(strings.TrimSuffix(strings.Repeat("?,", len(table.Columns)), ","))
	builder.WriteString(")")
	return &builder
}

func (c *Client) appendTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	builder := getInsertQueryBuild(table)
	builder.WriteString(";")

	for _, data := range resources {
		_, err := c.db.ExecContext(ctx, builder.String(), data...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) overwriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	builder := getInsertQueryBuild(table)
	builder.WriteString(" ON DUPLICATE KEY UPDATE ")
	for i, col := range table.Columns {
		builder.WriteString(fmt.Sprintf("%s = VALUES(%s)", identifier(col.Name), identifier(col.Name)))
		if i < len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(";")

	for _, data := range resources {
		_, err := c.db.ExecContext(ctx, builder.String(), data...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	switch c.spec.WriteMode {
	case specs.WriteModeAppend:
		return c.appendTableBatch(ctx, table, resources)
	case specs.WriteModeOverwrite, specs.WriteModeOverwriteDeleteStale:
		return c.overwriteTableBatch(ctx, table, resources)
	default:
		panic("unsupported write mode " + c.spec.WriteMode.String())
	}
}

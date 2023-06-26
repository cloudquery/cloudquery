package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func getInsertQueryBuild(table *schema.Table) *strings.Builder {
	builder := strings.Builder{}
	builder.WriteString("INSERT INTO " + identifier(table.Name))
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

func (c *Client) writeResources(ctx context.Context, query string, resources []arrow.Record) error {
	for _, data := range resources {
		transformedRecords, err := transformRecord(data)
		if err != nil {
			return err
		}
		for _, transformedRecord := range transformedRecords {
			_, err := c.db.ExecContext(ctx, query, transformedRecord...)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) appendTableBatch(ctx context.Context, table *schema.Table, resources []arrow.Record) error {
	builder := getInsertQueryBuild(table)
	builder.WriteString(";")
	return c.writeResources(ctx, builder.String(), resources)
}

func (c *Client) overwriteTableBatch(ctx context.Context, table *schema.Table, resources []arrow.Record) error {
	builder := getInsertQueryBuild(table)
	builder.WriteString(" ON DUPLICATE KEY UPDATE ")
	for i, col := range table.Columns {
		builder.WriteString(fmt.Sprintf("%s = VALUES(%s)", identifier(col.Name), identifier(col.Name)))
		if i < len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	return c.writeResources(ctx, builder.String(), resources)
}

func (c *Client) Write(ctx context.Context, options plugin.WriteOptions, res <-chan message.Message) error {
	if err := c.writer.Write(ctx, res); err != nil {
		return err
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, name string, msgs []*message.Insert) error {
	if len(msgs) == 0 {
		return nil
	}
	table := msgs[0].GetTable()
	hasPks := len(table.PrimaryKeys()) > 0
	if hasPks {
		return c.overwriteTableBatch(ctx, table, resources)
	} else {
		return c.appendTableBatch(ctx, table, resources)
	}
}

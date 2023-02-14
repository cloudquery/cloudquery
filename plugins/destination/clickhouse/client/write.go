package client

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	batch, err := c.conn.PrepareBatch(ctx, queries.Insert(table))
	if err != nil {
		return err
	}

	for _, row := range data {
		if err := batch.Append(row...); err != nil {
			_ = batch.Abort()
			return err
		}
	}

	return batch.Send()
}

package client

import (
	"context"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *arrow.Schema, data []arrow.Record) error {
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

package client

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/values"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, records []arrow.Record) error {
	batch, err := c.conn.PrepareBatch(ctx, queries.Insert(table))
	if err != nil {
		return err
	}

	if err := values.BatchAddRecords(ctx, batch, table, records); err != nil {
		_ = batch.Abort()
		return err
	}

	return batch.Send()
}

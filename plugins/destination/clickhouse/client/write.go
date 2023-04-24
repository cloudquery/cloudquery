package client

import (
	"context"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/values"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *arrow.Schema, records []arrow.Record) error {
	batch, err := c.conn.PrepareBatch(ctx, queries.Insert(table))
	if err != nil {
		return err
	}

	for _, record := range records {
		if err := values.BatchAddRecord(batch, record); err != nil {
			_ = batch.Abort()
			return err
		}
	}

	return batch.Send()
}

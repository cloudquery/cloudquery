package client

import (
	"context"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/values"
)

func (c *Client) WriteTableBatch(ctx context.Context, sc *arrow.Schema, records []arrow.Record) error {
	batch, err := c.conn.PrepareBatch(ctx, queries.Insert(sc))
	if err != nil {
		return err
	}

	if err := values.BatchAddRecords(ctx, sc, batch, records); err != nil {
		_ = batch.Abort()
		return err
	}

	return batch.Send()
}

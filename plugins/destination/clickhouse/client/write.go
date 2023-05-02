package client

import (
	"context"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/values"
)

func (c *Client) WriteTableBatch(ctx context.Context, sc *arrow.Schema, records []arrow.Record) error {
	reader, err := array.NewRecordReader(sc, records)
	if err != nil {
		return err
	}

	batch, err := c.conn.PrepareBatch(ctx, queries.Insert(sc))
	if err != nil {
		return err
	}

	if err := values.BatchAddRecords(ctx, batch, reader); err != nil {
		_ = batch.Abort()
		return err
	}

	return batch.Send()
}

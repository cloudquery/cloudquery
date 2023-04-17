package client

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *arrow.Schema, records []arrow.Record) error {
	batch, err := c.conn.PrepareBatch(ctx, queries.Insert(table))
	if err != nil {
		return err
	}

	for _, record := range records {
		if err := addRecord(batch, record); err != nil {
			_ = batch.Abort()
			return err
		}
	}

	return batch.Send()
}

func addRecord(batch driver.Batch, record arrow.Record) error {
	for i, col := range record.Columns() {
		if err := addColumn(batch.Column(i), col); err != nil {
			return err
		}
	}
	return nil
}

func addColumn(column driver.BatchColumn, arr arrow.Array) error {
	for i := 0; i < arr.Len(); i++ {
		switch {
		case arr.IsNull(i), !arr.IsValid(i):
			if err := column.Append(nil); err != nil {
				return err
			}
		default:
			if err := column.Append(arr.GetOneForMarshal(i)); err != nil {
				return err
			}
		}
	}
	return nil
}

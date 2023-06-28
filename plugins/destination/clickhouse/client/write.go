package client

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/values"
	"github.com/cloudquery/plugin-sdk/v4/message"
)

func (c *Client) WriteTableBatch(ctx context.Context, name string, messages []*message.Insert) error {
	if len(messages) == 0 {
		return nil
	}

	table := messages[0].GetTable()
	records := make([]arrow.Record, len(messages))
	for i, m := range messages {
		records[i] = m.Record
	}

	batch, err := c.conn.PrepareBatch(ctx, queries.Insert(table))
	if err != nil {
		return err
	}

	if err := values.BatchAddRecords(ctx, batch, table.ToArrowSchema(), records); err != nil {
		_ = batch.Abort()
		return err
	}

	return batch.Send()
}

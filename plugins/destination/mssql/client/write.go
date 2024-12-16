package client

import (
	"context"
	"database/sql"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
)

func (c *Client) Write(ctx context.Context, messages <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, messages); err != nil {
		return err
	}
	return c.writer.Flush(ctx)
}

func (c *Client) WriteTableBatch(ctx context.Context, _ string, messages message.WriteInserts) error {
	if len(messages) == 0 {
		return nil
	}

	table := messages[0].GetTable()
	records := make([]arrow.Record, len(messages))
	for i, m := range messages {
		records[i] = m.Record
	}

	if useTVP(table) {
		return c.insertTVP(ctx, table, records)
	}

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		return c.bulkInsert(ctx, tx, table, records)
	})
}

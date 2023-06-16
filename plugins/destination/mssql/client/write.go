package client

import (
	"context"
	"database/sql"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, records []arrow.Record) error {
	if c.useTVP(table) {
		return c.insertTVP(ctx, table, records)
	}

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		return c.bulkInsert(ctx, tx, table, records)
	})
}

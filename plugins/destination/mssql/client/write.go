package client

import (
	"context"
	"database/sql"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	if c.spec.WriteMode == specs.WriteModeAppend {
		return c.doInTx(ctx, func(tx *sql.Tx) error {
			return c.bulkInsert(ctx, tx, table, data)
		})
	}

	return c.insertTVP(ctx, table, data)
}

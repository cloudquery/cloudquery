package client

import (
	"context"
	"database/sql"

	"github.com/apache/arrow/go/v13/arrow"
)

func (c *Client) WriteTableBatch(ctx context.Context, sc *arrow.Schema, records []arrow.Record) error {
	if c.useTVP(sc) {
		return c.insertTVP(ctx, sc, records)
	}

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		return c.bulkInsert(ctx, tx, sc, records)
	})
}

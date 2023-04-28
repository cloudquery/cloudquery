package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
)

func (c *Client) WriteTableBatch(ctx context.Context, sc *arrow.Schema, records []arrow.Record) error {
	return fmt.Errorf("todo")

	if c.useTVP(sc) {
		return c.insertTVP(ctx, sc, records)
	}

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		return c.bulkInsert(ctx, tx, sc, records)
	})
}

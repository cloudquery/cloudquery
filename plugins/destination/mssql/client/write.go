package client

import (
	"context"
	"database/sql"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

func (c *Client) WriteTableBatch(ctx context.Context, sc *arrow.Schema, records []arrow.Record) error {
	if c.spec.WriteMode == specs.WriteModeAppend {
		return c.doInTx(ctx, func(tx *sql.Tx) error {
			return c.bulkInsert(ctx, tx, sc, records)
		})
	}

	return c.insertTVP(ctx, sc, records)
}

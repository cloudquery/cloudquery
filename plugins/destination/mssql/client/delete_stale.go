package client

import (
	"context"
	"database/sql"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	return c.doInTx(ctx, func(tx *sql.Tx) error {
		for _, table := range tables {
			query, params := queries.DeleteStale(c.spec.Schema, table, sourceName, syncTime)
			_, err := tx.ExecContext(ctx, query, params...)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

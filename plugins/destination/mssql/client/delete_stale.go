package client

import (
	"context"
	"database/sql"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/v5/queries"
	"github.com/cloudquery/plugin-sdk/v4/message"
)

func (c *Client) DeleteStale(ctx context.Context, messages message.WriteDeleteStales) error {
	return c.doInTx(ctx, func(tx *sql.Tx) error {
		for _, m := range messages {
			query, params := queries.DeleteStale(c.spec.Schema, m)
			_, err := tx.ExecContext(ctx, query, params...)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

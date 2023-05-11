package client

import (
	"context"
	"database/sql"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) DeleteStale(ctx context.Context, scs schema.Schemas, sourceName string, syncTime time.Time) error {
	return c.doInTx(ctx, func(tx *sql.Tx) error {
		for _, sc := range scs {
			query, params := queries.DeleteStale(c.schemaName, sc, sourceName, syncTime)
			_, err := tx.ExecContext(ctx, query, params...)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

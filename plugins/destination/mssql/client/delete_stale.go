package client

import (
	"context"
	"database/sql"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	return c.doInTx(ctx, func(tx *sql.Tx) error {
		for _, table := range tables.FlattenTables() {
			query, params := queries.DeleteStale(c.schemaName, table, sourceName, syncTime)
			_, err := tx.ExecContext(ctx, query, params...)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

package client

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/client/queries"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	return c.doInTx(ctx, func(tx *sql.Tx) error {
		for _, table := range tables.FlattenTables() {
			_, err := tx.ExecContext(ctx,
				fmt.Sprintf(`delete from %s where %s = @sourceName and %s < @syncTime`,
					c.tableName(table),
					queries.SanitizeID(schema.CqSourceNameColumn.Name),
					queries.SanitizeID(schema.CqSyncTimeColumn.Name),
				),
				sql.Named("sourceName", sourceName),
				sql.Named("syncTime", syncTime),
			)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

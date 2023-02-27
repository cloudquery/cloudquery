package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	for _, table := range tables.FlattenTables() {
		query := fmt.Sprintf(`delete from %s where %s = ? and %s < ?`, identifier(table.Name), identifier(schema.CqSourceNameColumn.Name), identifier(schema.CqSyncTimeColumn.Name))
		if _, err := c.db.ExecContext(ctx, query, source, syncTime); err != nil {
			return err
		}
	}
	return nil
}

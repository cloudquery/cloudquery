package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Schemas, source string, syncTime time.Time) error {
	for _, table := range tables {
		name := schema.TableName(table)
		query := fmt.Sprintf(`delete from %s where %s = ? and %s < ?`, identifier(name), identifier(schema.CqSourceNameColumn.Name), identifier(schema.CqSyncTimeColumn.Name))
		if _, err := c.db.ExecContext(ctx, query, source, syncTime); err != nil {
			return err
		}
	}
	return nil
}

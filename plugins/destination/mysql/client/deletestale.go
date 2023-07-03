package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) DeleteStale(ctx context.Context, msgs message.WriteDeleteStales) error {
	for _, msg := range msgs {
		tableName := msg.TableName
		query := fmt.Sprintf(`delete from %s where %s = ? and %s < ?`, identifier(tableName), identifier(schema.CqSourceNameColumn.Name), identifier(schema.CqSyncTimeColumn.Name))
		if _, err := c.db.ExecContext(ctx, query, msg.SourceName, msg.SyncTime); err != nil {
			return err
		}
	}
	return nil
}

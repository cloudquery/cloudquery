package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) DeleteStale(ctx context.Context, msgs message.WriteDeleteStales) error {
	for _, msg := range msgs {
		if _, err := c.db.ExecContext(ctx, deleteStaleQuery(msg.TableName), msg.SourceName, msg.SyncTime); err != nil {
			return err
		}
	}
	return nil
}

func deleteStaleQuery(tableName string) string {
	return fmt.Sprintf(`delete from %s where %s = ? and %s < ?`, identifier(tableName), identifier(schema.CqSourceNameColumn.Name), identifier(schema.CqSyncTimeColumn.Name))
}

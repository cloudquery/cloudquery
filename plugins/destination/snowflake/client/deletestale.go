package client

import (
	"context"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) DeleteStale(ctx context.Context, msgs message.WriteDeleteStales) error {
	for _, msg := range msgs {
		tableName := msg.TableName
		var sb strings.Builder
		sb.WriteString("delete from ")
		sb.WriteString(tableName)
		sb.WriteString(" where ")
		sb.WriteString(`"` + strings.ToUpper(schema.CqSourceNameColumn.Name) + `"`)
		sb.WriteString(" = ? and \"")
		sb.WriteString(strings.ToUpper(schema.CqSyncTimeColumn.Name))
		sb.WriteString("\" < CONVERT_TIMEZONE('UTC','UTC', TO_TIMESTAMP(?))")
		sql := sb.String()
		if _, err := c.db.ExecContext(ctx, sql, msg.SourceName, msg.SyncTime); err != nil {
			return err
		}
	}
	return nil
}

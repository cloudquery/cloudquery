package client

import (
	"context"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var delayAfterDeleteStale bool // used in tests

func (c *Client) DeleteStale(ctx context.Context, msgs message.WriteDeleteStales) error {
	for _, msg := range msgs {
		tableName := msg.TableName
		source := msg.SourceName
		syncTime := msg.SyncTime
		var sb strings.Builder
		sb.WriteString("delete from ")
		sb.WriteString(sanitizeID(tableName))
		sb.WriteString(" where ")
		sb.WriteString(sanitizeID(schema.CqSourceNameColumn.Name))
		sb.WriteString(" = $1 and ")
		sb.WriteString(schema.CqSyncTimeColumn.Name)
		sb.WriteString(" < make_timestamp($2)")
		if err := c.exec(ctx, sb.String(), source, syncTime.UnixMicro()); err != nil {
			return err
		}
	}
	if delayAfterDeleteStale {
		time.Sleep(time.Second)
	}

	return nil
}

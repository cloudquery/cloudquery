package client

import (
	"context"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) deleteStale(ctx context.Context, tableName string, source string, syncTime time.Time) error {
	var sb strings.Builder
	sb.WriteString("delete from ")
	sb.WriteString(`"` + tableName + `"`)
	sb.WriteString(" where ")
	sb.WriteString(`"` + schema.CqSourceNameColumn.Name + `"`)
	sb.WriteString(" = $1 and datetime(")
	sb.WriteString(schema.CqSyncTimeColumn.Name)
	sb.WriteString(") < datetime($2)")
	sql := sb.String()
	if _, err := c.db.ExecContext(ctx, sql, source, syncTime); err != nil {
		return err
	}
	return nil
}

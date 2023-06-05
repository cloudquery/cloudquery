package client

import (
	"context"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	for _, table := range tables {
		var sb strings.Builder
		sb.WriteString("delete from ")
		sb.WriteString(sanitizeID(table.Name))
		sb.WriteString(" where ")
		sb.WriteString(sanitizeID(schema.CqSourceNameColumn.Name))
		sb.WriteString(" = $1 and ")
		sb.WriteString(schema.CqSyncTimeColumn.Name)
		sb.WriteString(" < to_timestamp($2)")
		sql := sb.String()
		if _, err := c.db.ExecContext(ctx, sql, source, syncTime.Unix()); err != nil {
			return err
		}
	}
	return nil
}

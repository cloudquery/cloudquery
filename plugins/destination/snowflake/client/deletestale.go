package client

import (
	"context"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	for _, table := range tables {
		var sb strings.Builder
		sb.WriteString("delete from ")
		sb.WriteString(table.Name)
		sb.WriteString(" where ")
		sb.WriteString(`"` + schema.CqSourceNameColumn.Name + `"`)
		sb.WriteString(" = ? and \"")
		sb.WriteString(schema.CqSyncTimeColumn.Name)
		sb.WriteString("\"::timestamp_tz < ?::timestamp_tz")
		sql := sb.String()
		if _, err := c.db.Exec(sql, source, syncTime); err != nil {
			return err
		}
	}
	return nil
}

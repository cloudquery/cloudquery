package client

import (
	"context"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v4"
)

func (p *Client) DeleteStale(ctx context.Context, table string, source string, syncTime time.Time) error {
	var sb strings.Builder
	sb.WriteString("delete from ")
	sb.WriteString(pgx.Identifier{table}.Sanitize())
	sb.WriteString(" where ")
	sb.WriteString(schema.CqSourceNameColumn.Name)
	sb.WriteString(" = $1 and ")
	sb.WriteString(schema.CqSyncTimeColumn.Name)
	sb.WriteString(" < $2")
	if _, err := p.conn.Exec(ctx, sb.String(), source, syncTime); err != nil {
		return err
	}
	return nil
}

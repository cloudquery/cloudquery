package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/jackc/pgx/v5"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Schemas, source string, syncTime time.Time) error {
	batch := &pgx.Batch{}
	for _, table := range tables {
		var sb strings.Builder
		sb.WriteString("delete from ")
		sb.WriteString(pgx.Identifier{schema.TableName(table)}.Sanitize())
		sb.WriteString(" where ")
		sb.WriteString(schema.CqSourceNameColumn.Name)
		sb.WriteString(" = $1 and ")
		sb.WriteString(schema.CqSyncTimeColumn.Name)
		sb.WriteString(" < $2")
		batch.Queue(sb.String(), source, syncTime)
	}
	br := c.conn.SendBatch(ctx, batch)
	if err := br.Close(); err != nil {
		return fmt.Errorf("failed to execute batch: %w", err)
	}
	return nil
}

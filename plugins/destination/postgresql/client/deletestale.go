package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/jackc/pgx/v5"
)

// DeleteStaleBatch deletes stale records from the destination table. It forms part of the writer.MixedBatchWriter interface.
func (c *Client) DeleteStaleBatch(ctx context.Context, messages message.WriteDeleteStales) error {
	batch := &pgx.Batch{}
	for _, msg := range messages {
		var sb strings.Builder
		sb.WriteString("delete from ")
		sb.WriteString(pgx.Identifier{msg.TableName}.Sanitize())
		sb.WriteString(" where ")
		sb.WriteString(schema.CqSourceNameColumn.Name)
		sb.WriteString(" = $1 and ")
		sb.WriteString(schema.CqSyncTimeColumn.Name)
		sb.WriteString(" < $2")
		batch.Queue(sb.String(), msg.SourceName, msg.SyncTime)
	}
	br := c.conn.SendBatch(ctx, batch)
	if err := br.Close(); err != nil {
		return fmt.Errorf("failed to execute batch: %w", err)
	}
	return nil
}

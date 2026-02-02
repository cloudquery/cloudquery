package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (c *Client) WriteTableBatch(ctx context.Context, tableName string, msgs message.WriteInserts) error {
	if len(msgs) == 0 {
		return nil
	}

	table, err := schema.NewTableFromArrowSchema(msgs[0].Record.Schema())
	if err != nil {
		return err
	}

	session := c.Session(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	rows := make([]map[string]any, 0, len(msgs))
	for i := range msgs {
		rows = append(rows, transformValues(msgs[i].Record)...)
	}

	var sb strings.Builder
	pks := table.PrimaryKeys()
	if len(pks) == 0 {
		sb.WriteString("UNWIND $rows AS row CREATE (t:")
		sb.WriteString(tableName)
		sb.WriteString(") SET t = row")
	} else {
		sb.WriteString("UNWIND $rows AS row MERGE (t:")
		sb.WriteString(tableName)
		sb.WriteString(" {")
		for i, column := range pks {
			if i != 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(column)
			sb.WriteString(": row.")
			sb.WriteString(column)
		}
		sb.WriteString("}) SET t = row")
	}
	stmt := sb.String()
	c.logger.Debug().Str("stmt", stmt).Any("rows", rows).Msg("Executing statement")
	if _, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, stmt, map[string]any{"rows": rows})
		return nil, err
	}); err != nil {
		return err
	}

	return session.Close(ctx)
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, msgs); err != nil {
		return err
	}
	if err := c.writer.Flush(ctx); err != nil {
		return fmt.Errorf("failed to flush: %w", err)
	}
	return nil
}

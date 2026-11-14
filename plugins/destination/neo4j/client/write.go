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

	stmt := writeStatement(tableName, table.PrimaryKeys())
	c.logger.Debug().Str("stmt", stmt).Any("rows", rows).Msg("Executing statement")
	if _, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, stmt, map[string]any{"rows": rows})
		return nil, err
	}); err != nil {
		return err
	}

	return session.Close(ctx)
}

func writeStatement(tableName string, primaryKeys []string) string {
	label := quoteIdentifier(tableName)
	if len(primaryKeys) == 0 {
		return "UNWIND $rows AS row CREATE (t:" + label + ") SET t = row"
	}
	keys := make([]string, len(primaryKeys))
	for i, column := range primaryKeys {
		key := quoteIdentifier(column)
		keys[i] = key + ": row." + key
	}
	return "UNWIND $rows AS row MERGE (t:" + label + " {" + strings.Join(keys, ", ") + "}) SET t = row"
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

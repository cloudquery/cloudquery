package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

const deleteCypher = "MATCH (n:%s) WHERE n._cq_source_name = $cq_source_name AND n._cq_sync_time < $cq_sync_time DETACH DELETE n"

func (c *Client) DeleteStale(ctx context.Context, msgs message.WriteDeleteStales) error {
	stmts := make([]string, len(msgs))
	for i, msg := range msgs {
		stmt, err := deleteStaleStatement(msg.TableName)
		if err != nil {
			return err
		}
		stmts[i] = stmt
	}
	session := c.Session(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	for i, msg := range msgs {
		if _, err := session.Run(ctx, stmts[i], map[string]any{"cq_source_name": msg.SourceName, "cq_sync_time": msg.SyncTime.Truncate(time.Microsecond)}); err != nil {
			return err
		}
	}
	return session.Close(ctx)
}

func deleteStaleStatement(tableName string) (string, error) {
	label, err := quoteIdentifier(tableName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(deleteCypher, label), nil
}

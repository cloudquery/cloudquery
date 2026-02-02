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
	session := c.Session(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	for _, msg := range msgs {
		stmt := fmt.Sprintf(deleteCypher, msg.TableName)
		if _, err := session.Run(ctx, stmt, map[string]any{"cq_source_name": msg.SourceName, "cq_sync_time": msg.SyncTime.Truncate(time.Microsecond)}); err != nil {
			return err
		}
	}
	return session.Close(ctx)
}

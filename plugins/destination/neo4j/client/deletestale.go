package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const deleteCypher = "MATCH (n:%s) WHERE n._cq_source_name = $cq_source_name AND n._cq_sync_time < $cq_sync_time DETACH DELETE n"

func (c *Client) DeleteStale(ctx context.Context, tables schema.Schemas, source string, syncTime time.Time) error {
	session := c.LoggedSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	for _, table := range tables {
		tableName := schema.TableName(table)
		stmt := fmt.Sprintf(deleteCypher, tableName)
		if _, err := session.Run(ctx, stmt, map[string]any{"cq_source_name": source, "cq_sync_time": syncTime.Truncate(time.Microsecond)}); err != nil {
			return err
		}
	}
	return session.Close(ctx)
}

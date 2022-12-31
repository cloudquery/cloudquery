package client

import (
	"context"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	session := c.client.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	rows := make([]map[string]interface{}, len(resources))
	for i, resource := range resources {
		rows[i] = make(map[string]interface{})
		for j, column := range table.Columns {
			rows[i][column.Name] = resource[j]
		}
	}
	var sb strings.Builder
	sb.WriteString("UNWIND $rows as row MERGE (t:")
	sb.WriteString(table.Name)
	sb.WriteString(" {")
	for i, column := range table.PrimaryKeys() {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(column)
		sb.WriteString(": row.")
		sb.WriteString(column)
	}
	sb.WriteString("}) SET t = row")
	stmt := sb.String()
	if _, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, stmt, map[string]interface{}{"rows": rows})
		if err != nil {
			return nil, err
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

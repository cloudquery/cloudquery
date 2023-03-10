package client

import (
	"context"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	session := c.LoggedSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	rows := make([]map[string]any, len(resources))
	for i, resource := range resources {
		rows[i] = make(map[string]any)
		for j, column := range table.Columns {
			rows[i][column.Name] = resource[j]
		}
	}
	var sb strings.Builder
	sb.WriteString("UNWIND $rows AS row MERGE (t:")
	sb.WriteString(table.Name)
	sb.WriteString(" {")
	pks := table.PrimaryKeys()
	if len(pks) == 0 {
		// If no primary keys are defined, use all columns
		pks = table.Columns.Names()
	}
	for i, column := range pks {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(column)
		sb.WriteString(": row.")
		sb.WriteString(column)
	}
	sb.WriteString("}) SET t = row")
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

package client

import (
	"context"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *arrow.Schema, records []arrow.Record) error {
	session := c.LoggedSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	tableName := schema.TableName(table)
	rows := make([]map[string]any, 0)
	for _, record := range records {
		rows = append(rows, transformValues(record)...)
		record.Release()
	}
	var sb strings.Builder
	sb.WriteString("UNWIND $rows AS row MERGE (t:")
	sb.WriteString(tableName)
	sb.WriteString(" {")
	pks := schema.PrimaryKeyIndices(table)
	if len(pks) == 0 {
		// If no primary keys are defined, use _cq_id
		pks = table.FieldIndices(schema.CqIDColumn.Name)
	}
	for i, columnIndice := range pks {
		if i != 0 {
			sb.WriteString(", ")
		}
		column := table.Field(columnIndice).Name
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

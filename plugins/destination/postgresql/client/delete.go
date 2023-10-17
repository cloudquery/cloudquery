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

func (c *Client) DeleteRecordsBatch(ctx context.Context, messages message.WriteDeleteRecords) error {
	batch := &pgx.Batch{}
	for _, msg := range messages {
		sql := generateDeleteCTE(msg.DeleteRecord)
		vals := extractPredicateValues(msg.DeleteRecord.WhereClause)
		batch.Queue(sql, vals...)
	}
	br := c.conn.SendBatch(ctx, batch)
	if err := br.Close(); err != nil {
		return fmt.Errorf("failed to execute batch: %w", err)
	}
	return nil
}

func generateInitialDelete(tableName string, whereClause message.PredicateGroups) string {
	var sb strings.Builder
	sb.WriteString("DELETE from ")
	sb.WriteString(pgx.Identifier{tableName}.Sanitize())
	if len(whereClause) > 0 {
		sb.WriteString(" where ")
		counter := 1
		for i, predicateGroup := range whereClause {
			if len(predicateGroup.Predicates) == 0 {
				continue
			}
			sb.WriteString("( ")
			for i, predicate := range predicateGroup.Predicates {
				if i > 0 {
					sb.WriteString(fmt.Sprintf(" %s ", predicateGroup.GroupingType))
				}
				sb.WriteString(pgx.Identifier{predicate.Column}.Sanitize())
				sb.WriteString(fmt.Sprintf(" = $%d", counter))
				counter++
			}
			sb.WriteString(" )")
			if i < len(whereClause)-1 {
				sb.WriteString(" AND ")
			}
		}
	}

	sb.WriteString(" RETURNING *")
	return sb.String()
}

func generateRelationsDelete(tableRelation message.TableRelation) string {
	var sb strings.Builder
	sb.WriteString("DELETE from ")
	sb.WriteString(pgx.Identifier{tableRelation.TableName}.Sanitize())
	sb.WriteString(" where ")
	sb.WriteString(pgx.Identifier{schema.CqParentIDColumn.Name}.Sanitize())
	sb.WriteString(" in (select ")
	sb.WriteString(pgx.Identifier{schema.CqIDColumn.Name}.Sanitize())
	sb.WriteString(" from ")
	sb.WriteString(pgx.Identifier{tableRelation.ParentTable + "_CTE"}.Sanitize())
	sb.WriteString(")")
	return sb.String()
}

func generateDeleteCTE(deleteRecord message.DeleteRecord) string {
	tables := make([]string, len(deleteRecord.TableRelations))
	initialDelete := generateInitialDelete(deleteRecord.TableName, deleteRecord.WhereClause)
	var sb strings.Builder
	sb.WriteString("WITH ")
	sb.WriteString(pgx.Identifier{deleteRecord.TableName + "_CTE"}.Sanitize())
	sb.WriteString(" AS (")
	sb.WriteString(initialDelete)
	sb.WriteString(") ")

	for i, tableRelation := range deleteRecord.TableRelations {
		sb.WriteString(", ")
		sb.WriteString(pgx.Identifier{tableRelation.TableName + "_CTE"}.Sanitize())
		sb.WriteString(" AS (")
		sb.WriteString(generateRelationsDelete(tableRelation))
		sb.WriteString(" RETURNING ")
		sb.WriteString(pgx.Identifier{schema.CqIDColumn.Name}.Sanitize())
		sb.WriteString(") ")
		tables[i] = tableRelation.TableName
	}
	for _, table := range tables {
		sb.WriteString("Select count(*) from ")
		sb.WriteString(pgx.Identifier{table + "_CTE"}.Sanitize())
		sb.WriteString(" UNION ALL ")
	}
	sb.WriteString("Select count(*) from ")
	sb.WriteString(pgx.Identifier{deleteRecord.TableName + "_CTE"}.Sanitize())

	return sb.String()
}

func extractPredicateValues(where message.PredicateGroups) []any {
	predicateCount := 0
	for _, predicateGroup := range where {
		predicateCount += len(predicateGroup.Predicates)
	}
	results := make([]any, predicateCount)
	counter := 0
	for _, predicateGroup := range where {
		for _, predicate := range predicateGroup.Predicates {
			col := predicate.Record.Column(0)
			transformed := transformArr(col)
			results[counter] = transformed[0]
			counter++
		}
	}
	return results
}

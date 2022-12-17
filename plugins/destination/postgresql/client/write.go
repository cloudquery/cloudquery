package client

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

var cqStatusToPgStatus = map[schema.Status]pgtype.Status{
	schema.Null:      pgtype.Null,
	schema.Undefined: pgtype.Null,
	schema.Present:   pgtype.Present,
}


// Postgres best performance for CQ per our benchmark is achieved with batch that includes different tables
// given we can't predict how many objects will be in each table.
// This is why if we encounter an error we can't know which object/table caused and related the error to the object
// we we need to find the error and exit/stop the sync.
// Also, it is not advice to split the batch into different tables (for postgresql) as
// this will mean increased number of goroutines and will mean more connections which PostgreSQL is not great with.
// Also, I've benchmarked CopyFrom and batch and seems there is no significant difference in CQ case.
// maybe in the future we can provide a flag to switch batching mechanism.
func (c *Client) findErrorInBatch(ctx context.Context, items []*batchItem) error {
	for _, item := range items {
		_, err := c.conn.Exec(ctx, item.sql, item.arguments...)
		if err != nil {
			return fmt.Errorf("failed to insert into table %s with sql %s: %w", item.table, item.sql, err)
		}
	}
	return nil
}

type batchItem struct {
	table string
	sql     string
	arguments []interface{}
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *plugins.ClientResource) error {
	var sql string
	batch := &pgx.Batch{}
	batchItems := make([]*batchItem, 0, c.batchSize)
	// resources := 

	for r := range res {
		table := tables.Get(r.TableName)
		if table == nil {
			panic(fmt.Errorf("table %s not found", r.TableName))
		}
		if c.spec.WriteMode == specs.WriteModeAppend {
			sql = c.insert(table)
		} else {
			sql = c.upsert(table)
		}
		batchItems = append(batchItems, &batchItem{
			table: r.TableName,
			sql: sql,
			arguments: r.Data,
		})
		batch.Queue(sql, r.Data...)
		if batch.Len() >= c.batchSize {
			br := c.conn.SendBatch(ctx, batch)
			if batchErr := br.Close(); batchErr != nil {
				if err := c.findErrorInBatch(ctx, batchItems); err != nil {
					return err
				} else {
					// this should never happen
					return fmt.Errorf("failed to execute batch and was unable to pinpoint table: %w", batchErr)
				}
			}
			atomic.AddUint64(&c.metrics.Writes, uint64(c.batchSize))
			batch = &pgx.Batch{}
			batchItems = make([]*batchItem, 0, c.batchSize)
		}
	}

	if batch.Len() > 0 {
		br := c.conn.SendBatch(ctx, batch)
		if err := br.Close(); err != nil {
			if batchErr := br.Close(); batchErr != nil {
				if err := c.findErrorInBatch(ctx, batchItems); err != nil {
					return err
				} else {
					// this should never happen
					return fmt.Errorf("failed to execute batch and was unable to pinpoint table: %w", batchErr)
				}
			}
		}
		atomic.AddUint64(&c.metrics.Writes, uint64(batch.Len()))
	}

	return nil
}

func (*Client) insert(table *schema.Table) string {
	var sb strings.Builder
	sb.WriteString("insert into ")
	sb.WriteString(pgx.Identifier{table.Name}.Sanitize())
	sb.WriteString(" (")
	columns := table.Columns
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(pgx.Identifier{c.Name}.Sanitize())
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i := range columns {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	return sb.String()
}

func (c *Client) upsert(table *schema.Table) string {
	var sb strings.Builder

	sb.WriteString(c.insert(table))
	columns := table.Columns
	columnsLen := len(columns)

	constraintName := fmt.Sprintf("%s_cqpk", table.Name)
	sb.WriteString(" on conflict on constraint ")
	sb.WriteString(constraintName)
	sb.WriteString(" do update set ")
	for i, column := range columns {
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		sb.WriteString("=excluded.") // excluded references the new values
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString("")
		}
	}

	return sb.String()
}

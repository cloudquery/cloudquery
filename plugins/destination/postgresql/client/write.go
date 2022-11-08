package client

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

var cqStatusToPgStatus = map[schema.Status]pgtype.Status{
	schema.Null:      pgtype.Null,
	schema.Undefined: pgtype.Null,
	schema.Present:   pgtype.Present,
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *plugins.ClientResource) error {
	var sql string
	batch := &pgx.Batch{}

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

		batch.Queue(sql, r.Data...)
		if batch.Len() >= c.batchSize {
			br := c.conn.SendBatch(ctx, batch)
			if err := br.Close(); err != nil {
				if _, ok := err.(*pgconn.PgError); !ok {
					// not recoverable error
					return fmt.Errorf("failed to execute batch: %w", err)
				}
				atomic.AddUint64(&c.metrics.Errors, 1)
				c.logger.Error().Err(err).Msgf("failed to execute batch with pgerror")
			}
			atomic.AddUint64(&c.metrics.Writes, uint64(c.batchSize))
			batch = &pgx.Batch{}
		}
	}

	if batch.Len() > 0 {
		br := c.conn.SendBatch(ctx, batch)
		if err := br.Close(); err != nil {
			if _, ok := err.(*pgconn.PgError); !ok {
				// no recoverable error
				return fmt.Errorf("failed to execute batch: %w", err)
			}
			c.logger.Error().Err(err).Msgf("failed to execute batch with pgerror")
		}
		atomic.AddUint64(&c.metrics.Writes, uint64(c.batchSize))
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

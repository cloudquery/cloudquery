package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// WriteTable inserts records into the destination table.
// Part of the streamingbatchwriter.Client interface.
func (c *Client) WriteTable(ctx context.Context, messages <-chan *message.WriteInsert) error {
	var table *schema.Table
	var pgTable *schema.Table
	var err error

	var rows [][]any
	for msg := range messages {
		if table == nil {
			table = msg.GetTable()
			pgTable, err = c.getDBTable(ctx, table.Name)
			if err != nil {
				return fmt.Errorf("failed getting postgres table %s: %w", table.Name, err)
			}
			if pgTable == nil {
				return fmt.Errorf("missing expected postgres table %s", table.Name)
			}
			table = c.normalizeTable(table, pgTable)
		}

		rows = append(rows, transformValues(msg.Record)...)
	}

	if len(rows) == 0 {
		return nil // no data to insert
	}

	if len(table.PrimaryKeysIndexes()) > 0 {
		err = c.insert(ctx, table, rows)
	} else {
		err = c.copyFrom(ctx, table, rows)
	}

	if err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) {
			// not recoverable error
			return fmt.Errorf("failed to execute batch: %w", err)
		}
		return fmt.Errorf("failed to execute batch with pgerror: %s: %w", pgErrToStr(pgErr), err)
	}

	return nil
}

func (c *Client) insert(ctx context.Context, table *schema.Table, rows [][]any) error {
	var sb strings.Builder
	sb.WriteString("insert into ")
	sb.WriteString(pgx.Identifier{table.Name}.Sanitize())
	sb.WriteString(" (")
	for i, c := range table.Columns {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(pgx.Identifier{c.Name}.Sanitize())
	}
	sb.WriteString(") values (")
	for i := range table.Columns {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("$%d", i+1))
	}

	constraintName := table.PkConstraintName
	sb.WriteString(") on conflict on constraint ")
	sb.WriteString(pgx.Identifier{constraintName}.Sanitize())
	sb.WriteString(" do update set ")
	for i, column := range table.Columns {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		sb.WriteString("=excluded.") // excluded references the new values
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
	}

	query := sb.String()

	batch := new(pgx.Batch)
	for _, row := range rows {
		batch.Queue(query, row...)
	}

	return c.conn.SendBatch(ctx, batch).Close()
}

func (c *Client) copyFrom(ctx context.Context, table *schema.Table, rows [][]any) error {
	inserted, err := c.conn.CopyFrom(ctx, pgx.Identifier{table.Name}, table.Columns.Names(), pgx.CopyFromRows(rows))
	c.logger.Debug().
		Err(err).
		Str("table", table.Name).
		Int("rows", len(rows)).
		Int64("inserted", inserted).
		Msg("copied from rows")
	return err
}

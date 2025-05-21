package client

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (c *Client) pgTables(ctx context.Context) (map[string]struct{}, error) {
	c.pgTablesToPKConstraintsMu.RLock()
	// This happens when the CLI was invoked with `sync --no-migrate`
	if len(c.pgTablesToPKConstraints) == 0 {
		c.pgTablesToPKConstraintsMu.RUnlock()
		// listTables populates c.pgTablesToPKConstraints
		_, err := c.listTables(ctx)
		if err != nil {
			return nil, err
		}
		c.pgTablesToPKConstraintsMu.RLock()
	}

	list := make(map[string]struct{}, len(c.pgTablesToPKConstraints))
	for k := range c.pgTablesToPKConstraints {
		list[k] = struct{}{}
	}
	c.pgTablesToPKConstraintsMu.RUnlock()
	return list, nil
}

// InsertBatch inserts records into the destination table. It forms part of the writer.MixedBatchWriter interface.
func (c *Client) InsertBatch(ctx context.Context, messages message.WriteInserts) error {
	pgTables, err := c.pgTables(ctx)
	if err != nil {
		return err
	}

	c.bm = &batchManager{
		pgxBatch: new(pgx.Batch),
	}

	// Queries cache.
	// We may consider LRU cache in the future, but even for 10K records it may be OK to just save.
	// queries := make(map[string]string, 100)

	for _, msg := range messages {
		r := msg.Record

		table := c.normalizeTable(msg.GetTable())

		if _, ok := pgTables[table.Name]; !ok {
			return fmt.Errorf("table %s not found", table.Name)
		}

		rows := c.transformValues(r)
		for _, rowVals := range rows {
			err := c.addRow(ctx, table, rowVals)
			if err != nil {
				return fmt.Errorf("failed to add row: %w", err)
			}
		}
	}

	return c.flushBatch(ctx)
}

func (*Client) insert(table *schema.Table, rowCount int) string {
	var sb strings.Builder
	tableName := table.Name
	sb.WriteString("insert into ")
	sb.WriteString(pgx.Identifier{tableName}.Sanitize())
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
	totalArgs := len(columns) * rowCount
	for i := range totalArgs + 1 {
		if i == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("$%d", i))
		if i < totalArgs {
			if len(columns) == 1 || (i > 1 && (i)%len(columns) == 0) {
				sb.WriteString("),(")
			} else {
				sb.WriteString(",")
			}
		} else {
			sb.WriteString(")")
		}
	}
	return sb.String()
}

func (c *Client) upsert(table *schema.Table, rowCount int) string {
	if c.pgType == pgTypeCrateDB {
		return c.upsertCrateDB(table)
	}
	var sb strings.Builder
	sb.WriteString(c.insert(table, rowCount))
	columns := table.Columns
	columnsLen := len(columns)
	constraintName := table.PkConstraintName
	sb.WriteString(" on conflict on constraint ")
	sb.WriteString(pgx.Identifier{constraintName}.Sanitize())
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

// CrateDB differs from normal Postgres in that it does not support constraint-name
// based upserts, and errors out if a primary key is written to during upsert.
func (c *Client) upsertCrateDB(table *schema.Table) string {
	var sb strings.Builder
	sb.WriteString(c.insert(table, 1))
	columns := table.Columns
	if len(table.PrimaryKeysIndexes()) == len(table.Columns) {
		sb.WriteString(" on conflict do nothing")
		return sb.String()
	}
	pks := table.PrimaryKeys()
	sb.WriteString(fmt.Sprintf(" on conflict (%s) ", strings.Join(pks, ",")))

	sb.WriteString(" do update set ")
	cols := make([]string, 0, len(columns)-len(pks))
	for _, column := range columns {
		if column.PrimaryKey {
			continue
		}
		cols = append(cols, pgx.Identifier{column.Name}.Sanitize()+"=excluded."+pgx.Identifier{column.Name}.Sanitize())
	}
	sb.WriteString(strings.Join(cols, ","))

	return sb.String()
}

func pgErrToStr(err *pgconn.PgError) string {
	var sb strings.Builder
	sb.WriteString("severity: ")
	sb.WriteString(err.Severity)
	sb.WriteString(", code: ")
	sb.WriteString(err.Code)
	sb.WriteString(", message: ")
	sb.WriteString(err.Message)
	sb.WriteString(", detail :")
	sb.WriteString(err.Detail)
	sb.WriteString(", hint: ")
	sb.WriteString(err.Hint)
	sb.WriteString(", position: ")
	sb.WriteString(strconv.FormatInt(int64(err.Position), 10))
	sb.WriteString(", internal_position: ")
	sb.WriteString(strconv.FormatInt(int64(err.InternalPosition), 10))
	sb.WriteString(", internal_query: ")
	sb.WriteString(err.InternalQuery)
	sb.WriteString(", where: ")
	sb.WriteString(err.Where)
	sb.WriteString(", schema_name: ")
	sb.WriteString(err.SchemaName)
	sb.WriteString(", table_name: ")
	sb.WriteString(err.TableName)
	sb.WriteString(", column_name: ")
	sb.WriteString(err.ColumnName)
	sb.WriteString(", data_type_name: ")
	sb.WriteString(err.DataTypeName)
	sb.WriteString(", constraint_name: ")
	sb.WriteString(err.ConstraintName)
	sb.WriteString(", file: ")
	sb.WriteString(err.File)
	sb.WriteString(", line: ")
	sb.WriteString(strconv.FormatUint(uint64(err.Line), 10))
	sb.WriteString(", routine: ")
	sb.WriteString(err.Routine)
	return sb.String()
}

func (c *Client) addRow(ctx context.Context, table *schema.Table, row []any) error {
	for i, entry := range c.bm.batch {
		if entry.table != table {
			continue
		}
		if len(entry.data)*entry.columnCount >= 60000 {
			continue
		}
		c.bm.batch[i].data = append(c.bm.batch[i].data, row)
		c.bm.batch[i].recordCount++
		c.bm.recordCount++
		if c.bm.recordCount >= int(c.batchSize) {
			err := c.flushBatch(ctx)
			if err != nil {
				return fmt.Errorf("failed to flush batch: %w", err)
			}
			c.bm.batch = nil
			c.bm.recordCount = 0
			c.bm.pgxBatch = new(pgx.Batch)
		}
		return nil
	}
	c.bm.batch = append(c.bm.batch, batchEntry{
		table:       table,
		columnCount: len(row),
		recordCount: 1,
		data:        [][]any{row},
	})
	c.bm.recordCount++

	return nil
}

func (c *Client) flushBatch(ctx context.Context) error {
	if len(c.bm.batch) == 0 {
		return nil
	}
	for _, entry := range c.bm.batch {
		sql := ""
		if len(entry.table.PrimaryKeysIndexes()) > 0 {
			sql = c.upsert(entry.table, entry.recordCount)
		} else {
			sql = c.insert(entry.table, entry.recordCount)
		}
		print(sql)
		c.bm.pgxBatch.Queue(sql, entry.getValues()...)
	}
	err := c.conn.SendBatch(ctx, c.bm.pgxBatch).Close()
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return fmt.Errorf("failed to execute batch with pgerror: %s: %w", pgErrToStr(pgErr), err)
	}

	// not recoverable error
	return fmt.Errorf("failed to execute batch: %w", err)
}

func (be batchEntry) getValues() []any {
	values := make([]any, 0, len(be.data))
	for _, row := range be.data {
		values = append(values, row...)
	}
	return values
}

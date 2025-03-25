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

	batch := new(pgx.Batch)

	// Queries cache.
	// We may consider LRU cache in the future, but even for 10K records it may be OK to just save.
	queries := make(map[string]string, 100)

	for _, msg := range messages {
		r := msg.Record
		md := r.Schema().Metadata()
		tableName, ok := md.GetValue(schema.MetadataTableName)
		if !ok {
			return errors.New("table name not found in metadata")
		}

		if _, ok := pgTables[tableName]; !ok {
			return fmt.Errorf("table %s not found", tableName)
		}

		sql, ok := queries[tableName]
		if !ok {
			// cache the query
			table := c.normalizeTable(msg.GetTable())
			if len(table.PrimaryKeysIndexes()) > 0 {
				sql = c.upsert(table)
			} else {
				sql = c.insert(table)
			}
			queries[tableName] = sql
		}

		rows := c.transformValues(r)
		for _, rowVals := range rows {
			batch.Queue(sql, rowVals...)
		}
		if int64(batch.Len()) >= c.batchSize {
			if err := c.flushBatch(ctx, batch); err != nil {
				return err
			}
			batch = new(pgx.Batch)
		}
	}

	return c.flushBatch(ctx, batch)
}

func (c *Client) flushBatch(ctx context.Context, batch *pgx.Batch) error {
	if batch.Len() == 0 {
		return nil
	}
	err := c.conn.SendBatch(ctx, batch).Close()
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

func (*Client) insert(table *schema.Table) string {
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
	if c.pgType == pgTypeCrateDB {
		return c.upsertCrateDB(table)
	}
	var sb strings.Builder
	sb.WriteString(c.insert(table))
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
	sb.WriteString(c.insert(table))
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

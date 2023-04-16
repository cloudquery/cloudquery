package client

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

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

func (c *Client) Write(ctx context.Context, tables schema.Schemas, res <-chan arrow.Record) error {
	var sql string
	batch := &pgx.Batch{}
	pgTables, err := c.listPgTables(ctx, tables)
	if err != nil {
		return err
	}
	tables = c.normalizeTables(tables, pgTables)
	if err != nil {
		return err
	}
	for r := range res {
		tableName := schema.TableName(r.Schema())
		table := tables.SchemaByName(tableName)
		if table == nil {
			panic(fmt.Errorf("table %s not found", tableName))
		}
		if c.spec.WriteMode == specs.WriteModeAppend {
			sql = c.insert(table)
		} else {
			if len(schema.PrimaryKeyIndices(table)) > 0 {
				sql = c.upsert(table)
			} else {
				sql = c.insert(table)
			}
		}
		rows := transformValues(r)
		for _, rowVals := range rows {
			batch.Queue(sql, rowVals...)
		}
		batchSize := batch.Len()
		if batchSize >= c.batchSize {
			br := c.conn.SendBatch(ctx, batch)
			if err := br.Close(); err != nil {
				var pgErr *pgconn.PgError
				if !errors.As(err, &pgErr) {
					// not recoverable error
					return fmt.Errorf("failed to execute batch: %w", err)
				}
				return fmt.Errorf("failed to execute batch with pgerror: %s: %w", pgErrToStr(pgErr), err)
			}
			atomic.AddUint64(&c.metrics.Writes, uint64(batchSize))
			batch = &pgx.Batch{}
		}
	}

	batchSize := batch.Len()
	if batchSize > 0 {
		br := c.conn.SendBatch(ctx, batch)
		if err := br.Close(); err != nil {
			var pgErr *pgconn.PgError
			if !errors.As(err, &pgErr) {
				// not recoverable error
				return fmt.Errorf("failed to execute batch: %w", err)
			}
			return fmt.Errorf("failed to execute batch with pgerror: %s: %w", pgErrToStr(pgErr), err)
		}
		atomic.AddUint64(&c.metrics.Writes, uint64(batchSize))
	}

	return nil
}

func (*Client) insert(table *arrow.Schema) string {
	var sb strings.Builder
	tableName := schema.TableName(table)
	sb.WriteString("insert into ")
	sb.WriteString(pgx.Identifier{tableName}.Sanitize())
	sb.WriteString(" (")
	columns := table.Fields()
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

func (c *Client) upsert(table *arrow.Schema) string {
	var sb strings.Builder

	sb.WriteString(c.insert(table))
	columns := table.Fields()
	columnsLen := len(columns)

	constraintName, ok := table.Metadata().GetValue(schema.MetadataConstraintName)
	if !ok {
		panic(fmt.Errorf("constraint_name not found in table metadata"))
	}
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

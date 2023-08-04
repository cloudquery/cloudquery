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

// WriteTable inserts records into the destination table.
// Part of the streamingbatchwriter.Client interface.
func (c *Client) WriteTable(ctx context.Context, messages <-chan *message.WriteInsert) error {
	var table *schema.Table
	var pgTable *schema.Table
	var err error

	var sql string
	batch := &pgx.Batch{}
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

			if len(table.PrimaryKeysIndexes()) > 0 {
				sql = c.upsert(table)
			} else {
				sql = c.insert(table)
			}
		}

		for _, row := range transformValues(msg.Record) {
			batch.Queue(sql, row...)
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
	}
	return nil
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

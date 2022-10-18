package client

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgx/v4"
)

func (c *Client) Write(ctx context.Context, table string, data map[string]interface{}) error {
	var sql string
	var values []interface{}

	if c.spec.WriteMode == specs.WriteModeAppend {
		sql, values = insert(table, data)
	} else {
		sql, values = upsert(table, data)
	}
	c.batch.Queue(sql, values...)
	if c.batch.Len() >= c.batchSize {
		br := c.conn.SendBatch(ctx, c.batch)
		if err := br.Close(); err != nil {
			c.batch = &pgx.Batch{}
			return fmt.Errorf("failed to execute batch: %v", err)
		}
		c.batch = &pgx.Batch{}
	}
	return nil
}

func insert(table string, data map[string]interface{}) (string, []any) {
	sb := new(strings.Builder)
	_, values := header(sb, table, data)
	return sb.String(), values
}

func upsert(table string, data map[string]interface{}) (string, []any) {
	sb := new(strings.Builder)
	columns, values := header(sb, table, data)

	constraintName := table + `_cqpk`
	sb.WriteString(" on conflict on constraint ")
	sb.WriteString(constraintName)

	sb.WriteString(" do update set ")
	excluded := func(column string) {
		sb.WriteString(column + `=excluded.` + column) // excluded references the new values
	}
	for _, column := range columns[:len(columns)-1] {
		excluded(column)
		sb.WriteString(",")
	}
	excluded(columns[len(columns)-1])

	return sb.String(), values
}

func header(sb *strings.Builder, table string, data map[string]any) (columns []string, values []any) {
	columns = make([]string, 0, len(data))
	values = make([]any, 0, len(data))

	// Sort the columns prior to adding data to columns and values arrays
	// Columns need to be in the same order so that the query can be cached during the statement preparation stage
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		columns = append(columns, pgx.Identifier{key}.Sanitize())
		values = append(values, data[key])
	}

	sb.WriteString("insert into ")
	sb.WriteString(table)
	sb.WriteString(" (")
	sb.WriteString(strings.Join(columns, ","))
	sb.WriteString(") values (")
	for i := 1; i < len(values); i++ {
		sb.WriteString(`$` + strconv.Itoa(i) + `,`)
	}
	sb.WriteString(`$` + strconv.Itoa(len(values)) + `)`)
	return columns, values
}

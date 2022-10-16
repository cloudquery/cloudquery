package client

import (
	"context"
	"fmt"
	"sort"
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

func insert(table string, data map[string]interface{}) (string, []interface{}) {
	var sb strings.Builder

	columns := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))

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
	sort.Strings(columns)
	for i, c := range columns {
		sb.WriteString(c)
		// sb.WriteString("::" + SchemaTypeToPg())
		if i < len(columns)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i := range values {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < len(values)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	return sb.String(), values
}

func upsert(table string, data map[string]interface{}) (string, []interface{}) {
	var sb strings.Builder

	columns := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))
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
	sort.Strings(columns)
	for i, c := range columns {
		sb.WriteString(c)
		// sb.WriteString("::" + SchemaTypeToPg())
		if i < len(columns)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i := range values {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < len(values)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	constraintName := fmt.Sprintf("%s_cqpk", table)
	sb.WriteString(" on conflict on constraint ")
	sb.WriteString(constraintName)
	sb.WriteString(" do update set ")
	for i, column := range columns {
		sb.WriteString(column)
		sb.WriteString("=excluded.") // excluded references the new values
		sb.WriteString(column)
		if i < len(columns)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString("")
		}
	}

	return sb.String(), values
}

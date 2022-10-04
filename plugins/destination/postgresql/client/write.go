package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgx/v4"
)

func (p *Client) Write(ctx context.Context, table string, data map[string]interface{}) error {
	var sql string
	var values []interface{}
	if p.spec.WriteMode == specs.WriteModeAppend {
		sql, values = insert(table, data)
	} else {
		sql, values = upsert(table, data)
	}
	_, err := p.conn.Exec(ctx, sql, values...)
	if err != nil {
		return fmt.Errorf("failed to insert data with sql '%s': %w", sql, err)
	}
	return nil
}

func insert(table string, data map[string]interface{}) (string, []interface{}) {
	var sb strings.Builder

	columns := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))
	for c, v := range data {
		columns = append(columns, pgx.Identifier{c}.Sanitize())
		values = append(values, v)
	}
	sb.WriteString("insert into ")
	sb.WriteString(table)
	sb.WriteString(" (")
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
	for c, v := range data {
		columns = append(columns, pgx.Identifier{c}.Sanitize())
		values = append(values, v)
	}

	sb.WriteString("insert into ")
	sb.WriteString(table)
	sb.WriteString(" (")
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

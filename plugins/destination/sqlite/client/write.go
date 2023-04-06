package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan arrow.Record) error {
	var sql string
	for r := range res {
		tableName, ok := r.Schema().Metadata().GetValue(schema.MetadataTableName)
		if !ok {
			return fmt.Errorf("failed to get table name from record metadata")
		}
		table := tables.Get(tableName)
		if c.spec.WriteMode == specs.WriteModeAppend {
			sql = c.insert(table)
		} else {
			sql = c.upsert(table)
		}
		vals, err := transformRecord(r)
		if err != nil {
			return fmt.Errorf("failed to transform record: %w", err)
		}
		for _, v := range vals {
			if _, err := c.db.Exec(sql, v...); err != nil {
				return fmt.Errorf("failed to execute '%s': %w", sql, err)
			}
		}
	}

	return nil
}

func (*Client) insert(table *schema.Table) string {
	var sb strings.Builder
	sb.WriteString("insert into ")
	sb.WriteString(`"` + table.Name + `"`)
	sb.WriteString(" (")
	columns := table.Columns
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(`"` + c.Name + `"`)
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

func (*Client) upsert(table *schema.Table) string {
	var sb strings.Builder
	sb.WriteString("insert or replace into ")
	sb.WriteString(`"` + table.Name + `"`)
	sb.WriteString(" (")
	columns := table.Columns
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(`"` + c.Name + `"`)
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

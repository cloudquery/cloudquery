package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *destination.ClientResource) error {
	var sql string
	for r := range res {
		table := tables.Get(r.TableName)
		if c.spec.WriteMode == specs.WriteModeAppend {
			sql = c.insertSQL(table, r.Data)
			expanded := expandData(table, r.Data)
			if _, err := c.db.Exec(sql, expanded...); err != nil {
				return fmt.Errorf("failed to execute '%s': %w", sql, err)
			}
		} else {
			err := c.upsert(table, r.Data)
			if err != nil {
				return fmt.Errorf("failed to execute upsert: %w", err)
			}
		}
	}
	return nil
}

func expandData(table *schema.Table, data []any) []any {
	var expanded []any
	for i, d := range data {
		if isArray(table.Columns[i]) {
			expanded = append(expanded, toArray(table.Columns[i], d)...)
		} else {
			expanded = append(expanded, d)
		}
	}
	return expanded
}

func (c *Client) insertSQL(table *schema.Table, data []any) string {
	var sb strings.Builder
	sb.WriteString("insert into ")
	c.insertQuery(&sb, table, data)
	return sb.String()
}

func (c *Client) upsert(table *schema.Table, data []any) error {
	// At time of writing (March 2023), duckdb does not support updating list columns.
	// As a workaround, we delete the row and insert it again. This makes it non-atomic, unfortunately,
	// but this is unavoidable until support is added to duckdb itself.
	// See https://github.com/duckdb/duckdb/blob/c5d9afb97bbf0be12216f3b89ae3131afbbc3156/src/storage/table/list_column_data.cpp#L243-L251
	var sb strings.Builder
	if len(table.PrimaryKeys()) > 0 {
		sb.WriteString("delete from ")
		sb.WriteString(`"` + table.Name + `"`)
		sb.WriteString(" where ")
		pks := table.PrimaryKeys()
		pkData := make([]any, len(pks))
		for i, k := range pks {
			col := table.Columns.Get(k)
			sb.WriteString(`"` + col.Name + `"`)
			sb.WriteString(" = ")
			sb.WriteString(fmt.Sprintf("$%d ", i+1))
			if i < len(pks)-1 {
				sb.WriteString("and ")
			}
			pkData[i] = data[table.Columns.Index(k)]
		}
		sql := sb.String()
		if _, err := c.db.Exec(sql, pkData...); err != nil {
			return fmt.Errorf("failed to execute '%s': %w", sql, err)
		}
		sb.Reset()
	}
	sb.WriteString("insert into ")
	c.insertQuery(&sb, table, data)
	expanded := expandData(table, data)
	sql := sb.String()
	if _, err := c.db.Exec(sql, expanded...); err != nil {
		return fmt.Errorf("failed to execute '%s': %w", sql, err)
	}
	return nil
}

func (*Client) insertQuery(sb *strings.Builder, table *schema.Table, data []any) {
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
	counter := 0
	for i := range columns {
		if isArray(table.Columns[i]) {
			n := len(toArray(table.Columns[i], data[i]))
			sb.WriteString("[")
			for j := 0; j < n; j++ {
				sb.WriteString(fmt.Sprintf("$%d", counter+1))
				counter++
				if j < n-1 {
					sb.WriteString(",")
				}
			}
			sb.WriteString("]")
		} else {
			sb.WriteString(fmt.Sprintf("$%d", counter+1))
			counter++
		}
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
}

func isArray(col schema.Column) bool {
	switch col.Type {
	case schema.TypeIntArray, schema.TypeStringArray, schema.TypeCIDRArray, schema.TypeInetArray, schema.TypeUUIDArray, schema.TypeMacAddrArray:
		return true
	}
	return false
}

func toArray(col schema.Column, data any) []any {
	switch col.Type {
	case schema.TypeIntArray:
		if v, ok := data.([]int64); ok {
			a := make([]any, len(v))
			for i := range v {
				a[i] = v[i]
			}
			return a
		}
		return nil
	case schema.TypeStringArray, schema.TypeCIDRArray, schema.TypeInetArray, schema.TypeUUIDArray, schema.TypeMacAddrArray:
		if v, ok := data.([]string); ok {
			a := make([]any, len(v))
			for i := range v {
				a[i] = v[i]
			}
			return a
		}
		return nil
	}
	panic("not a supported array type: " + col.Type.String() + " for column: " + col.Name)
}

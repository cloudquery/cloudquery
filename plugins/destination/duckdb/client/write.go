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
			sql = c.insert(table, r.Data)
		} else {
			sql = c.upsert(table, r.Data)
		}
		expanded := expandData(table, r.Data)
		if _, err := c.db.Exec(sql, expanded...); err != nil {
			return fmt.Errorf("failed to execute '%s': %w", sql, err)
		}
	}
	return nil
}

func expandData(table *schema.Table, data []any) []any {
	var expanded []any
	for i, d := range data {
		if isArray(table.Columns[i]) {
			for _, s := range d.([]string) {
				expanded = append(expanded, s)
			}
		} else {
			expanded = append(expanded, d)
		}
	}
	return expanded
}

func (c *Client) insert(table *schema.Table, data []any) string {
	var sb strings.Builder
	sb.WriteString("insert into ")
	return c.insertQuery(&sb, table, data)
}

func (c *Client) upsert(table *schema.Table, data []any) string {
	var sb strings.Builder
	sb.WriteString("insert or replace into ")
	return c.insertQuery(&sb, table, data)
}

func (*Client) insertQuery(sb *strings.Builder, table *schema.Table, data []any) string {
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
			n := arrayLength(table.Columns[i], data[i])
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
	return sb.String()
}

func isArray(col schema.Column) bool {
	switch col.Type {
	case schema.TypeStringArray, schema.TypeCIDRArray, schema.TypeIntArray, schema.TypeInetArray, schema.TypeUUIDArray, schema.TypeMacAddrArray:
		return true
	}
	return false
}

func arrayLength(col schema.Column, data any) int {
	switch col.Type {
	case schema.TypeStringArray, schema.TypeCIDRArray, schema.TypeIntArray, schema.TypeInetArray, schema.TypeUUIDArray, schema.TypeMacAddrArray:
		return len(data.([]string))
	}
	panic("not a supported array type: " + col.Type.String() + " for column: " + col.Name)
}

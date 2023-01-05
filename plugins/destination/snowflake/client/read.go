package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	readSQL = "SELECT * FROM %s WHERE \"_cq_source_name\" = ?"
)

// https://github.com/snowflakedb/gosnowflake/issues/674
func snowflakeStrToIntArray(val string) []string {
	val = strings.TrimPrefix(val, "[\n  ")
	val = strings.TrimSuffix(val, "\n]")
	strs := strings.Split(val, ",\n  ")
	for i := range strs {
		strs[i] = strings.ReplaceAll(strs[i], "\\\"", "\"")
		strs[i] = strings.ReplaceAll(strs[i], "\\n", "\n")
	}
	return strs
}

// https://github.com/snowflakedb/gosnowflake/issues/674
func snowflakeStrToArray(val string) []string {
	val = strings.TrimPrefix(val, "[\n  \"")
	val = strings.TrimSuffix(val, "\"\n]")
	strs := strings.Split(val, "\",\n  \"")
	for i := range strs {
		strs[i] = strings.ReplaceAll(strs[i], "\\\"", "\"")
		strs[i] = strings.ReplaceAll(strs[i], "\\n", "\n")
	}
	return strs
}

func (*Client) createResultsArray(values []any, table *schema.Table) []any {
	results := make([]any, 0, len(table.Columns))
	for i, col := range table.Columns {
		switch col.Type {
		case schema.TypeBool:
			r := (*values[i].(*any)).(bool)
			results = append(results, r)
		case schema.TypeInt:
			r := (*values[i].(*any)).(string)
			results = append(results, r)
		case schema.TypeFloat:
			r := (*values[i].(*any)).(float64)
			results = append(results, r)
		case schema.TypeUUID:
			r := (*values[i].(*any)).(string)
			results = append(results, r)
		case schema.TypeString:
			r := (*values[i].(*any)).(string)
			results = append(results, r)
		case schema.TypeByteArray:
			r := (*values[i].(*any)).([]uint8)
			results = append(results, r)
		case schema.TypeStringArray:
			r := snowflakeStrToArray((*values[i].(*any)).(string))
			results = append(results, r)
		case schema.TypeTimestamp:
			r := (*values[i].(*any)).(time.Time)
			results = append(results, r)
		case schema.TypeJSON:
			r := (*values[i].(*any)).(string)
			results = append(results, r)
		case schema.TypeUUIDArray:
			r := snowflakeStrToArray((*values[i].(*any)).(string))
			results = append(results, r)
		case schema.TypeCIDR:
			r := (*values[i].(*any)).(string)
			results = append(results, r)
		case schema.TypeCIDRArray:
			r := snowflakeStrToArray((*values[i].(*any)).(string))
			results = append(results, r)
		case schema.TypeMacAddr:
			r := (*values[i].(*any)).(string)
			results = append(results, r)
		case schema.TypeMacAddrArray:
			r := snowflakeStrToArray((*values[i].(*any)).(string))
			results = append(results, r)
		case schema.TypeInet:
			r := (*values[i].(*any)).(string)
			results = append(results, r)
		case schema.TypeInetArray:
			r := snowflakeStrToArray((*values[i].(*any)).(string))
			results = append(results, r)
		case schema.TypeIntArray:
			r := snowflakeStrToIntArray((*values[i].(*any)).(string))
			results = append(results, r)
		}
	}
	return results
}

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	stmt := fmt.Sprintf(readSQL, table.Name)
	rows, err := c.db.Query(stmt, sourceName)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		values := make([]any, len(table.Columns))
		for i := range values {
			values[i] = new(any)
		}
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		goValues := c.createResultsArray(values, table)
		res <- goValues
	}
	rows.Close()
	return nil
}

package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow/go/arrow/memory"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

const (
	readSQL = "SELECT %s FROM %s WHERE \"_cq_source_name\" = ?"
)

func (c *Client) reverseTransform(f arrow.Field, bldr array.Builder, val any) error {
	if val == nil {
		bldr.AppendNull()
		return nil
	}
	switch b := bldr.(type) {
	case *array.BooleanBuilder:
		b.Append(val.(bool))
	case *array.Int8Builder:
		b.Append(val.(int8))
	case *array.Int16Builder:
		b.Append(val.(int16))
	case *array.Int32Builder:
		b.Append(val.(int32))
	case *array.Int64Builder:
		b.Append(val.(int64))
	case *array.Uint8Builder:
		b.Append(val.(uint8))
	case *array.Uint16Builder:
		b.Append(val.(uint16))
	case *array.Uint32Builder:
		b.Append(val.(uint32))
	case *array.Uint64Builder:
		b.Append(val.(uint64))
	case *array.Float32Builder:
		b.Append(val.(float32))
	case *array.Float64Builder:
		b.Append(val.(float64))
	case *array.StringBuilder:
		b.Append(val.(string))
	case *array.LargeStringBuilder:
		b.Append(val.(string))
	case *array.BinaryBuilder:
		b.Append(val.([]uint8))
	case *array.TimestampBuilder:
		b.Append(arrow.Timestamp(val.(time.Time).UnixMicro()))
	case *types.JSONBuilder:
		b.Append(val)
	case array.ListLikeBuilder:
		b.Append(true)
		valBuilder := b.ValueBuilder()
		for _, v := range val.([]any) {
			if err := c.reverseTransform(f, valBuilder, v); err != nil {
				return err
			}
		}
	default:
		v, ok := val.(string)
		if !ok {
			return fmt.Errorf("unsupported type %T with builder %T", val, bldr)
		}
		if err := bldr.AppendValueFromString(v); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) reverseTransformer(sc *arrow.Schema, values []any) (arrow.Record, error) {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	for i, f := range sc.Fields() {
		if err := c.reverseTransform(f, bldr.Field(i), values[i]); err != nil {
			return nil, err
		}
	}
	rec := bldr.NewRecord()
	bldr.Release()
	return rec, nil
}

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
		if values[i] == nil {
			results = append(results, nil)
			continue
		}
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

func (c *Client) Read(ctx context.Context, table *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	tableName := schema.TableName(table)
	colNames := make([]string, 0, len(table.Fields()))
	for _, col := range table.Fields() {
		colNames = append(colNames, `"`+col.Name+`"`)
	}
	cols := strings.Join(colNames, ", ")
	stmt := fmt.Sprintf(readSQL, cols, tableName)
	rows, err := c.db.Query(stmt, sourceName)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		values := make([]any, len(table.Fields()))
		for i := range values {
			values[i] = new(any)
		}
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to read from table %s: %w", tableName, err)
		}
		rec, err := c.reverseTransformer(table, values)
		if err != nil {
			return err
		}
		res <- rec
	}
	rows.Close()
	return nil
}

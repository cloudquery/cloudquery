package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/apache/arrow/go/arrow/memory"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/plugin-sdk/v2/schema"
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
		u, err := strconv.ParseInt(val.(string), 10, 8)
		if err != nil {
			return err
		}
		b.Append(int8(u))
	case *array.Int16Builder:
		u, err := strconv.ParseInt(val.(string), 10, 16)
		if err != nil {
			return err
		}
		b.Append(int16(u))
	case *array.Int32Builder:
		u, err := strconv.ParseInt(val.(string), 10, 32)
		if err != nil {
			return err
		}
		b.Append(int32(u))
	case *array.Int64Builder:
		u, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			return err
		}
		b.Append(u)
	case *array.Uint8Builder:
		u, err := strconv.ParseUint(val.(string), 10, 8)
		if err != nil {
			return err
		}
		b.Append(uint8(u))
	case *array.Uint16Builder:
		u, err := strconv.ParseUint(val.(string), 10, 16)
		if err != nil {
			return err
		}
		b.Append(uint16(u))
	case *array.Uint32Builder:
		u, err := strconv.ParseUint(val.(string), 10, 32)
		if err != nil {
			return err
		}
		b.Append(uint32(u))
	case *array.Uint64Builder:
		u, err := strconv.ParseUint(val.(string), 10, 64)
		if err != nil {
			return err
		}
		b.Append(u)
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
	case array.ListLikeBuilder:
		b.Append(true)
		valBuilder := b.ValueBuilder()
		s := val.(string)
		var values []string
		// nolint:gocritic,revive
		if strings.HasPrefix(s, "[\n  \"") {
			values = snowflakeStrToArray(s)
		} else if strings.HasPrefix(s, "[\n  ") {
			values = snowflakeStrToIntArray(s)
		} else {
			return fmt.Errorf("unknown array format %s", s)
		}

		for _, v := range values {
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
		if err := c.reverseTransform(f, bldr.Field(i), *values[i].(*any)); err != nil {
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
		strs[i] = strings.ReplaceAll(strs[i], "\\u0000", "\u0000")
		strs[i] = strings.ReplaceAll(strs[i], "\\n", "\n")
	}
	return strs
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

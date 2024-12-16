package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	readSQL = "SELECT %s FROM %s"
)

func (c *Client) reverseTransform(f arrow.Field, bldr array.Builder, val any) error {
	if val == nil {
		bldr.AppendNull()
		return nil
	}

	if s, ok := val.(string); ok {
		if s == "null" {
			bldr.AppendNull()
			return nil
		}
	}

	switch b := bldr.(type) {
	case *array.BooleanBuilder:
		if boolVal, ok := val.(bool); ok {
			b.Append(boolVal)
			return nil
		}
		return b.AppendValueFromString(val.(string))
	case *array.Int8Builder:
		u, err := strconv.ParseInt(val.(string), 10, 8)
		if err != nil {
			return fmt.Errorf("failed to parse int8: %w", err)
		}
		b.Append(int8(u))
	case *array.Int16Builder:
		u, err := strconv.ParseInt(val.(string), 10, 16)
		if err != nil {
			return fmt.Errorf("failed to parse int16: %w", err)
		}
		b.Append(int16(u))
	case *array.Int32Builder:
		u, err := strconv.ParseInt(val.(string), 10, 32)
		if err != nil {
			return fmt.Errorf("failed to parse int32: %w", err)
		}
		b.Append(int32(u))
	case *array.Int64Builder:
		u, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse int64: %w", err)
		}
		b.Append(u)
	case *array.Uint8Builder:
		u, err := strconv.ParseUint(val.(string), 10, 8)
		if err != nil {
			return fmt.Errorf("failed to parse uint8: %w", err)
		}
		b.Append(uint8(u))
	case *array.Uint16Builder:
		u, err := strconv.ParseUint(val.(string), 10, 16)
		if err != nil {
			return fmt.Errorf("failed to parse uint16: %w", err)
		}
		b.Append(uint16(u))
	case *array.Uint32Builder:
		u, err := strconv.ParseUint(val.(string), 10, 32)
		if err != nil {
			return fmt.Errorf("failed to parse uint32: %w", err)
		}
		b.Append(uint32(u))
	case *array.Uint64Builder:
		u, err := strconv.ParseUint(val.(string), 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse uint64: %w", err)
		}
		b.Append(u)
	case *array.Float32Builder:
		if floatVal, ok := val.(float64); ok {
			b.Append(float32(floatVal))
		} else {
			floatVal, err := strconv.ParseFloat(val.(string), 32)
			if err != nil {
				return fmt.Errorf("failed to parse float32: %w", err)
			}
			b.Append(float32(floatVal))
		}
	case *array.Float64Builder:
		if floatVal, ok := val.(float64); ok {
			b.Append(floatVal)
		} else {
			floatVal, err := strconv.ParseFloat(val.(string), 64)
			if err != nil {
				return fmt.Errorf("failed to parse float64: %w", err)
			}
			b.Append(floatVal)
		}
	case *array.StringBuilder:
		b.Append(val.(string))
	case *array.LargeStringBuilder:
		b.Append(val.(string))
	case *array.BinaryBuilder:
		b.Append(val.([]uint8))
	case *array.TimestampBuilder:
		var timeVal time.Time
		// nolint:revive
		if t, ok := val.(time.Time); ok {
			timeVal = t
		} else {
			t, err := arrow.TimestampFromString(val.(string), b.Type().(*arrow.TimestampType).Unit)
			if err != nil {
				return fmt.Errorf("failed to parse timestamp: %w", err)
			}
			b.Append(t)
			return nil
		}

		switch b.Type().(*arrow.TimestampType).Unit {
		case arrow.Second:
			b.Append(arrow.Timestamp(timeVal.UTC().Unix()))
		case arrow.Millisecond:
			b.Append(arrow.Timestamp(timeVal.UTC().UnixMilli()))
		case arrow.Microsecond:
			b.Append(arrow.Timestamp(timeVal.UTC().UnixMicro()))
		case arrow.Nanosecond:
			b.Append(arrow.Timestamp(timeVal.UTC().UnixNano()))
		default:
			return fmt.Errorf("unsupported timestamp unit %s", f.Type.(*arrow.TimestampType).Unit)
		}
	case array.ListLikeBuilder:
		b.Append(true)
		valBuilder := b.ValueBuilder()
		s := val.(string)
		var values []string
		// nolint:gocritic,revive
		if !strings.HasPrefix(s, "[\n  ") {
			return fmt.Errorf("unknown array format %s", s)
		}
		values = snowflakeStrToArray(s)

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
			return fmt.Errorf("failed to AppendValueFromString %s: %w", v, err)
		}
	}
	return nil
}

func (c *Client) reverseTransformer(table *schema.Table, values []any) (arrow.Record, error) {
	sc := table.ToArrowSchema()
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	for i, f := range sc.Fields() {
		if err := c.reverseTransform(f, bldr.Field(i), *values[i].(*any)); err != nil {
			return nil, fmt.Errorf("failed to transform field %s: %w", f.Name, err)
		}
	}
	rec := bldr.NewRecord()
	bldr.Release()
	return rec, nil
}

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	tableName := table.Name
	colNames := make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		colNames = append(colNames, `"`+strings.ToUpper(col.Name)+`"`)
	}
	cols := strings.Join(colNames, ", ")
	stmt := fmt.Sprintf(readSQL, cols, tableName)
	rows, err := c.db.QueryContext(ctx, stmt)
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

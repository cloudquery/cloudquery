package client

import (
	"context"
	"fmt"
	"net"
	"net/netip"
	"strings"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	readSQL     = "SELECT %s FROM %s WHERE _cq_source_name = $1 order by _cq_sync_time asc"
	readSQLJSON = `
	SELECT row_to_json(t)
FROM (
  SELECT *
  FROM %s
  WHERE _cq_source_name = $1
  ORDER BY _cq_sync_time ASC
) t;
	`
)

func reverseTransform(f arrow.Field, bldr array.Builder, val any) error {
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
		va, ok := val.(string)
		if !ok {
			panic(fmt.Sprintf("unsupported type %T with builder %T and column %s", val, bldr, f.Name))
		}
		b.Append(va)
	case *array.LargeStringBuilder:
		b.Append(val.(string))
	case *array.BinaryBuilder:
		b.Append(val.([]byte))
	case *array.TimestampBuilder:
		b.Append(arrow.Timestamp(val.(time.Time).UnixMicro()))
	case *types.UUIDBuilder:
		va, ok := val.([16]byte)
		if !ok {
			panic(fmt.Sprintf("unsupported type %T with builder %T", val, bldr))
		}
		u, err := uuid.FromBytes(va[:])
		if err != nil {
			return err
		}
		b.Append(u)
	case *types.JSONBuilder:
		b.Append(val)
	case *types.InetBuilder:
		if v, ok := val.(netip.Prefix); ok {
			_, ipnet, err := net.ParseCIDR(v.String())
			if err != nil {
				return err
			}
			b.Append(*ipnet)
			return nil
		}
		b.Append(val.(net.IPNet))
	case *types.MacBuilder:
		b.Append(val.(net.HardwareAddr))
	case array.ListLikeBuilder:
		b.Append(true)
		valBuilder := b.ValueBuilder()
		for _, v := range val.([]any) {
			if err := reverseTransform(f, valBuilder, v); err != nil {
				return err
			}
		}
	default:
		v, ok := val.(string)
		if !ok {
			panic(fmt.Sprintf("unsupported type %T with builder %T", val, bldr))
		}
		if err := bldr.AppendValueFromString(v); err != nil {
			return err
		}
	}
	return nil
}

func reverseTransformer(sc *arrow.Schema, values []any) (arrow.Record, error) {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	for i, f := range sc.Fields() {
		if err := reverseTransform(f, bldr.Field(i), values[i]); err != nil {
			return nil, err
		}
	}
	rec := bldr.NewRecord()
	bldr.Release()
	return rec, nil
}

func (c *Client) Read(ctx context.Context, table *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	colNames := make([]string, 0, len(table.Fields()))
	for _, col := range table.Fields() {
		colNames = append(colNames, pgx.Identifier{col.Name}.Sanitize())
	}
	cols := strings.Join(colNames, ",")
	tableName := schema.TableName(table)
	sql := fmt.Sprintf(readSQL, cols, pgx.Identifier{tableName}.Sanitize())
	rows, err := c.conn.Query(ctx, sql, sourceName)
	if err != nil {
		return err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err
		}
		rec, err := reverseTransformer(table, values)
		if err != nil {
			return err
		}
		res <- rec
	}
	rows.Close()
	return nil
}

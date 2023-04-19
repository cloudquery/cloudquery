package client

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) Read(ctx context.Context, table *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	session, closer, err := c.newSession()
	if err != nil {
		return err
	}
	defer closer()

	tableName := schema.TableName(table)
	g := gremlingo.Traversal_().WithRemote(session).
		V().
		HasLabel(tableName).
		Has(schema.CqSourceNameColumn.Name, sourceName).
		Group().By(gremlingo.T.Id).
		By(AnonT.ValueMap())

	rs, err := g.GetResultSet()
	if err != nil {
		return fmt.Errorf("GetResultSet: %w", err)
	}
	defer rs.Close()

	for row := range rs.Channel() {
		m := row.Data.(map[any]any)
		for _, rowCols := range m {
			rowData := rowCols.(map[any]any)
			rec, err := c.reverseTransformer(table, rowData)
			if err != nil {
				return err
			}
			res <- rec
		}
	}

	return rs.GetError()
}

func (c *Client) reverseTransform(f arrow.Field, bldr array.Builder, val any) error {
	if val == nil {
		bldr.AppendNull()
		return nil
	}
	switch b := bldr.(type) {
	case *array.BooleanBuilder:
		b.Append(val.(bool))
	case *array.Int8Builder:
		b.Append(int8(val.(int64)))
	case *array.Int16Builder:
		b.Append(int16(val.(int64)))
	case *array.Int32Builder:
		b.Append(int32(val.(int64)))
	case *array.Int64Builder:
		b.Append(val.(int64))
	case *array.Uint8Builder:
		b.Append(uint8(val.(int64)))
	case *array.Uint16Builder:
		b.Append(uint16(val.(int64)))
	case *array.Uint32Builder:
		b.Append(uint32(val.(int64)))
	case *array.Uint64Builder:
		b.Append(uint64(val.(int64)))
	case *array.Float32Builder:
		b.Append(float32(val.(float64)))
	case *array.Float64Builder:
		b.Append(val.(float64))
	case *array.StringBuilder:
		va, ok := val.(string)
		if !ok {
			return fmt.Errorf("unsupported type %T with builder %T and column %s", val, bldr, f.Name)
		}
		b.Append(va)
	case *array.LargeStringBuilder:
		b.Append(val.(string))
	case *array.BinaryBuilder:
		v := val.([]any)
		byteArray := make([]byte, len(v))
		for i := range v {
			byteArray[i] = v[i].(uint8)
		}
		b.Append(byteArray)
	case *array.TimestampBuilder:
		b.Append(arrow.Timestamp(val.(time.Time). /*.Round(time.Millisecond)*/ UnixMicro()))
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

func (c *Client) reverseTransformer(sc *arrow.Schema, data map[any]any) (arrow.Record, error) {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)

	for i, f := range sc.Fields() {
		if data[f.Name] == nil {
			if err := c.reverseTransform(f, bldr.Field(i), nil); err != nil {
				return nil, err
			}
			continue
		}
		data := data[f.Name].([]any)
		if l := len(data); l != 1 {
			return nil, fmt.Errorf("expected 1 value for %s, got %v", f.Name, l)
		}
		if err := c.reverseTransform(f, bldr.Field(i), data[0]); err != nil {
			return nil, err
		}
	}
	rec := bldr.NewRecord()
	bldr.Release()
	return rec, nil
}

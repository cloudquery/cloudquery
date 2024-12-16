package client

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const readCypher = "MATCH (t:%s) RETURN t"

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
		b.Append(val.([]byte))
	case *array.TimestampBuilder:
		switch b.Type().(*arrow.TimestampType).Unit {
		case arrow.Second:
			b.Append(arrow.Timestamp(val.(time.Time).Unix()))
		case arrow.Millisecond:
			b.Append(arrow.Timestamp(val.(time.Time).UnixMilli()))
		case arrow.Microsecond:
			b.Append(arrow.Timestamp(val.(time.Time).UnixMicro()))
		case arrow.Nanosecond:
			b.Append(arrow.Timestamp(val.(time.Time).UnixNano()))
		default:
			return fmt.Errorf("unsupported timestamp unit %s", f.Type.(*arrow.TimestampType).Unit)
		}
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

func (c *Client) reverseTransformer(table *schema.Table, node *neo4j.Node) (arrow.Record, error) {
	sc := table.ToArrowSchema()
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	for i, f := range sc.Fields() {
		if err := c.reverseTransform(f, bldr.Field(i), node.Props[f.Name]); err != nil {
			return nil, err
		}
	}
	rec := bldr.NewRecord()
	return rec, nil
}

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	stmt := fmt.Sprintf(readCypher, table.Name)

	session := c.Session(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		r, err := tx.Run(ctx, stmt, map[string]any{})
		if err != nil {
			return nil, err
		}
		records, err := r.Collect(ctx)
		if err != nil {
			return nil, err
		}
		for _, record := range records {
			values := record.Values
			for _, value := range values {
				node := value.(neo4j.Node)
				rec, err := c.reverseTransformer(table, &node)
				if err != nil {
					return nil, err
				}
				res <- rec
			}
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return session.Close(ctx)
}

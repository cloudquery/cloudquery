package client

import (
	"bytes"
	"context"
	"fmt"

	"github.com/goccy/go-json"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		b.Append(int8(val.(int32)))
	case *array.Int16Builder:
		b.Append(int16(val.(int32)))
	case *array.Int32Builder:
		b.Append(val.(int32))
	case *array.Int64Builder:
		b.Append(val.(int64))
	case *array.Uint8Builder:
		b.Append(uint8(val.(int32)))
	case *array.Uint16Builder:
		b.Append(uint16(val.(int32)))
	case *array.Uint32Builder:
		b.Append(uint32(val.(int64)))
	case *array.Uint64Builder:
		b.Append(uint64(val.(int64)))
	case *array.Float32Builder:
		b.Append(float32(val.(float64)))
	case *array.Float64Builder:
		b.Append(val.(float64))
	case *array.StringBuilder:
		b.Append(val.(string))
	case *array.LargeStringBuilder:
		b.Append(val.(string))
	case *array.BinaryBuilder:
		b.Append(val.(primitive.Binary).Data)
	case *array.TimestampBuilder:
		switch b.Type().(*arrow.TimestampType).Unit {
		case arrow.Second:
			b.Append(arrow.Timestamp((val).(primitive.DateTime).Time().UTC().Unix()))
		case arrow.Millisecond:
			b.Append(arrow.Timestamp((val).(primitive.DateTime).Time().UTC().UnixMilli()))
		case arrow.Microsecond:
			b.Append(arrow.Timestamp((val).(primitive.DateTime).Time().UTC().UnixMicro()))
		case arrow.Nanosecond:
			b.Append(arrow.Timestamp((val).(primitive.DateTime).Time().UTC().UnixNano()))
		default:
			return fmt.Errorf("unsupported timestamp unit %s", f.Type.(*arrow.TimestampType).Unit)
		}
	case *types.JSONBuilder:
		b.Append(val)
	case *array.StructBuilder:
		v, err := json.Marshal(val.(primitive.M))
		if err != nil {
			return err
		}
		dec := json.NewDecoder(bytes.NewReader(v))
		if err := b.UnmarshalOne(dec); err != nil {
			return err
		}
	case array.ListLikeBuilder:
		b.Append(true)
		valBuilder := b.ValueBuilder()
		for _, v := range val.(primitive.A) {
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

func (c *Client) reverseTransformer(table *schema.Table, values primitive.M) (arrow.Record, error) {
	sc := table.ToArrowSchema()
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	for i, f := range sc.Fields() {
		if err := c.reverseTransform(f, bldr.Field(i), values[sc.Field(i).Name]); err != nil {
			return nil, err
		}
	}
	rec := bldr.NewRecord()
	return rec, nil
}

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	tableName := table.Name
	cur, err := c.client.Database(c.spec.Database).Collection(tableName).Find(
		ctx,
		bson.D{})
	if err != nil {
		return fmt.Errorf("failed to read table %s: %w", tableName, err)
	}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return fmt.Errorf("failed to read from table %s: %w", tableName, err)
		}
		rec, err := c.reverseTransformer(table, result)
		if err != nil {
			return fmt.Errorf("failed to read from table %s: %w", tableName, err)
		}
		res <- rec
	}
	return nil
}

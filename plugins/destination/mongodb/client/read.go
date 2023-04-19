package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		b.Append(val.(primitive.Binary).Data)
	case *array.TimestampBuilder:
		b.Append(arrow.Timestamp((val).(primitive.DateTime).Time().UTC().UnixMicro()))
	case *types.JSONBuilder:
		b.Append(val.(primitive.M))
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

func (c *Client) reverseTransformer(sc *arrow.Schema, values primitive.M) (arrow.Record, error) {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	for i, f := range sc.Fields() {
		if err := c.reverseTransform(f, bldr.Field(i), values[sc.Field(i).Name]); err != nil {
			return nil, err
		}
	}
	rec := bldr.NewRecord()
	bldr.Release()
	return rec, nil
}

func (c *Client) Read(ctx context.Context, table *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	tableName := schema.TableName(table)
	cur, err := c.client.Database(c.pluginSpec.Database).Collection(tableName).Find(
		ctx,
		bson.M{"_cq_source_name": sourceName},
		options.Find().SetSort(bson.M{"_cq_sync_time": 1}))
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

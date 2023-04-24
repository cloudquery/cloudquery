package client

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
)

const (
	readSQL = `SELECT %s FROM "%s" WHERE _cq_source_name = $1 order by _cq_sync_time asc`
)


func (c *Client) reverseTransform(f arrow.Field, bldr array.Builder, val any) error {
	if val == nil {
		bldr.AppendNull()
		return nil
	}
	switch b := bldr.(type) {
	case *array.BooleanBuilder:
		b.Append(*val.(*bool))
	case *array.Int8Builder:
		b.Append(*val.(*int8))
	case *array.Int16Builder:
		b.Append(*val.(*int16))
	case *array.Int32Builder:
		b.Append(*val.(*int32))
	case *array.Int64Builder:
		b.Append(*val.(*int64))
	case *array.Uint8Builder:
		b.Append(*val.(*uint8))
	case *array.Uint16Builder:
		b.Append(*val.(*uint16))
	case *array.Uint32Builder:
		b.Append(*val.(*uint32))
	case *array.Uint64Builder:
		b.Append(*val.(*uint64))
	case *array.Float32Builder:
		b.Append(*val.(*float32))
	case *array.Float64Builder:
		b.Append(*val.(*float64))
	case *array.StringBuilder:
		b.Append(*val.(*string))
	case *array.LargeStringBuilder:
		b.Append(*val.(*string))
	case *array.BinaryBuilder:
		b.Append(*val.(*[]byte))
	case *array.TimestampBuilder:
		b.Append(arrow.Timestamp(val.(*time.Time).UnixMicro()))
	case *types.JSONBuilder:
		b.Append(val)
	case *types.UUIDBuilder:
		u, err := uuid.FromBytes(*val.(*[]byte))
		if err != nil {
			return err
		}
		b.Append(u)
	case array.ListLikeBuilder:
		b.Append(true)
		valBuilder := b.ValueBuilder()
		for _, v := range val.([]any) {
			if err := c.reverseTransform(f, valBuilder, v); err != nil {
				return err
			}
		}
	default:
		if err := bldr.AppendValueFromString(*val.(*string)); err != nil {
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
	return rec, nil
}

func (*Client) createResultsArray(table *arrow.Schema) []any {
	results := make([]any, 0, len(table.Fields()))
	for _, col := range table.Fields() {
		switch col.Type.ID() {
		case arrow.BOOL:
			var r bool
			results = append(results, &r)
		case arrow.BINARY, arrow.LARGE_BINARY:
			var r []byte
			results = append(results, &r)
		case arrow.INT8, arrow.INT16, arrow.INT32, arrow.INT64, arrow.UINT8, arrow.UINT16, arrow.UINT32, arrow.UINT64:
			var r int64
			results = append(results, &r)
		case arrow.FLOAT16, arrow.FLOAT32, arrow.FLOAT64:
			var r float64
			results = append(results, &r)
		case arrow.TIMESTAMP:
			var r time.Time
			results = append(results, &r)
		case arrow.LIST, arrow.LARGE_LIST:
			var r []any
			results = append(results, &r)
		case arrow.EXTENSION:
			if arrow.TypeEqual(col.Type, types.ExtensionTypes.UUID) {
				var r []byte
				results = append(results, &r)
			} else {
				var r string
				results = append(results, &r)
			}
		default:
			var r string
			results = append(results, &r)
		}
	}
	return results
}

func (c *Client) Read(ctx context.Context, sc *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	tableName := schema.TableName(sc)
	f, err := os.CreateTemp("", fmt.Sprintf("%s-*.json", tableName))
	if err != nil {
		return err
	}
	fName := f.Name()
	if err := f.Close(); err != nil {
		return err
	}

	_, err = c.db.Exec("copy " + tableName + " to '" + f.Name() + "' (timestampformat '%Y-%m-%d %H:%M:%S.%f')")
	if err != nil {
		return err
	}
	f, err = os.Open(fName)
	if err != nil {
		return err
	}

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(f)

	// Loop through the scanner, reading line by line
	for scanner.Scan() {
			line := scanner.Bytes()
			bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
			if err := bldr.UnmarshalJSON(line); err != nil {
				return err
			}
			res <- bldr.NewRecord()
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
			return err
	}

	return nil
}

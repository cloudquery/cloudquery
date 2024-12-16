package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func (c *Client) transformArr(arr arrow.Array, isCQTime bool) []any {
	dbArr := make([]any, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) || !arr.IsValid(i) {
			dbArr[i] = nil
			continue
		}
		switch a := arr.(type) {
		case *array.Boolean:
			dbArr[i] = a.Value(i)
		case *array.Int16:
			dbArr[i] = int64(a.Value(i))
		case *array.Int32:
			dbArr[i] = int64(a.Value(i))
		case *array.Int64:
			dbArr[i] = a.Value(i)
		case *array.Float32:
			dbArr[i] = float64(a.Value(i))
		case *array.Float64:
			dbArr[i] = a.Value(i)
		case *array.Binary:
			dbArr[i] = a.Value(i)
		case *array.LargeBinary:
			dbArr[i] = a.Value(i)
		case *array.String:
			dbArr[i] = stripNulls(a.Value(i))
		case *array.LargeString:
			dbArr[i] = stripNulls(a.Value(i))
		case *array.Timestamp:
			if isCQTime {
				dbArr[i] = a.Value(i).ToTime(a.DataType().(*arrow.TimestampType).Unit).UTC()
				continue
			}
			dbArr[i] = a.Value(i).ToTime(a.DataType().(*arrow.TimestampType).Unit).UTC().Format("2006-01-02 15:04:05.999999999")
		case array.ListLike:
			if !c.spec.CompleteTypes {
				dbArr[i] = stripNulls(arr.ValueStr(i))
				continue
			}
			start, end := a.ValueOffsets(i)
			nested := array.NewSlice(a.ListValues(), start, end)
			dbArr[i] = c.transformArr(nested, false)
		default:
			dbArr[i] = stripNulls(arr.ValueStr(i))
		}
	}

	return dbArr
}

func (c *Client) transformValues(r arrow.Record, cqTimeIndex int) []map[string]any {
	results := make([]map[string]any, r.NumRows())

	for i := range results {
		results[i] = make(map[string]any, r.NumCols())
	}
	sc := r.Schema()
	for i := 0; i < int(r.NumCols()); i++ {
		col := r.Column(i)
		transformed := c.transformArr(col, i == cqTimeIndex)
		for l := 0; l < col.Len(); l++ {
			results[l][sc.Field(i).Name] = transformed[l]
		}
	}
	return results
}

func stripNulls(s string) string {
	return strings.ReplaceAll(s, "\x00", "")
}

func reverseTransform(f arrow.Field, bldr array.Builder, val any) error {
	if val == nil {
		bldr.AppendNull()
		return nil
	}
	if str, ok := val.(string); ok {
		return bldr.AppendValueFromString(str)
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
		t := val.(time.Time).UTC()
		ts, err := arrow.TimestampFromString(t.Format(time.RFC3339Nano), b.Type().(*arrow.TimestampType).Unit)
		if err != nil {
			return err
		}
		b.Append(ts)
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
			return fmt.Errorf("unsupported type %T with builder %T", val, bldr)
		}
		if err := bldr.AppendValueFromString(v); err != nil {
			return err
		}
	}
	return nil
}

func reverseTransformer(sc *arrow.Schema, data map[any]any) (arrow.Record, error) {
	bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)

	for i, f := range sc.Fields() {
		if data[f.Name] == nil {
			if err := reverseTransform(f, bldr.Field(i), nil); err != nil {
				return nil, err
			}
			continue
		}
		data := data[f.Name].([]any)
		if l := len(data); l != 1 {
			return nil, fmt.Errorf("expected 1 value for %s, got %v", f.Name, l)
		}
		if err := reverseTransform(f, bldr.Field(i), data[0]); err != nil {
			return nil, err
		}
	}
	rec := bldr.NewRecord()
	return rec, nil
}

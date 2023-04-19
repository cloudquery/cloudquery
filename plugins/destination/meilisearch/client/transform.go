package client

import (
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func timestampValues(arr *array.Timestamp) []any {
	unit := arr.DataType().(*arrow.TimestampType).Unit
	res := make([]any, arr.Len())
	for i := range res {
		if arr.IsValid(i) && !arr.IsNull(i) {
			res[i] = arr.Value(i).ToTime(unit).UTC()
		}
	}
	return res
}

func getValues(arr arrow.Array) []any {
	if tsArr, ok := arr.(*array.Timestamp); ok {
		return timestampValues(tsArr)
	}

	res := make([]any, arr.Len())
	for i := range res {
		if arr.IsValid(i) && !arr.IsNull(i) {
			res[i] = arr.GetOneForMarshal(i)
		}
	}
	return res
}

func transpose(m map[string][]any, l int) []map[string]any {
	rows := make([]map[string]any, l)
	for i := range rows {
		rows[i] = make(map[string]any)
	}

	for i, row := range rows {
		for k, v := range m {
			row[k] = v[i]
		}
	}

	return rows
}

func reverseTransform(builder array.Builder, val any) error {
	if val == nil {
		builder.AppendNull()
		return nil
	}

	switch builder := builder.(type) {
	case *array.BooleanBuilder:
		builder.Append(val.(bool))
	case *array.Int8Builder:
		builder.Append(int8(val.(float64)))
	case *array.Int16Builder:
		builder.Append(int16(val.(float64)))
	case *array.Int32Builder:
		builder.Append(int32(val.(float64)))
	case *array.Int64Builder:
		builder.Append(int64(val.(float64)))
	case *array.Uint8Builder:
		builder.Append(uint8(val.(float64)))
	case *array.Uint16Builder:
		builder.Append(uint16(val.(float64)))
	case *array.Uint32Builder:
		builder.Append(uint32(val.(float64)))
	case *array.Uint64Builder:
		builder.Append(uint64(val.(float64)))
	case *array.Float32Builder:
		builder.Append(float32(val.(float64)))
	case *array.Float64Builder:
		builder.Append(val.(float64))
	case *types.JSONBuilder:
		builder.Append(val)
	case array.ListLikeBuilder:
		builder.Append(true)
		valueBuilder := builder.ValueBuilder()
		for _, v := range val.([]any) {
			if err := reverseTransform(valueBuilder, v); err != nil {
				return err
			}
		}
	default:
		v, ok := val.(string)
		if !ok {
			return fmt.Errorf("unsupported type %T with builder %T", val, builder)
		}
		if err := builder.AppendValueFromString(v); err != nil {
			return err
		}
	}
	return nil
}

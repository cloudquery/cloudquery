package client

import (
	"bytes"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/goccy/go-json"
)

func timestampValues(arr *array.Timestamp) []any {
	unit := arr.DataType().(*arrow.TimestampType).Unit
	res := make([]any, arr.Len())
	for i := range res {
		if arr.IsValid(i) {
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
		if arr.IsValid(i) {
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
		v, err := toInt64(val)
		if err != nil {
			return err
		}
		builder.Append(int8(v))
	case *array.Int16Builder:
		v, err := toInt64(val)
		if err != nil {
			return err
		}
		builder.Append(int16(v))
	case *array.Int32Builder:
		v, err := toInt64(val)
		if err != nil {
			return err
		}
		builder.Append(int32(v))
	case *array.Int64Builder:
		v, err := toInt64(val)
		if err != nil {
			return err
		}
		builder.Append(v)
	case *array.Uint8Builder:
		v, err := toUint64(val)
		if err != nil {
			return err
		}
		builder.Append(uint8(v))
	case *array.Uint16Builder:
		v, err := toUint64(val)
		if err != nil {
			return err
		}
		builder.Append(uint16(v))
	case *array.Uint32Builder:
		v, err := toUint64(val)
		if err != nil {
			return err
		}
		builder.Append(uint32(v))
	case *array.Uint64Builder:
		v, err := toUint64(val)
		if err != nil {
			return err
		}
		builder.Append(v)
	case *array.Float32Builder:
		v, err := toFloat64(val)
		if err != nil {
			return err
		}
		builder.Append(float32(v))
	case *array.Float64Builder:
		v, err := toFloat64(val)
		if err != nil {
			return err
		}
		builder.Append(v)
	case *array.BinaryBuilder:
		return appendFromString(builder, val.(string))
	case *array.StringBuilder:
		builder.Append(val.(string))
	case *array.LargeStringBuilder:
		builder.Append(val.(string))
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
		data, err := json.MarshalWithOption(val, json.DisableHTMLEscape())
		if err != nil {
			return err
		}

		dec := json.NewDecoder(bytes.NewReader(data))
		dec.UseNumber()
		if err := builder.UnmarshalOne(dec); err != nil {
			return err
		}
	}

	return nil
}

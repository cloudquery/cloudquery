package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
)

func transformArr(arr arrow.Array) []any {
	pgArr := make([]any, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) || !arr.IsValid(i) {
			pgArr[i] = nil
			continue
		}
		switch a := arr.(type) {
		case *array.Boolean:
			pgArr[i] = a.Value(i)
		case *array.Int8:
			pgArr[i] = int64(a.Value(i))
		case *array.Int16:
			pgArr[i] = int64(a.Value(i))
		case *array.Int32:
			pgArr[i] = int64(a.Value(i))
		case *array.Int64:
			pgArr[i] = a.Value(i)
		case *array.Uint8:
			pgArr[i] = uint64(a.Value(i))
		case *array.Uint16:
			pgArr[i] = uint64(a.Value(i))
		case *array.Uint32:
			pgArr[i] = uint64(a.Value(i))
		case *array.Uint64:
			pgArr[i] = a.Value(i)
		case *array.Float32:
			pgArr[i] = float64(a.Value(i))
		case *array.Float64:
			pgArr[i] = a.Value(i)
		case *array.Binary:
			pgArr[i] = a.Value(i)
		case *array.LargeBinary:
			pgArr[i] = a.Value(i)
		case *array.String:
			pgArr[i] = a.Value(i)
		case *array.LargeString:
			pgArr[i] = a.Value(i)
		case *array.Timestamp:
			pgArr[i] = a.Value(i).ToTime(a.DataType().(*arrow.TimestampType).Unit)
		case array.ListLike:
			start, end := a.ValueOffsets(i)
			nested := array.NewSlice(a.ListValues(), start, end)
			pgArr[i] = transformArr(nested)
		default:
			pgArr[i] = arr.ValueStr(i)
		}
	}

	return pgArr
}

func transformValues(r arrow.Record) []map[string]any {
	results := make([]map[string]any, r.NumRows())

	for i := range results {
		results[i] = make(map[string]any, r.NumCols())
	}
	sc := r.Schema()
	for i := 0; i < int(r.NumCols()); i++ {
		col := r.Column(i)
		transformed := transformArr(col)
		for l := 0; l < col.Len(); l++ {
			results[l][sc.Field(i).Name] = transformed[l]
		}
	}
	return results
}

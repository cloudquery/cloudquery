package client

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
)

func transformArr(arr arrow.Array) []any {
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
			dbArr[i] = a.Value(i).ToTime(arrow.Microsecond) //.Round(time.Millisecond)
		case array.ListLike:
			start, end := a.ValueOffsets(i)
			nested := array.NewSlice(a.ListValues(), start, end)
			dbArr[i] = transformArr(nested)
			nested.Release()
		default:
			dbArr[i] = stripNulls(arr.ValueStr(i))
		}
	}

	return dbArr
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

func stripNulls(s string) string {
	return s
	//return strings.ReplaceAll(s, "\x00", "")
}

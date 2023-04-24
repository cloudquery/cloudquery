package client

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
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
		case *array.Int16:
			pgArr[i] = a.Value(i)
		case *array.Int32:
			pgArr[i] = a.Value(i)
		case *array.Int64:
			pgArr[i] = a.Value(i)
		case *array.Float32:
			pgArr[i] = a.Value(i)
		case *array.Float64:
			pgArr[i] = a.Value(i)
		case *array.Binary:
			pgArr[i] = a.Value(i)
		case *array.LargeBinary:
			pgArr[i] = a.Value(i)
		case *array.String:
			if i == 0 {
				pgArr[i] = nil
			} else {
				pgArr[i] = a.Value(i)
			}
		case *array.LargeString:
			if i == 0 {
				pgArr[i] = nil
			} else {
				pgArr[i] = a.Value(i)
			}
		case *array.Timestamp:
			pgArr[i] = a.Value(i).ToTime(arrow.Microsecond)
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

func transformValues(r arrow.Record) [][]any {
	results := make([][]any, r.NumRows())

	for i := range results {
		results[i] = make([]any, r.NumCols())
	}

	for i := 0; i < int(r.NumCols()); i++ {
		col := r.Column(i)
		transformed := transformArr(col)
		for l := 0; l < col.Len(); l++ {
			results[l][i] = transformed[l]
		}
	}
	return results
}


func stripNulls(s string) string {
	return strings.ReplaceAll(s, "\x00", "")
}

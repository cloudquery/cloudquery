package client

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

func transformArr(arr arrow.Array) []any {
	pgArr := make([]any, 0, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if !arr.IsValid(i) {
			// we are using append and not pgArr[i] = nil
			// because neo4j doesn't support nil values in collection so
			// we just skip those
			// pgArr = append(pgArr, nil)
			continue
		}
		switch a := arr.(type) {
		case *array.Boolean:
			pgArr = append(pgArr, a.Value(i))
		case *array.Int8:
			pgArr = append(pgArr, int64(a.Value(i)))
		case *array.Int16:
			pgArr = append(pgArr, int64(a.Value(i)))
		case *array.Int32:
			pgArr = append(pgArr, int64(a.Value(i)))
		case *array.Int64:
			pgArr = append(pgArr, a.Value(i))
		case *array.Uint8:
			pgArr = append(pgArr, uint64(a.Value(i)))
		case *array.Uint16:
			pgArr = append(pgArr, uint64(a.Value(i)))
		case *array.Uint32:
			pgArr = append(pgArr, uint64(a.Value(i)))
		case *array.Uint64:
			// Neo4j fails if uint64 overflows int64
			pgArr = append(pgArr, int64(a.Value(i)))
		case *array.Float32:
			pgArr = append(pgArr, float64(a.Value(i)))
		case *array.Float64:
			pgArr = append(pgArr, a.Value(i))
		case *array.Binary:
			pgArr = append(pgArr, a.Value(i))
		case *array.LargeBinary:
			pgArr = append(pgArr, a.Value(i))
		case *array.String:
			pgArr = append(pgArr, a.Value(i))
		case *array.LargeString:
			pgArr = append(pgArr, a.Value(i))
		case *array.Timestamp:
			pgArr = append(pgArr, a.Value(i).ToTime(a.DataType().(*arrow.TimestampType).Unit))
		case array.ListLike:
			start, end := a.ValueOffsets(i)
			nested := array.NewSlice(a.ListValues(), start, end)
			pgArr = append(pgArr, transformArr(nested))
		default:
			pgArr = append(pgArr, arr.ValueStr(i))
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
			if l < len(transformed) {
				results[l][sc.Field(i).Name] = transformed[l]
			}
		}
	}
	return results
}

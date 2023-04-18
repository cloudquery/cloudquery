package client

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func getValue(arr arrow.Array, i int) any {
	if arr.IsNull(i) || !arr.IsValid(i) {
		return nil
	}
	switch a := arr.(type) {
	case *array.Boolean:
		return a.Value(i)
	case *array.Int8:
		return a.Value(i)
	case *array.Int16:
		return a.Value(i)
	case *array.Int32:
		return a.Value(i)
	case *array.Int64:
		return a.Value(i)
	case *array.Uint8:
		return a.Value(i)
	case *array.Uint16:
		return a.Value(i)
	case *array.Uint32:
		return a.Value(i)
	case *array.Uint64:
		return a.Value(i)
	case *array.Float32:
		return a.Value(i)
	case *array.Float64:
		return a.Value(i)
	case *array.String:
		return a.Value(i)
	case *array.Binary:
		return a.Value(i)
	case *array.FixedSizeBinary:
		return a.Value(i)
	case *array.Timestamp:
		val := arr.(*array.Timestamp).Value(i).ToTime(arrow.Microsecond)
		return val
	case *types.UUIDArray:
		bUUID, err := a.Value(i).MarshalBinary()
		if err != nil {
			panic(err)
		}
		return bUUID
	case *types.JSONArray:
		return a.ValueStr(i)
	}

	return arr.ValueStr(i)
}

func transformRecord(record arrow.Record) [][]any {
	var res [][]any
	for i := int64(0); i < record.NumRows(); i++ {
		var row []any
		for j := 0; int64(j) < record.NumCols(); j++ {
			v := getValue(record.Column(j), int(i))
			row = append(row, v)
		}
		res = append(res, row)
	}
	return res
}

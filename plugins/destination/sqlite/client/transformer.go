package client

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

func getValue(arr arrow.Array, i int) any {
	if !arr.IsValid(i) {
		return nil
	}
	switch arr.DataType().ID() {
	case arrow.BOOL:
		return arr.(*array.Boolean).Value(i)
	case arrow.INT8:
		return arr.(*array.Int8).Value(i)
	case arrow.INT16:
		return arr.(*array.Int16).Value(i)
	case arrow.INT32:
		return arr.(*array.Int32).Value(i)
	case arrow.INT64:
		return arr.(*array.Int64).Value(i)
	case arrow.UINT8:
		return arr.(*array.Uint8).Value(i)
	case arrow.UINT16:
		return arr.(*array.Uint16).Value(i)
	case arrow.UINT32:
		return arr.(*array.Uint32).Value(i)
	case arrow.UINT64:
		// SQLite fails if uint64 overflows int64
		return int64(arr.(*array.Uint64).Value(i))
	case arrow.FLOAT32:
		return arr.(*array.Float32).Value(i)
	case arrow.FLOAT64:
		return arr.(*array.Float64).Value(i)
	case arrow.STRING:
		return arr.(*array.String).Value(i)
	case arrow.BINARY:
		return arr.(*array.Binary).Value(i)
	case arrow.LARGE_BINARY:
		return arr.(*array.LargeBinary).Value(i)
	case arrow.FIXED_SIZE_BINARY:
		return arr.(*array.FixedSizeBinary).Value(i)
	default:
		return arr.ValueStr(i)
	}
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

package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func getValue(arr arrow.Array, i int) (any, error) {
	if arr.IsNull(i) || !arr.IsValid(i) {
		return nil, nil
	}
	switch a := arr.(type) {
	case *array.Boolean:
		return a.Value(i), nil
	case *array.Int8:
		return a.Value(i), nil
	case *array.Int16:
		return a.Value(i), nil
	case *array.Int32:
		return a.Value(i), nil
	case *array.Int64:
		return a.Value(i), nil
	case *array.Uint8:
		return a.Value(i), nil
	case *array.Uint16:
		return a.Value(i), nil
	case *array.Uint32:
		return a.Value(i), nil
	case *array.Uint64:
		return a.Value(i), nil
	case *array.Float32:
		return a.Value(i), nil
	case *array.Float64:
		return a.Value(i), nil
	case *array.String:
		return a.Value(i), nil
	case *array.Binary:
		return a.Value(i), nil
	case *array.FixedSizeBinary:
		return a.Value(i), nil
	case *array.Timestamp:
		val := arr.(*array.Timestamp).Value(i).ToTime(arrow.Microsecond)
		return val, nil
	case *types.UUIDArray:
		bUUID, err := a.Value(i).MarshalBinary()
		if err != nil {
			return nil, err
		}
		return bUUID, nil
	default:
		return a.ValueStr(i), nil
	}
}

func transformRecord(record arrow.Record) ([][]any, error) {
	var res [][]any
	for i := int64(0); i < record.NumRows(); i++ {
		var row []any
		for j := 0; int64(j) < record.NumCols(); j++ {
			v, err := getValue(record.Column(j), int(i))
			if err != nil {
				return nil, err
			}
			row = append(row, v)
		}
		res = append(res, row)
	}
	return res, nil
}

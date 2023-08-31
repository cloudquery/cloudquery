package client

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/apache/arrow/go/v14/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/types"
	go_ora "github.com/sijms/go-ora/v2"
)

func GetValue(arr arrow.Array, i int) (any, error) {
	if arr.IsNull(i) {
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
	case *array.Float16:
		return a.Value(i), nil
	case *array.Float32:
		return a.Value(i), nil
	case *array.Float64:
		return a.Value(i), nil
	case *array.String:
		return a.Value(i), nil
	case *array.LargeString:
		return a.Value(i), nil
	case *array.Binary:
		return a.Value(i), nil
	case *array.LargeBinary:
		return a.Value(i), nil
	case *array.FixedSizeBinary:
		return a.Value(i), nil
	case *array.Timestamp:
		toTime, err := a.DataType().(*arrow.TimestampType).GetToTimeFunc()
		if err != nil {
			return nil, err
		}
		t := toTime(a.Value(i))
		asTimeStamp := go_ora.TimeStamp(t)
		return asTimeStamp, nil
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

// used in the tests to insert the arrow.Record
func TransformRecord(record arrow.Record) ([][]any, error) {
	numRows := record.NumRows()
	res := make([][]any, numRows)
	var err error
	for i := int64(0); i < numRows; i++ {
		numCols := record.NumCols()
		row := make([]any, numCols)
		for j := 0; int64(j) < numCols; j++ {
			row[j], err = GetValue(record.Column(j), int(i))
			if err != nil {
				return nil, err
			}
		}
		res[i] = row
	}
	return res, nil
}

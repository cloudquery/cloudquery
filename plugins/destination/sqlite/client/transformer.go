package client

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
)

func getValue(arr arrow.Array, i int) (any, error) {
	if !arr.IsValid(i) {
		return nil, nil
	}
	switch arr.DataType().ID() {
	case arrow.BOOL:
		return arr.(*array.Boolean).Value(i), nil
	case arrow.INT8:
		return arr.(*array.Int8).Value(i), nil
	case arrow.INT16:
		return arr.(*array.Int16).Value(i), nil
	case arrow.INT32:
		return arr.(*array.Int32).Value(i), nil
	case arrow.INT64:
		return arr.(*array.Int64).Value(i), nil
	case arrow.UINT8:
		return arr.(*array.Uint8).Value(i), nil
	case arrow.UINT16:
		return arr.(*array.Uint16).Value(i), nil
	case arrow.UINT32:
		return arr.(*array.Uint32).Value(i), nil
	case arrow.UINT64:
		return arr.(*array.Uint64).Value(i), nil
	case arrow.FLOAT32:
		return arr.(*array.Float32).Value(i), nil
	case arrow.FLOAT64:
		return arr.(*array.Float64).Value(i), nil
	case arrow.STRING:
		return arr.(*array.String).Value(i), nil
	case arrow.BINARY:
		return arr.(*array.Binary).Value(i), nil
	case arrow.FIXED_SIZE_BINARY:
		return arr.(*array.FixedSizeBinary).Value(i), nil
	// case arrow.DATE32, arrow.DATE64,
	// 		 arrow.TIMESTAMP,
	// 		 arrow.TIME32, arrow.TIME64,
	// 		 arrow.INTERVAL_DAY_TIME,
	// 		 arrow.DECIMAL128, arrow.DECIMAL256:
	// 		 v := arr.GetOneForMarshal(i)
	// 		 // check if v is a string with go reflection
	// 		 b, err := json.Marshal(v)
	// 		 if err != nil {
	// 			 return nil, err
	// 		 }
	// 		return strings.Trim(string(b), "\""), nil
	default:
		return arr.ValueStr(i), nil

		// v := arr.GetOneForMarshal(i)
		// // check if v is a string with go reflection
		// b, err := json.Marshal(v)
		// if err != nil {
		// 	return nil, err
		// }
		// return string(b), nil
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

func stripNulls(s string) string {
	return strings.ReplaceAll(s, "\x00", "")
}

package client

import (
	"database/sql/driver"
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/marcboeker/go-duckdb"
)

func getValue(arr arrow.Array, i int) any {
	if !arr.IsValid(i) {
		if arrow.TypeEqual(arr.DataType(), types.NewUUIDType()) {
			return duckdb.UUID{}
		}

		// return typed nil
		switch arr.DataType().ID() {
		case arrow.TIMESTAMP:
			return ptrOf[time.Time]()
		case arrow.BOOL:
			return ptrOf[bool]()
		case arrow.INT8:
			return ptrOf[int8]()
		case arrow.INT16:
			return ptrOf[int16]()
		case arrow.INT32:
			return ptrOf[int32]()
		case arrow.INT64:
			return ptrOf[int64]()
		case arrow.UINT8:
			return ptrOf[uint8]()
		case arrow.UINT16:
			return ptrOf[uint16]()
		case arrow.UINT32:
			return ptrOf[uint32]()
		case arrow.UINT64:
			return ptrOf[uint64]()
		case arrow.FLOAT32:
			return ptrOf[float32]()
		case arrow.FLOAT64:
			return ptrOf[float64]()
		case arrow.STRING:
			return ptrOf[string]()
		case arrow.BINARY, arrow.LARGE_BINARY, arrow.FIXED_SIZE_BINARY:
			return ptrOf[[]byte]()
		default:
			return ptrOf[string]()
		}
	}

	if arrow.TypeEqual(arr.DataType(), types.NewUUIDType()) {
		v, _ := arr.(*types.UUIDArray).Value(i).MarshalBinary()
		return duckdb.UUID(v)
	}

	switch arr.DataType().ID() {
	case arrow.TIMESTAMP:
		ts := arr.(*array.Timestamp)
		timeUnit := ts.DataType().(*arrow.TimestampType).Unit
		return ts.Value(i).ToTime(timeUnit)
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
		return arr.(*array.Uint64).Value(i)
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

func transformRecordToGoType(record arrow.Record) [][]driver.Value {
	var res [][]driver.Value
	for i := int64(0); i < record.NumRows(); i++ {
		nc := record.NumCols()
		row := make([]driver.Value, nc)
		for j := 0; int64(j) < nc; j++ {
			row[j] = getValue(record.Column(j), int(i))
		}
		res = append(res, row)
	}
	return res
}

func ptrOf[T any]() *T {
	var val *T
	return val
}

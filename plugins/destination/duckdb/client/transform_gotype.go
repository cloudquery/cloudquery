package client

import (
	"database/sql/driver"
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/marcboeker/go-duckdb"
)

func getTypedNilValue(arr arrow.Array) any {
	switch arr.DataType().(type) {
	case *types.UUIDType:
		return nilPtrOf[duckdb.UUID]()
	case *arrow.TimestampType:
		return nilPtrOf[time.Time]()
	case *arrow.BooleanType:
		return nilPtrOf[bool]()
	case *arrow.Int8Type:
		return nilPtrOf[int8]()
	case *arrow.Int16Type:
		return nilPtrOf[int16]()
	case *arrow.Int32Type:
		return nilPtrOf[int32]()
	case *arrow.Int64Type:
		return nilPtrOf[int64]()
	case *arrow.Uint8Type:
		return nilPtrOf[uint8]()
	case *arrow.Uint16Type:
		return nilPtrOf[uint16]()
	case *arrow.Uint32Type:
		return nilPtrOf[uint32]()
	case *arrow.Uint64Type:
		return nilPtrOf[uint64]()
	case *arrow.Float32Type:
		return nilPtrOf[float32]()
	case *arrow.Float64Type:
		return nilPtrOf[float64]()
	case *arrow.StringType:
		return nilPtrOf[string]()
	case *arrow.BinaryType, *arrow.LargeBinaryType, *arrow.FixedSizeBinaryType:
		return nilPtrOf[[]byte]()
	default:
		return nilPtrOf[string]()
	}
}

func getValue(arr arrow.Array, i int, firstRow bool) any {
	if !arr.IsValid(i) {
		if !firstRow {
			return nil // Regular nil will do
		}
		return getTypedNilValue(arr)
	}

	switch arr.DataType().(type) {
	case *types.UUIDType:
		v, _ := arr.(*types.UUIDArray).Value(i).MarshalBinary()
		return duckdb.UUID(v)
	case *arrow.TimestampType:
		ts := arr.(*array.Timestamp)
		timeUnit := ts.DataType().(*arrow.TimestampType).Unit
		return ts.Value(i).ToTime(timeUnit)
	case *arrow.BooleanType:
		return arr.(*array.Boolean).Value(i)
	case *arrow.Int8Type:
		return arr.(*array.Int8).Value(i)
	case *arrow.Int16Type:
		return arr.(*array.Int16).Value(i)
	case *arrow.Int32Type:
		return arr.(*array.Int32).Value(i)
	case *arrow.Int64Type:
		return arr.(*array.Int64).Value(i)
	case *arrow.Uint8Type:
		return arr.(*array.Uint8).Value(i)
	case *arrow.Uint16Type:
		return arr.(*array.Uint16).Value(i)
	case *arrow.Uint32Type:
		return arr.(*array.Uint32).Value(i)
	case *arrow.Uint64Type:
		return arr.(*array.Uint64).Value(i)
	case *arrow.Float32Type:
		return arr.(*array.Float32).Value(i)
	case *arrow.Float64Type:
		return arr.(*array.Float64).Value(i)
	case *arrow.StringType:
		return arr.(*array.String).Value(i)
	case *arrow.BinaryType:
		return arr.(*array.Binary).Value(i)
	case *arrow.LargeBinaryType:
		return arr.(*array.LargeBinary).Value(i)
	case *arrow.FixedSizeBinaryType:
		return arr.(*array.FixedSizeBinary).Value(i)
	default:
		return arr.ValueStr(i)
	}
}

func transformRecordToGoType(record arrow.Record, firstRow bool) [][]driver.Value {
	res := make([][]driver.Value, record.NumRows())
	nc := record.NumCols()
	for i := range res {
		res[i] = make([]driver.Value, nc)
	}

	for j := 0; j < int(nc); j++ {
		col := record.Column(j)
		for i := range res {
			res[i][j] = getValue(col, i, firstRow && i == 0)
		}
	}
	return res
}

func nilPtrOf[T any]() *T {
	return nil // typed nil
}

package client

import (
	"database/sql/driver"
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/marcboeker/go-duckdb"
)

// getValue assumes the value at i is valid
func getValue(arr arrow.Array, i int) any {
	switch arr := arr.(type) {
	case *types.UUIDArray:
		v, _ := arr.Value(i).MarshalBinary()
		return duckdb.UUID(v)
	case *array.Timestamp:
		timeUnit := arr.DataType().(*arrow.TimestampType).Unit
		return arr.Value(i).ToTime(timeUnit)
	case *array.Date32:
		return arr.Value(i).ToTime()
	case *array.Date64:
		return arr.Value(i).ToTime()
	case *array.Boolean:
		return arr.Value(i)
	case *array.Int8:
		return arr.Value(i)
	case *array.Int16:
		return arr.Value(i)
	case *array.Int32:
		return arr.Value(i)
	case *array.Int64:
		return arr.Value(i)
	case *array.Uint8:
		return uint32(arr.Value(i)) // use uint32
	case *array.Uint16:
		return uint32(arr.Value(i)) // use uint32
	case *array.Uint32:
		return arr.Value(i)
	case *array.Uint64:
		return arr.Value(i)
	case *array.Float32:
		return arr.Value(i)
	case *array.Float64:
		return arr.Value(i)
	case *array.String:
		return arr.Value(i)
	case *array.Binary:
		return arr.Value(i)
	case *array.LargeBinary:
		return arr.Value(i)
	case *array.FixedSizeBinary:
		return arr.Value(i)
	case *array.Map:
		// unsupported in appender: use string
		return arr.ValueStr(i)
	case array.ListLike: // should be after *array.Map
		return arrowListToGoSlice(arr, i)
	case *array.Struct:
		// Can't create a Go struct dynamically and maps are unsupported: use string
		return arr.ValueStr(i)
	default:
		return arr.ValueStr(i)
	}
}

// Convert an Arrow list to a Go slice of driver.Value, to be able to represent nulls
func arrowListToGoSlice(arr array.ListLike, i int) any {
	from, to := arr.ValueOffsets(i)
	slice := array.NewSlice(arr.ListValues(), from, to)
	defer slice.Release()

	val := make([]driver.Value, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		if slice.IsNull(i) {
			continue
		}

		val[i] = getValue(slice, i)
	}
	return val
}

func transformRecordToGoType(record arrow.Record, arrowFields []arrow.Field, colList schema.ColumnList) ([][]driver.Value, error) {
	res := make([][]driver.Value, record.NumRows())
	tc := len(colList)
	for i := range res {
		res[i] = make([]driver.Value, tc)
	}

	for i, f := range arrowFields { // i: arrow column index
		j := colList.Index(f.Name) // look up the column index in the destination table
		if j == -1 {
			return nil, fmt.Errorf("column not found: %s", f.String()) // should never happen
		}
		arr := record.Column(i)
		for k := range res {
			if !arr.IsValid(k) {
				continue
			}

			res[k][j] = getValue(arr, k)
		}
	}
	return res, nil
}

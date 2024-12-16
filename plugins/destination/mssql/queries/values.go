package queries

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/types"
	mssql "github.com/microsoft/go-mssqldb"
)

func GetRows(table arrow.Table) ([][]any, error) {
	rows := prealloc(table.NumRows(), table.NumCols())
	var err error

	for c := 0; c < int(table.NumCols()); c++ {
		row, col := 0, table.Column(c)
		for _, chunk := range col.Data().Chunks() {
			for i := 0; i < chunk.Len(); i++ {
				rows[row][c], err = getColValue(chunk, i)
				row++
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return rows, nil
}

func prealloc(rows, cols int64) [][]any {
	result := make([][]any, rows)
	for i := range result {
		result[i] = make([]any, cols)
	}
	return result
}

func getColValue(arr arrow.Array, idx int) (any, error) {
	if arr.IsNull(idx) {
		return nil, nil
	}

	switch arr := arr.(type) {
	case *array.Boolean:
		return ptr(arr.Value(idx)), nil

	case *array.Uint8:
		return ptr(arr.Value(idx)), nil
	case *array.Uint16:
		return ptr(int32(arr.Value(idx))), nil // no special uint16 type, upscale
	case *array.Uint32:
		return ptr(int64(arr.Value(idx))), nil // no special uint32 type, upscale
	case *array.Uint64:
		return ptr(int64(arr.Value(idx))), nil // we store this as int64, although it may produce overflow and negative numbers

	case *array.Int8:
		return ptr(int16(arr.Value(idx))), nil // no special int8 type, upscale
	case *array.Int16:
		return ptr(arr.Value(idx)), nil
	case *array.Int32:
		return ptr(arr.Value(idx)), nil
	case *array.Int64:
		return ptr(arr.Value(idx)), nil

	case *array.Float32:
		return ptr(arr.Value(idx)), nil
	case *array.Float64:
		return ptr(arr.Value(idx)), nil

	case *array.LargeString:
		return ptr(arr.Value(idx)), nil
	case *array.String:
		return ptr(arr.Value(idx)), nil

	case *array.Binary:
		return ptr(arr.Value(idx)), nil
	case *array.LargeBinary:
		return ptr(arr.Value(idx)), nil
	case *array.FixedSizeBinary:
		return ptr(arr.Value(idx)), nil

	case *array.Timestamp:
		toTime, err := arr.DataType().(*arrow.TimestampType).GetToTimeFunc()
		if err != nil {
			return nil, err
		}
		return ptr(toTime(arr.Value(idx))), nil

	case *types.UUIDArray:
		val, _ := mssql.UniqueIdentifier(arr.Value(idx)).Value()
		return ptr(val.([]byte)), nil

	default:
		return ptr(arr.ValueStr(idx)), nil
	}
}

func ptr[A any](a A) *A { return &a }

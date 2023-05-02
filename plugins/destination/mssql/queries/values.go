package queries

import (
	"reflect"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/plugin-sdk/v2/types"
	mssql "github.com/microsoft/go-mssqldb"
	"golang.org/x/exp/slices"
)

func GetRows(reader array.RecordReader) ([][]any, error) {
	var rows [][]any
	for reader.Next() {
		r, err := getRecordRows(reader.Record())
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return slices.Clip(rows), nil
}

func prealloc(rows, cols int64) [][]any {
	result := make([][]any, rows)
	for i := range result {
		result[i] = make([]any, cols)
	}
	return result
}

func getRecordRows(record arrow.Record) ([][]any, error) {
	rows := prealloc(record.NumRows(), record.NumCols())
	var err error

	for row := range rows {
		for idx, col := range record.Columns() {
			rows[row][idx], err = getColValue(col, row)
			if err != nil {
				return nil, err
			}
		}
	}

	return rows, nil
}

func getColValue(arr arrow.Array, idx int) (any, error) {
	if arr.IsNull(idx) {
		return nil, nil
	}

	switch arr := arr.(type) {
	case *array.Boolean:
		return arr.Value(idx), nil

	case *array.Uint8:
		return arr.Value(idx), nil
	case *array.Uint16:
		return int16(arr.Value(idx)), nil // as we map those to the signed types for now
	case *array.Uint32:
		return int32(arr.Value(idx)), nil // as we map those to the signed types for now
	case *array.Uint64:
		return int64(arr.Value(idx)), nil // as we map those to the signed types for now

	case *array.Int8:
		return arr.Value(idx), nil
	case *array.Int16:
		return arr.Value(idx), nil
	case *array.Int32:
		return arr.Value(idx), nil
	case *array.Int64:
		return arr.Value(idx), nil

	case *array.Float32:
		return arr.Value(idx), nil
	case *array.Float64:
		return arr.Value(idx), nil

	case *array.LargeString:
		return arr.Value(idx), nil
	case *array.String:
		return arr.Value(idx), nil

	case *types.UUIDArray:
		val, _ := mssql.UniqueIdentifier(arr.Value(idx)).Value()
		return val, nil

	case *array.Binary:
		return arr.Value(idx), nil
	case *array.LargeBinary:
		return arr.Value(idx), nil
	case *array.FixedSizeBinary:
		return arr.Value(idx), nil

	case *array.Timestamp:
		toTime, err := arr.DataType().(*arrow.TimestampType).GetToTimeFunc()
		if err != nil {
			return nil, err
		}
		return toTime(arr.Value(idx)), nil

	default:
		return arr.ValueStr(idx), nil
	}
}

// columnGoType has to be in sync with SQLType
func columnGoType(_type arrow.DataType) reflect.Type {
	switch _type.(type) {
	case *arrow.BooleanType:
		return reflect.TypeOf(true)

	case *arrow.Uint8Type:
		return reflect.TypeOf(uint8(0))
	case *arrow.Int8Type, *arrow.Uint16Type, *arrow.Int16Type:
		return reflect.TypeOf(int16(0))
	case *arrow.Uint32Type, *arrow.Int32Type:
		return reflect.TypeOf(int32(0))
	case *arrow.Uint64Type, *arrow.Int64Type:
		return reflect.TypeOf(int64(0))

	case *arrow.Float32Type:
		return reflect.TypeOf(float32(0))
	case *arrow.Float64Type:
		return reflect.TypeOf(float64(0))

	case *arrow.LargeStringType, *arrow.StringType, *types.InetType, *types.MacType:
		return reflect.TypeOf("")

	case arrow.BinaryDataType, *arrow.FixedSizeBinaryType:
		return reflect.TypeOf([]byte{})

	case *types.UUIDType:
		return reflect.TypeOf([]byte{})

	case *arrow.TimestampType:
		return reflect.TypeOf(time.Time{})

	default:
		return reflect.TypeOf("")
	}
}

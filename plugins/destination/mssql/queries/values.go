package queries

import (
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
		return ptr(arr.Value(idx)), nil

	case *array.Uint8:
		return ptr(arr.Value(idx)), nil
	case *array.Uint16:
		return ptr(int16(arr.Value(idx))), nil // as we map those to the signed types for now
	case *array.Uint32:
		return ptr(int32(arr.Value(idx))), nil // as we map those to the signed types for now
	case *array.Uint64:
		return ptr(int64(arr.Value(idx))), nil // as we map those to the signed types for now

	case *array.Int8:
		return ptr(arr.Value(idx)), nil
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

	case array.ListLike:
		data, err := arr.MarshalJSON()
		if err != nil {
			return nil, err
		}
		return ptr(string(data)), nil

	default:
		return ptr(arr.ValueStr(idx)), nil
	}
}

func ptr[A any](a A) *A { return &a }

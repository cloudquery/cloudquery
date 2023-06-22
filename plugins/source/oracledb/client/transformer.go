package client

import (
	"reflect"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
	go_ora "github.com/sijms/go-ora/v2"
)

func reverseTransform(bldr array.Builder, val any) error {
	val = reflect.ValueOf(val).Elem().Interface()
	if val == nil {
		bldr.AppendNull()
		return nil
	}

	switch b := bldr.(type) {
	case *array.BooleanBuilder:
		ptr := val.(*bool)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.Int8Builder:
		ptr := val.(*int8)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.Int16Builder:
		ptr := val.(*int16)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.Int32Builder:
		ptr := val.(*int32)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.Int64Builder:
		ptr := val.(*int64)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.Uint8Builder:
		ptr := val.(*uint8)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.Uint16Builder:
		ptr := val.(*uint16)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.Uint32Builder:
		ptr := val.(*uint32)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.Uint64Builder:
		ptr := val.(*uint64)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.Float32Builder:
		ptr := val.(*float32)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.Float64Builder:
		ptr := val.(*float64)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.BinaryBuilder:
		ptr := val.(*[]byte)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.Append(*ptr)
	case *array.TimestampBuilder:
		ptr := val.(*time.Time)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		b.AppendTime(*ptr)
	case *types.UUIDBuilder:
		ptr := val.(*[]byte)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		u, err := uuid.FromBytes(*ptr)
		if err != nil {
			return err
		}
		b.Append(u)
	case *array.Decimal128Builder, *array.Decimal256Builder:
		ptr := val.(*string)
		if ptr == nil {
			b.AppendNull()
			return nil
		}
		err := b.AppendValueFromString(*ptr)
		if err != nil {
			return err
		}
	default:
		ptr := val.(*string)
		if ptr == nil {
			bldr.AppendNull()
			return nil
		}
		if err := bldr.AppendValueFromString(*ptr); err != nil {
			return err
		}
	}
	return nil
}

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

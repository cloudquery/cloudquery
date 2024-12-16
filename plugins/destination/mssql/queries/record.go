package queries

import (
	"fmt"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
)

func Record(sc *arrow.Schema, data []any) (arrow.Record, error) {
	builder := array.NewRecordBuilder(memory.DefaultAllocator, sc)

	for i, elem := range data {
		if err := buildValue(builder.Field(i), elem); err != nil {
			return nil, err
		}
	}

	return builder.NewRecord(), nil
}

func buildValue(builder array.Builder, elem any) error {
	if elem == nil {
		builder.AppendNull()
		return nil
	}

	switch builder := builder.(type) {
	case *array.BooleanBuilder:
		builder.Append(elem.(bool))

	case *array.Uint8Builder:
		builder.Append(uint8(elem.(int64)))
	case *array.Uint16Builder:
		builder.Append(uint16(elem.(int64)))
	case *array.Uint32Builder:
		builder.Append(uint32(elem.(int64)))
	case *array.Uint64Builder:
		builder.Append(uint64(elem.(int64)))

	case *array.Int8Builder:
		builder.Append(int8(elem.(int64)))
	case *array.Int16Builder:
		builder.Append(int16(elem.(int64)))
	case *array.Int32Builder:
		builder.Append(int32(elem.(int64)))
	case *array.Int64Builder:
		builder.Append(elem.(int64))

	case *array.Float32Builder:
		builder.Append(float32(elem.(float64))) // mssql scans in into float64
	case *array.Float64Builder:
		builder.Append(elem.(float64))

	case *array.BinaryBuilder:
		builder.Append(elem.([]byte))
	case *array.FixedSizeBinaryBuilder:
		builder.Append(elem.([]byte))

	case *array.StringBuilder:
		builder.Append(elem.(string))
	case *array.LargeStringBuilder:
		builder.Append(elem.(string))

	case *array.TimestampBuilder:
		value, err := timeToTimestamp(elem.(time.Time), builder.Type().(*arrow.TimestampType))
		if err != nil {
			return err
		}
		builder.Append(value)
	case *types.UUIDBuilder:
		val := mssql.UniqueIdentifier{}
		if err := val.Scan(elem); err != nil {
			return err
		}
		builder.Append(uuid.UUID(val))

	default:
		return builder.AppendValueFromString(elem.(string))
	}
	return nil
}

func timeToTimestamp(value time.Time, tsType *arrow.TimestampType) (arrow.Timestamp, error) {
	loc, err := tsType.GetZone()
	if err != nil {
		return arrow.Timestamp(0), err
	}
	if loc != nil {
		value = value.In(loc)
	}

	switch tsType.Unit {
	case arrow.Second:
		return arrow.Timestamp(value.Unix()), nil
	case arrow.Millisecond:
		return arrow.Timestamp(value.Unix()*1e3 + int64(value.Nanosecond())/1e6), nil
	case arrow.Microsecond:
		return arrow.Timestamp(value.Unix()*1e6 + int64(value.Nanosecond())/1e3), nil
	case arrow.Nanosecond:
		return arrow.Timestamp(value.UnixNano()), nil
	default:
		return arrow.Timestamp(0), fmt.Errorf("unsupported Apache Arrow time unit: %s", tsType.Unit.String())
	}
}

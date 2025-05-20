package types

import (
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/util"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func fieldFromColumn(col column.Interface) (*arrow.Field, error) {
	name := util.UnquoteID(col.Name())
	switch col := col.(type) {
	case *column.Bool:
		return &arrow.Field{Name: name, Type: new(arrow.BooleanType)}, nil

	case *column.UInt8:
		return &arrow.Field{Name: name, Type: new(arrow.Uint8Type)}, nil
	case *column.UInt16:
		return &arrow.Field{Name: name, Type: new(arrow.Uint16Type)}, nil
	case *column.UInt32:
		return &arrow.Field{Name: name, Type: new(arrow.Uint32Type)}, nil
	case *column.UInt64:
		return &arrow.Field{Name: name, Type: new(arrow.Uint64Type)}, nil
	case *column.Int8:
		return &arrow.Field{Name: name, Type: new(arrow.Int8Type)}, nil
	case *column.Int16:
		return &arrow.Field{Name: name, Type: new(arrow.Int16Type)}, nil
	case *column.Int32:
		return &arrow.Field{Name: name, Type: new(arrow.Int32Type)}, nil
	case *column.Int64:
		return &arrow.Field{Name: name, Type: new(arrow.Int64Type)}, nil

	case *column.Float32:
		return &arrow.Field{Name: name, Type: new(arrow.Float32Type)}, nil
	case *column.Float64:
		return &arrow.Field{Name: name, Type: new(arrow.Float64Type)}, nil

	case *column.String:
		return &arrow.Field{Name: name, Type: new(arrow.StringType)}, nil

	case *column.FixedString:
		// sadly, we need to parse manually here
		var byteWidth int
		if _, err := fmt.Sscanf(string(col.Type()), "FixedString(%d)", &byteWidth); err != nil {
			return nil, err
		}
		return &arrow.Field{Name: name, Type: &arrow.FixedSizeBinaryType{ByteWidth: byteWidth}}, nil

	case *column.Date, *column.Date32:
		return &arrow.Field{Name: name, Type: new(arrow.Date32Type)}, nil

	case *column.DateTime:
		return dateTimeType(name, col)

	case *column.DateTime64:
		return dateTime64Type(name, col)

	case *column.Decimal:
		return decimalType(name, col), nil

	case *column.Array:
		return arrayType(name, col)
	case *column.Nested:
		// it'll be Array(Tuple(...))
		return fieldFromColumn(col.Interface)

	case *column.Nullable:
		return nullableType(name, col)

	case *column.Map:
		return mapType(name, col)

	case *column.Tuple:
		return structType(name, col)

	case *column.UUID:
		return &arrow.Field{Name: name, Type: types.NewUUIDType()}, nil

	default:
		return &arrow.Field{Name: name, Type: new(arrow.StringType)}, nil
	}
}

func Field(name, typ string) (*arrow.Field, error) {
	col, err := column.Type(typ).Column(name, time.UTC)
	if err != nil {
		return nil, err
	}

	return fieldFromColumn(col)
}

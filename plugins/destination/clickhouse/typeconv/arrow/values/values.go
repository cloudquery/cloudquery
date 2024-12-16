package values

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func buildValue(builder array.Builder, value any) error {
	if value == nil {
		// saves checks for untyped nil
		builder.AppendNull()
		return nil
	}

	switch builder := builder.(type) {
	case *array.BooleanBuilder:
		buildPrimitive(builder, value)

	case *array.Uint8Builder:
		buildPrimitive(builder, value)
	case *array.Uint16Builder:
		buildPrimitive(builder, value)
	case *array.Uint32Builder:
		buildPrimitive(builder, value)
	case *array.Uint64Builder:
		buildPrimitive(builder, value)

	case *array.Int8Builder:
		buildPrimitive(builder, value)
	case *array.Int16Builder:
		buildPrimitive(builder, value)
	case *array.Int32Builder:
		buildPrimitive(builder, value)
	case *array.Int64Builder:
		buildPrimitive(builder, value)

	case *array.Float16Builder:
		buildFloat16(builder, value)
	case *array.Float32Builder:
		buildPrimitive(builder, value)
	case *array.Float64Builder:
		buildPrimitive(builder, value)

	case *array.StringBuilder:
		buildPrimitive(builder, value)

	case *array.BinaryBuilder: // also handles the LargeSizeBinaryBuilder
		buildBinary(builder, value)
	case *array.FixedSizeBinaryBuilder:
		buildBinary(builder, value)

	case *array.Date32Builder:
		buildDate32Values(builder, value)
	case *array.Date64Builder:
		buildDate64Values(builder, value)

	case *array.Time32Builder:
		return buildTime32Values(builder, value, builder.Type().(*arrow.Time32Type))
	case *array.Time64Builder:
		return buildTime64Values(builder, value, builder.Type().(*arrow.Time64Type))
	case *array.TimestampBuilder:
		return buildTimestampValues(builder, value)

	case *array.Decimal128Builder:
		buildDecimal128(builder, value)
	case *array.Decimal256Builder:
		buildDecimal256(builder, value)

	case *types.UUIDBuilder:
		buildPrimitive(builder, value)

	case *types.InetBuilder:
		const zero = "0.0.0.0/0"
		return buildFromStringWithZero(builder, value, zero)

	case *types.MACBuilder:
		const zero = "00:00:00:00:00:00"
		return buildFromStringWithZero(builder, value, zero)

	case *array.ExtensionBuilder, *types.JSONBuilder:
		return buildFromString(builder, value)

	case *array.StructBuilder:
		return buildStruct(builder, value)

	case *array.MapBuilder:
		// just before array.ListLikeBuilder as it also matches
		return buildMap(builder, value)

	case array.ListLikeBuilder:
		return buildList(builder, value)

	default:
		return buildFromString(builder, value)
	}

	return nil
}

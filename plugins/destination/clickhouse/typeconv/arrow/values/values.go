package values

import (
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/cloudquery/plugin-sdk/v3/types"
	"github.com/google/uuid"
)

func buildValue(builder array.Builder, value any) error {
	if value == nil {
		// saves checks for untyped nil
		builder.AppendNull()
		return nil
	}

	switch builder := builder.(type) {
	case *array.BooleanBuilder:
		buildPrimitive[bool](builder, value)

	case *array.Uint8Builder:
		buildPrimitive[uint8](builder, value)
	case *array.Uint16Builder:
		buildPrimitive[uint16](builder, value)
	case *array.Uint32Builder:
		buildPrimitive[uint32](builder, value)
	case *array.Uint64Builder:
		buildPrimitive[uint64](builder, value)

	case *array.Int8Builder:
		buildPrimitive[int8](builder, value)
	case *array.Int16Builder:
		buildPrimitive[int16](builder, value)
	case *array.Int32Builder:
		buildPrimitive[int32](builder, value)
	case *array.Int64Builder:
		buildPrimitive[int64](builder, value)

	case *array.Float16Builder:
		buildFloat16(builder, value)
	case *array.Float32Builder:
		buildPrimitive[float32](builder, value)
	case *array.Float64Builder:
		buildPrimitive[float64](builder, value)

	case *array.StringBuilder:
		buildPrimitive[string](builder, value)

	case *array.BinaryBuilder: // also handles the LargeSizeBinaryBuilder
		buildBinary(builder, value)
	case *array.FixedSizeBinaryBuilder:
		buildBinary(builder, value)

	case *array.Date32Builder:
		buildDate32Values(builder, value)
	case *array.Date64Builder:
		buildDate64Values(builder, value)

	case *array.TimestampBuilder:
		return buildTimestampValues(builder, value)

	case *array.Decimal128Builder:
		buildDecimal128(builder, value)
	case *array.Decimal256Builder:
		buildDecimal256(builder, value)

	case *types.UUIDBuilder:
		buildPrimitive[uuid.UUID](builder, value)

	case *types.JSONBuilder, *types.InetBuilder, *types.MacBuilder:
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

package values

import (
	"time"

	"github.com/apache/arrow/go/v12/arrow/array"
)

func buildValue(builder array.Builder, value any) error {
	switch builder := builder.(type) {
	case *array.BooleanBuilder:
		buildPrimitive[bool](builder, *value.(**bool))

	case *array.Uint8Builder:
		buildPrimitive[uint8](builder, *value.(**uint8))
	case *array.Uint16Builder:
		buildPrimitive[uint16](builder, *value.(**uint16))
	case *array.Uint32Builder:
		buildPrimitive[uint32](builder, *value.(**uint32))
	case *array.Uint64Builder:
		buildPrimitive[uint64](builder, *value.(**uint64))

	case *array.Int8Builder:
		buildPrimitive[int8](builder, *value.(**int8))
	case *array.Int16Builder:
		buildPrimitive[int16](builder, *value.(**int16))
	case *array.Int32Builder:
		buildPrimitive[int32](builder, *value.(**int32))
	case *array.Int64Builder:
		buildPrimitive[int64](builder, *value.(**int64))

	case *array.Float16Builder:
		buildFloat16(builder, *value.(**float32))
	case *array.Float32Builder:
		buildPrimitive[float32](builder, *value.(**float32))
	case *array.Float64Builder:
		buildPrimitive[float64](builder, *value.(**float64))

	case *array.StringBuilder:
		buildPrimitive[string](builder, *value.(**string))

	case *array.BinaryBuilder: // also handles the LargeSizeBinaryBuilder
		buildBinary(builder, *value.(**string))
	case *array.FixedSizeBinaryBuilder:
		buildBinary(builder, *value.(**string))

	case *array.Date32Builder:
		buildDate32Values(builder, *value.(**time.Time))
	case *array.Date64Builder:
		buildDate64Values(builder, *value.(**time.Time))

	case *array.TimestampBuilder:

	case *array.MapBuilder:
		// just before other list-like builders, as this one is special
	case array.ListLikeBuilder:
		return buildList(builder, value)
	}
	return nil
}

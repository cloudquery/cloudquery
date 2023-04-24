package values

import (
	"github.com/apache/arrow/go/v12/arrow/array"
)

func buildValue(builder array.Builder, value any) error {
	switch builder := builder.(type) {
	case *array.BooleanBuilder:
		buildPrimitiveValues[bool](builder, *value.(**bool))

	case *array.Uint8Builder:
		buildPrimitiveValues[uint8](builder, *value.(**uint8))
	case *array.Uint16Builder:
		buildPrimitiveValues[uint16](builder, *value.(**uint16))
	case *array.Uint32Builder:
		buildPrimitiveValues[uint32](builder, *value.(**uint32))
	case *array.Uint64Builder:
		buildPrimitiveValues[uint64](builder, *value.(**uint64))

	case *array.Int8Builder:
		buildPrimitiveValues[int8](builder, *value.(**int8))
	case *array.Int16Builder:
		buildPrimitiveValues[int16](builder, *value.(**int16))
	case *array.Int32Builder:
		buildPrimitiveValues[int32](builder, *value.(**int32))
	case *array.Int64Builder:
		buildPrimitiveValues[int64](builder, *value.(**int64))

	case *array.Float16Builder:
		buildFloat16Values(builder, *value.(**float32))
	case *array.Float32Builder:
		buildPrimitiveValues[float32](builder, *value.(**float32))
	case *array.Float64Builder:
		buildPrimitiveValues[float64](builder, *value.(**float64))

	case *array.MapBuilder:
		// just before other list-like builders, as this one is special
	case array.ListLikeBuilder:
		return buildListValues(builder, value)
	}
	return nil
}

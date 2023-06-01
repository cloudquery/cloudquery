package client

import (
	"reflect"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case datadog.NullableInt, datadog.NullableInt32, datadog.NullableInt64:
		return arrow.PrimitiveTypes.Int64, nil
	case datadog.NullableBool:
		return arrow.FixedWidthTypes.Boolean, nil
	case datadog.NullableFloat32:
		return arrow.PrimitiveTypes.Float32, nil
	case datadog.NullableFloat64:
		return arrow.PrimitiveTypes.Float64, nil
	case datadog.NullableString:
		return arrow.BinaryTypes.String, nil
	case datadog.NullableTime:
		return arrow.FixedWidthTypes.Timestamp_us, nil
	default:
		return nil, nil
	}
}

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

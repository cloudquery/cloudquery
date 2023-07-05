package client

import (
	"context"
	"reflect"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/thoas/go-funk"
)

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case datadog.NullableInt, datadog.NullableInt64:
		return arrow.PrimitiveTypes.Int64, nil
	case datadog.NullableInt32:
		return arrow.PrimitiveTypes.Int32, nil
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

type nullable interface {
	IsSet() bool
}

func resolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case
		datadog.NullableBool,
		datadog.NullableFloat32,
		datadog.NullableFloat64,
		datadog.NullableInt32,
		datadog.NullableInt64,
		datadog.NullableInt,
		datadog.NullableString,
		datadog.NullableTime:

		return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
			val := funk.Get(r.Item, path, funk.WithAllowZero())
			if !val.(nullable).IsSet() {
				return r.Set(c.Name, nil)
			}
			getFunc, ok := reflect.TypeOf(val).MethodByName("Get")
			if !ok {
				panic("no Get() on nullable type " + field.Type.String())
			}

			ret := getFunc.Func.Call([]reflect.Value{reflect.ValueOf(val)})
			if len(ret) != 1 {
				panic("expected 1 return value from Get()")
			}

			return r.Set(c.Name, ret[0].Interface())
		}

	default:
		return nil
	}
}

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
	transformers.WithResolverTransformer(resolverTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

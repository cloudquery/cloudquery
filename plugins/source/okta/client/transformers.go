package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
	transformers.WithResolverTransformer(resolverTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if field.Type == reflect.TypeOf(okta.NullableTime{}) {
		return schema.TypeTimestamp, nil
	}

	return schema.TypeInvalid, nil
}

func resolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	if field.Type == reflect.TypeOf(okta.NullableTime{}) {
		return ResolveNullableTime(path)
	}

	return transformers.DefaultResolverTransformer(field, path)
}

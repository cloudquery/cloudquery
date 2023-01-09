package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func TransformerOptions() []transformers.StructTransformerOption {
	return []transformers.StructTransformerOption{
		transformers.WithTypeTransformer(typeTransformer),
		transformers.WithResolverTransformer(resolverTransformer),
	}
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if field.Type == reflect.TypeOf(okta.NullableTime{}) {
		return schema.TypeTimestamp, nil
	}
	return codegen.DefaultTypeTransformer(field)
}

func resolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	if field.Type == reflect.TypeOf(okta.NullableTime{}) {
		return ResolveNullableTime(path)
	}

	return transformers.DefaultResolverTransformer(field, path)
}

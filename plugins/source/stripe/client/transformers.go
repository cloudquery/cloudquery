package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SharedTransformers(others ...transformers.StructTransformerOption) []transformers.StructTransformerOption {
	return append([]transformers.StructTransformerOption{
		transformers.WithTypeTransformer(typeTransformer),
	}, others...)
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if field.Name == "Created" && field.Type == reflect.TypeOf(int64(0)) {
		return schema.TypeTimestamp, nil
	}

	return schema.TypeInvalid, nil
}

package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}
func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if field.Name == "Created" && field.Type == reflect.TypeOf(int64(0)) {
		return schema.TypeTimestamp, nil
	}

	return schema.TypeInvalid, nil
}

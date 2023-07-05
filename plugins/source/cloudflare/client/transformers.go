package client

import (
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case []any:
		return types.ExtensionTypes.JSON, nil
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

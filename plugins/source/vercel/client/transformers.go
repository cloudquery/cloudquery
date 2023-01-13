package client

import (
	"reflect"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SharedTransformers() []transformers.StructTransformerOption {
	return []transformers.StructTransformerOption{
		transformers.WithTypeTransformer(typeTransformer),
	}
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	switch field.Type {
	case reflect.TypeOf(vercel.MilliTime{}), reflect.TypeOf(&vercel.MilliTime{}):
		return schema.TypeTimestamp, nil
	default:
		return schema.TypeInvalid, nil
	}
}

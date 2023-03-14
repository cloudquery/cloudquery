package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

var options = []transformers.StructTransformerOption{
	transformers.WithUnwrapAllEmbeddedStructs(),
	transformers.WithUnwrapStructFields("Spec", "Status"),
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	timestamp := github.Timestamp{}
	switch field.Type {
	case reflect.TypeOf(timestamp), reflect.TypeOf(&timestamp):
		return schema.TypeTimestamp, nil
	default:
		return schema.TypeInvalid, nil
	}
}

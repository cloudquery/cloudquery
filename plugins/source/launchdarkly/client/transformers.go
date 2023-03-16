package client

import (
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

var options = []transformers.StructTransformerOption{
	transformers.WithNameTransformer(nameTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

func nameTransformer(f reflect.StructField) (string, error) {
	n, err := transformers.DefaultNameTransformer(f)
	if err != nil {
		return n, err
	}

	return strings.TrimPrefix(n, "_"), nil
}

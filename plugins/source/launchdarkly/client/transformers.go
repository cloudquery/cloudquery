package client

import (
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/transformers"
)

func SharedTransformers(others ...transformers.StructTransformerOption) []transformers.StructTransformerOption {
	return append([]transformers.StructTransformerOption{
		transformers.WithNameTransformer(nameTransformer),
	}, others...)
}

func nameTransformer(f reflect.StructField) (string, error) {
	n, err := transformers.DefaultNameTransformer(f)
	if err != nil {
		return n, err
	}

	return strings.TrimPrefix(n, "_"), nil
}

package client

import (
	"reflect"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SharedTransformers(others ...transformers.StructTransformerOption) []transformers.StructTransformerOption {
	return append([]transformers.StructTransformerOption{
		transformers.WithTypeTransformer(typeTransformer),
	}, others...)
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if field.Type == reflect.TypeOf(mixpanel.Time{}) {
		return schema.TypeTimestamp, nil
	}

	return schema.TypeInvalid, nil
}

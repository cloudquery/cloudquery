package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func SharedTransformers(others ...transformers.StructTransformerOption) []transformers.StructTransformerOption {
	return append([]transformers.StructTransformerOption{
		transformers.WithTypeTransformer(typeTransformer),
	}, others...)
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if field.Type == reflect.TypeOf(tailscale.Time{}) {
		return schema.TypeTimestamp, nil
	}

	return schema.TypeInvalid, nil
}

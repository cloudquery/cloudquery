package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if field.Type == reflect.TypeOf(tailscale.Time{}) {
		return schema.TypeTimestamp, nil
	}

	return schema.TypeInvalid, nil
}

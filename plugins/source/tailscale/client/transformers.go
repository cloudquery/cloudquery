package client

import (
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	if field.Type == reflect.TypeOf(tailscale.Time{}) {
		return arrow.FixedWidthTypes.Timestamp_us, nil
	}

	return nil, nil
}

package client

import (
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	switch field.Type {
	case reflect.TypeOf(mixpanel.Time{}):
		return arrow.FixedWidthTypes.Timestamp_us, nil
	case reflect.TypeOf([]any{}):
		return types.ExtensionTypes.JSON, nil
	default:
		return nil, nil
	}
}

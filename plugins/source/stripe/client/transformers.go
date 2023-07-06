package client

import (
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}
func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	if field.Name == "Created" && field.Type == reflect.TypeOf(int64(0)) {
		return arrow.FixedWidthTypes.Timestamp_us, nil
	}

	return nil, nil
}

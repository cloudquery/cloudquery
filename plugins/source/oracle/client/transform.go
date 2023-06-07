package client

import (
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/oracle/oci-go-sdk/v65/common"
)

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	fieldType := field.Type

	for fieldType.Kind() == reflect.Pointer {
		fieldType = fieldType.Elem()
	}

	if fieldType == reflect.TypeOf(common.SDKTime{}) {
		return arrow.FixedWidthTypes.Timestamp_us, nil
	}

	return nil, nil
}

var options = []transformers.StructTransformerOption{
	transformers.WithPrimaryKeys("Id"),
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

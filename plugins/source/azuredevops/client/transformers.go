package client

import (
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
)

func isUUID(fieldType reflect.Type) bool {
	fieldKind := fieldType.Kind()

	if fieldKind == reflect.Ptr {
		return isUUID(fieldType.Elem())
	}

	return fieldType == reflect.TypeOf(uuid.UUID{})
}

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	if isUUID(field.Type) {
		return types.ExtensionTypes.UUID, nil
	}
	return nil, nil
}

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

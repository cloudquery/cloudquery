package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/google/uuid"
)

func isUUID(fieldType reflect.Type) bool {
	fieldKind := fieldType.Kind()

	if fieldKind == reflect.Ptr {
		return isUUID(fieldType.Elem())
	}

	return fieldType == reflect.TypeOf(uuid.UUID{})
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if isUUID(field.Type) {
		return schema.TypeUUID, nil
	}
	return schema.TypeInvalid, nil
}

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

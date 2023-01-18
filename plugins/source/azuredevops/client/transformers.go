package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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

func Options() []transformers.StructTransformerOption {
	options := []transformers.StructTransformerOption{
		transformers.WithTypeTransformer(typeTransformer),
	}

	return options
}

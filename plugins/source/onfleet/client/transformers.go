package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"golang.org/x/exp/slices"
)

func OnfleetTypeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if isTimestampField(field) {
		return schema.TypeTimestamp, nil
	}

	return transformers.DefaultTypeTransformer(field)
}

func isTimestampField(field reflect.StructField) bool {
	timestampFieldNames := []string{
		"TimeCreated",
		"TimeLastModified",
		"TimeLastSeen",
		"CompletedAfter",
		"CompleteAfter",
		"EstimatedCompletionTime",
		"EstimatedArrivalTime",
	}

	fieldType := field.Type
	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}

	return fieldType.Kind() == reflect.Int64 &&
		slices.Contains(timestampFieldNames, field.Name)
}

func OnfleetResolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	if isTimestampField(field) {
		return ResolveTimestampField(path)
	}

	return nil
}

func OnfleetSharedTransformers() []transformers.StructTransformerOption {
	return []transformers.StructTransformerOption{
		transformers.WithPrimaryKeys("Id"),
		transformers.WithResolverTransformer(OnfleetResolverTransformer),
		transformers.WithTypeTransformer(OnfleetTypeTransformer),
	}
}

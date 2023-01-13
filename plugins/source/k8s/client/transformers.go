package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SharedTransformersWithMoreSkipFields(skips []string) []transformers.StructTransformerOption {
	return []transformers.StructTransformerOption{
		transformers.WithUnwrapAllEmbeddedStructs(),
		transformers.WithSkipFields(
			append([]string{
				"GenerateName",
				"SelfLink",
				"CreationTimestamp",
				"DeletionTimestamp",
				"ZZZ_DeprecatedClusterName",
				"ManagedFields",
			}, skips...)...,
		),
		transformers.WithUnwrapStructFields("Spec", "Status"),
		transformers.WithTypeTransformer(typeTransformer),
	}
}

func SharedTransformers() []transformers.StructTransformerOption {
	return SharedTransformersWithMoreSkipFields(nil)
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if isK8sTimeStruct(field.Type) {
		return schema.TypeTimestamp, nil
	}

	return transformers.DefaultTypeTransformer(field)
}

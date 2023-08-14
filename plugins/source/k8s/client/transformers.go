package client

import (
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

var (
	options = []transformers.StructTransformerOption{
		transformers.WithUnwrapAllEmbeddedStructs(),
		transformers.WithSkipFields(skipFields...),
		transformers.WithUnwrapStructFields("Spec", "Status"),
		transformers.WithTypeTransformer(typeTransformer),
	}
	skipFields = []string{
		"GenerateName",
		"SelfLink",
		"CreationTimestamp",
		"DeletionTimestamp",
		"ZZZ_DeprecatedClusterName",
		"ManagedFields",
	}
)

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

func WithMoreSkipFields(extra ...string) transformers.StructTransformerOption {
	return transformers.WithSkipFields(append(skipFields, extra...)...)
}

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	if isK8sTimeStruct(field.Type) {
		return arrow.FixedWidthTypes.Timestamp_us, nil
	}

	return transformers.DefaultTypeTransformer(field)
}

package client

import (
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SharedTransformers() []transformers.StructTransformerOption {
	return []transformers.StructTransformerOption{
		transformers.WithUnwrapAllEmbeddedStructs(),
		transformers.WithSkipFields(
			"GenerateName",
			"SelfLink",
			"CreationTimestamp",
			"DeletionTimestamp",
			"ZZZ_DeprecatedClusterName",
			"ManagedFields",
		),
	}
}

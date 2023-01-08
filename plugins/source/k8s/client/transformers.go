package client

import "github.com/cloudquery/plugin-sdk/transformers"

func SharedTransformers() []transformers.StructTransformerOption {
	return []transformers.StructTransformerOption{
		transformers.WithUnwrapAllEmbeddedStructs(),
		transformers.WithSkipFields("ManagedFields"),
	}
}
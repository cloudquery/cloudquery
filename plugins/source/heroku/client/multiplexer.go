package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func NoMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	return []schema.ClientMeta{meta.(*Client)}
}

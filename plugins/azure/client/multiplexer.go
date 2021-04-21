package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func SubscriptionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var c = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, subId := range client.subscriptions {
		c = append(c, client.withSubscription(subId))
	}
	return c
}

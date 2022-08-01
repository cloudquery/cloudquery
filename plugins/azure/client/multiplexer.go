package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

// SingleSubscriptionMultiplex can be used in cases where any single subscription would do. This can reduce the number
// of API calls.
func SingleSubscriptionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	if len(client.subscriptions) == 0 {
		return []schema.ClientMeta{}
	}
	return []schema.ClientMeta{
		client.withSubscription(client.subscriptions[0]),
	}
}

func SubscriptionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	var c = make([]schema.ClientMeta, len(client.subscriptions))
	for i, subId := range client.subscriptions {
		c[i] = client.withSubscription(subId)
	}
	return c
}

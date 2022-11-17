package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/exp/maps"
)

// SingleSubscriptionMultiplex can be used in cases where any single subscription would do. This can reduce the number
// of API calls.
func SingleSubscriptionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	subscriptions := maps.Keys(client.services)
	if len(subscriptions) == 0 {
		return []schema.ClientMeta{}
	}
	return []schema.ClientMeta{
		client.withSubscription(subscriptions[0]),
	}
}
func SubscriptionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	res := make([]schema.ClientMeta, 0, len(client.services))
	for subscriptionID := range client.services {
		res = append(res, client.withSubscription(subscriptionID))
	}
	return res
}

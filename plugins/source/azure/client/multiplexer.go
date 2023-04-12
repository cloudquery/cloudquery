package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func SubscriptionMultiplexRegisteredNamespace(table, namespace string) func(schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)
		var c = make([]schema.ClientMeta, 0)
		for _, subId := range client.subscriptions {
			if _, ok := client.registeredNamespaces[subId][namespace]; ok {
				c = append(c, client.withSubscription(subId))
			} else {
				client.Logger().Info().
					Str("subscription_id", subId).
					Str("namespace", namespace).
					Str("table", table).
					Msg("Skipping namespace, not registered for subscription")
			}
		}
		return c
	}
}

func SubscriptionResourceGroupMultiplexRegisteredNamespace(table string, namespace string) func(schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)
		var c = make([]schema.ClientMeta, 0)
		for _, subId := range client.subscriptions {
			if _, ok := client.registeredNamespaces[subId][namespace]; ok {
				for _, rg := range client.ResourceGroups[subId] {
					c = append(c, client.withSubscription(subId).withResourceGroup(*rg.Name))
				}
			} else {
				client.Logger().Info().
					Str("subscription_id", subId).
					Str("namespace", namespace).
					Str("table", table).
					Msg("Skipping namespace, not registered for subscription")
			}
		}
		return c
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

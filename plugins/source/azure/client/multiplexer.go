package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

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

func SubscriptionMultiplexRegisteredNamespace(namespace string) func(schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)
		var c = make([]schema.ClientMeta, 0)
		for _, subId := range client.subscriptions {
			if client.debugMode {
				c = append(c, client.withSubscription(subId))
				continue
			}
			if _, ok := client.registeredNamespace[subId][namespace]; ok {
				c = append(c, client.withSubscription(subId))
			}
		}
		return c
	}
}

func SubscriptionResourceGroupMultiplexRegisteredNamespace(namespace string) func(schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)
		var c = make([]schema.ClientMeta, 0)
		for _, subId := range client.subscriptions {
			for _, rg := range client.resourceGroups[subId] {
				if client.debugMode {
					c = append(c, client.withSubscription(subId).withResourceGroup(*rg.Name))
					continue
				}
				if _, ok := client.registeredNamespace[subId][namespace]; ok {
					c = append(c, client.withSubscription(subId).withResourceGroup(*rg.Name))
				}
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

func SubscriptionResourceGroupMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	var c = make([]schema.ClientMeta, 0)
	for _, subId := range client.subscriptions {
		for _, rg := range client.resourceGroups[subId] {
			c = append(c, client.withSubscription(subId).withResourceGroup(*rg.Name))
		}
	}
	return c
}

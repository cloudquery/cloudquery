package client

import (
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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

func BillingAccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	var c = make([]schema.ClientMeta, len(client.BillingAccounts))
	for i := range client.BillingAccounts {
		c[i] = client.withBillingAccount(client.BillingAccounts[i])
	}
	return c
}
func isModernAccount(account *armbilling.Account) bool {
	return strings.Contains(*account.Name, ":")
}

func LegacyBillingAccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	var c = make([]schema.ClientMeta, 0)
	for i := range client.BillingAccounts {
		if !isModernAccount(client.BillingAccounts[i]) {
			c = append(c, client.withBillingAccount(client.BillingAccounts[i]))
		}
	}
	return c
}

func ModernBillingAccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	var c = make([]schema.ClientMeta, 0)
	for i := range client.BillingAccounts {
		if isModernAccount(client.BillingAccounts[i]) {
			c = append(c, client.withBillingAccount(client.BillingAccounts[i]))
		}
	}
	return c
}

func BillingAccountProfileMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	var c = make([]schema.ClientMeta, 0)
	for i := range client.BillingAccounts {
		if client.BillingAccounts[i].Properties.BillingProfiles == nil {
			continue
		}
		profiles := client.BillingAccounts[i].Properties.BillingProfiles.Value
		for j := range profiles {
			c = append(c, client.withBillingAccount(client.BillingAccounts[i]).withBillingProfile(profiles[j]))
		}
	}
	return c
}

func SubscriptionBillingPeriodMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	var c = make([]schema.ClientMeta, 0)
	for _, subID := range client.subscriptions {
		periodsForSubscription := client.BillingPeriods[subID]
		for _, period := range periodsForSubscription {
			c = append(c, client.withSubscription(subID).withBillingPeriod(period))
		}
	}
	return c
}

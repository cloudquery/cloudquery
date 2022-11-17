package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/subscription"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Subscription() []*resource.Resource {
	return []*resource.Resource{
		{
			Name:      "azure_subscriptions",
			Struct:    new(armsubscription.Subscription),
			Multiplex: "client.SingleSubscriptionMultiplex",
			Resolver: &resource.FuncParams{
				Func: subscription.SubscriptionsClient.NewListPager,
			},
		},
		{
			Struct: new(armsubscription.Location),
			Resolver: &resource.FuncParams{
				Func:   subscription.SubscriptionsClient.NewListLocationsPager,
				Params: []string{"c.SubscriptionID"},
			},
		},
		{
			SubService: "tenants",
			Struct:     new(armsubscription.TenantIDDescription),
			Resolver: &resource.FuncParams{
				Func: subscription.TenantsClient.NewListPager,
			},
		},
	}
}

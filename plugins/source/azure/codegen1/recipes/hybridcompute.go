// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"

func Armhybridcompute() []*Table {
	tables := []*Table{
		{
			NewFunc:        armhybridcompute.NewPrivateLinkScopesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.HybridCompute/privateLinkScopes",
			Namespace:      "Microsoft.HybridCompute",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_HybridCompute)`,
			Pager:          `NewListPager`,
			ResponseStruct: "PrivateLinkScopesClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armhybridcompute())
}

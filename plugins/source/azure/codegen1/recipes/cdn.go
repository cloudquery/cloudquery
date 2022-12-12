// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"

func Armcdn() []*Table {
	tables := []*Table{
		{
			NewFunc:        armcdn.NewEdgeNodesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn",
			URL:            "/providers/Microsoft.Cdn/edgenodes",
			Namespace:      "Microsoft.Cdn",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Cdn)`,
			Pager:          `NewListPager`,
			ResponseStruct: "EdgeNodesClientListResponse",
		},
		{
			NewFunc:        armcdn.NewManagedRuleSetsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/cdnWebApplicationFirewallManagedRuleSets",
			Namespace:      "Microsoft.Cdn",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Cdn)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ManagedRuleSetsClientListResponse",
		},
		{
			NewFunc:        armcdn.NewResourceUsageClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/checkResourceUsage",
			Namespace:      "Microsoft.Cdn",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Cdn)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ResourceUsageClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcdn())
}

// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric"

func Armengagementfabric() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armengagementfabric.NewAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.EngagementFabric/Accounts",
		},
		{
			NewFunc: armengagementfabric.NewChannelsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric",
			URL: "",
		},
		{
			NewFunc: armengagementfabric.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric",
			URL: "",
		},
		{
			NewFunc: armengagementfabric.NewSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.EngagementFabric/skus",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armengagementfabric())
}
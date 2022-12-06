// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"

func Armcdn() []*Table {
	tables := []*Table{
		{
			NewFunc: armcdn.NewResourceUsageClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/checkResourceUsage",
		},
		{
			NewFunc: armcdn.NewEdgeNodesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn",
			URL: "/providers/Microsoft.Cdn/edgenodes",
		},
		{
			NewFunc: armcdn.NewProfilesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/profiles",
		},
		{
			NewFunc: armcdn.NewPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/cdnWebApplicationFirewallPolicies",
		},
		{
			NewFunc: armcdn.NewManagedRuleSetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/cdnWebApplicationFirewallManagedRuleSets",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcdn())
}
// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"

func Armfrontdoor() []*Table {
	tables := []*Table{
		{
			NewFunc:   armfrontdoor.NewPoliciesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor",
			URL:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies",
			Namespace: "Microsoft.Network",
			Multiplex: `client.SubscriptionResourceGroupMultiplexRegisteredNamespace("Microsoft.Network")`,
		},
		{
			NewFunc:   armfrontdoor.NewFrontDoorsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Network/frontDoors",
			Namespace: "Microsoft.Network",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Network")`,
		},
		{
			NewFunc:   armfrontdoor.NewManagedRuleSetsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallManagedRuleSets",
			Namespace: "Microsoft.Network",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Network")`,
		},
		{
			NewFunc:   armfrontdoor.NewNetworkExperimentProfilesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Network/NetworkExperimentProfiles",
			Namespace: "Microsoft.Network",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Network")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armfrontdoor())
}

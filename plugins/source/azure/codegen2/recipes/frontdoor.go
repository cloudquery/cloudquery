// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"

func Armfrontdoor() []Table {
	tables := []Table{
		{
			Service:        "armfrontdoor",
			Name:           "front_doors",
			Struct:         &armfrontdoor.FrontDoor{},
			ResponseStruct: &armfrontdoor.FrontDoorsClientListResponse{},
			Client:         &armfrontdoor.FrontDoorsClient{},
			ListFunc:       (&armfrontdoor.FrontDoorsClient{}).NewListPager,
			NewFunc:        armfrontdoor.NewFrontDoorsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Network/frontDoors",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Network)`,
		},
		{
			Service:        "armfrontdoor",
			Name:           "managed_rule_sets",
			Struct:         &armfrontdoor.ManagedRuleSetDefinition{},
			ResponseStruct: &armfrontdoor.ManagedRuleSetsClientListResponse{},
			Client:         &armfrontdoor.ManagedRuleSetsClient{},
			ListFunc:       (&armfrontdoor.ManagedRuleSetsClient{}).NewListPager,
			NewFunc:        armfrontdoor.NewManagedRuleSetsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallManagedRuleSets",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Network)`,
		},
		{
			Service:        "armfrontdoor",
			Name:           "network_experiment_profiles",
			Struct:         &armfrontdoor.Profile{},
			ResponseStruct: &armfrontdoor.NetworkExperimentProfilesClientListResponse{},
			Client:         &armfrontdoor.NetworkExperimentProfilesClient{},
			ListFunc:       (&armfrontdoor.NetworkExperimentProfilesClient{}).NewListPager,
			NewFunc:        armfrontdoor.NewNetworkExperimentProfilesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Network/NetworkExperimentProfiles",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Network)`,
		},
		{
			Service:        "armfrontdoor",
			Name:           "policies",
			Struct:         &armfrontdoor.WebApplicationFirewallPolicy{},
			ResponseStruct: &armfrontdoor.PoliciesClientListResponse{},
			Client:         &armfrontdoor.PoliciesClient{},
			ListFunc:       (&armfrontdoor.PoliciesClient{}).NewListPager,
			NewFunc:        armfrontdoor.NewPoliciesClient,
			URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies",
			Multiplex:      `client.SubscriptionResourceGroupMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Network)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armfrontdoor()...)
}

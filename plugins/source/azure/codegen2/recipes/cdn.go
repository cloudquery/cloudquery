// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"

func Armcdn() []Table {
	tables := []Table{
		{
			Name:           "managed_rule_sets",
			Struct:         &armcdn.ManagedRuleSetDefinition{},
			ResponseStruct: &armcdn.ManagedRuleSetsClientListResponse{},
			Client:         &armcdn.ManagedRuleSetsClient{},
			ListFunc:       (&armcdn.ManagedRuleSetsClient{}).NewListPager,
			NewFunc:        armcdn.NewManagedRuleSetsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/cdnWebApplicationFirewallManagedRuleSets",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Cdn")`,
		},
		{
			Name:           "policies",
			Struct:         &armcdn.WebApplicationFirewallPolicy{},
			ResponseStruct: &armcdn.PoliciesClientListResponse{},
			Client:         &armcdn.PoliciesClient{},
			ListFunc:       (&armcdn.PoliciesClient{}).NewListPager,
			NewFunc:        armcdn.NewPoliciesClient,
			URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/cdnWebApplicationFirewallPolicies",
			Multiplex:      `client.SubscriptionResourceGroupMultiplexRegisteredNamespace("Microsoft.Cdn")`,
		},
		{
			Name:           "profiles",
			Struct:         &armcdn.Profile{},
			ResponseStruct: &armcdn.ProfilesClientListResponse{},
			Client:         &armcdn.ProfilesClient{},
			ListFunc:       (&armcdn.ProfilesClient{}).NewListPager,
			NewFunc:        armcdn.NewProfilesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/profiles",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Cdn")`,
		},
		{
			Name:           "edge_nodes",
			Struct:         &armcdn.EdgeNode{},
			ResponseStruct: &armcdn.EdgeNodesClientListResponse{},
			Client:         &armcdn.EdgeNodesClient{},
			ListFunc:       (&armcdn.EdgeNodesClient{}).NewListPager,
			NewFunc:        armcdn.NewEdgeNodesClient,
			URL:            "/providers/Microsoft.Cdn/edgenodes",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Cdn")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armcdn"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcdn()...)
}

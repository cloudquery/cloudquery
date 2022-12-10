// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"

func Armcdn() []Table {
	tables := []Table{
		{
			Service:        "armcdn",
			Name:           "edge_nodes",
			Struct:         &armcdn.EdgeNode{},
			ResponseStruct: &armcdn.EdgeNodesClientListResponse{},
			Client:         &armcdn.EdgeNodesClient{},
			ListFunc:       (&armcdn.EdgeNodesClient{}).NewListPager,
			NewFunc:        armcdn.NewEdgeNodesClient,
			URL:            "/providers/Microsoft.Cdn/edgenodes",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Cdn)`,
		},
		{
			Service:        "armcdn",
			Name:           "managed_rule_sets",
			Struct:         &armcdn.ManagedRuleSetDefinition{},
			ResponseStruct: &armcdn.ManagedRuleSetsClientListResponse{},
			Client:         &armcdn.ManagedRuleSetsClient{},
			ListFunc:       (&armcdn.ManagedRuleSetsClient{}).NewListPager,
			NewFunc:        armcdn.NewManagedRuleSetsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/cdnWebApplicationFirewallManagedRuleSets",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Cdn)`,
		},
		{
			Service:        "armcdn",
			Name:           "policies",
			Struct:         &armcdn.WebApplicationFirewallPolicy{},
			ResponseStruct: &armcdn.PoliciesClientListResponse{},
			Client:         &armcdn.PoliciesClient{},
			ListFunc:       (&armcdn.PoliciesClient{}).NewListPager,
			NewFunc:        armcdn.NewPoliciesClient,
			URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/cdnWebApplicationFirewallPolicies",
			Multiplex:      `client.SubscriptionResourceGroupMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Cdn)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armcdn()...)
}

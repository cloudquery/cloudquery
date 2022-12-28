package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"

func ArmcdnE() []Table {
	tables := []Table{
		{
			Name:           "profiles",
			Service:        "armcdn",
			Struct:         &armcdn.Profile{},
			ResponseStruct: &armcdn.ProfilesClientListResponse{},
			Client:         &armcdn.ProfilesClient{},
			ListFunc:       (&armcdn.ProfilesClient{}).NewListPager,
			NewFunc:        armcdn.NewProfilesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/profiles",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_cdn)`,
			Relations: []*Table{
				{
					Name:           "endpoints",
					Service:        "armcdn",
					Struct:         &armcdn.Endpoint{},
					ResponseStruct: &armcdn.EndpointsClientListByProfileResponse{},
					Client:         &armcdn.EndpointsClient{},
					ListFunc:       (&armcdn.EndpointsClient{}).NewListByProfilePager,
					NewFunc:        armcdn.NewEndpointsClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints",
					SkipFetch:      true,
				},
				{
					Name:           "rule_sets",
					Service:        "armcdn",
					Struct:         &armcdn.RuleSet{},
					ResponseStruct: &armcdn.RuleSetsClientListByProfileResponse{},
					Client:         &armcdn.RuleSetsClient{},
					ListFunc:       (&armcdn.RuleSetsClient{}).NewListByProfilePager,
					NewFunc:        armcdn.NewRuleSetsClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets",
					SkipFetch:      true,
				},
				{
					Name:           "security_policies",
					Service:        "armcdn",
					Struct:         &armcdn.SecurityPolicy{},
					ResponseStruct: &armcdn.SecurityPoliciesClientListByProfileResponse{},
					Client:         &armcdn.SecurityPoliciesClient{},
					ListFunc:       (&armcdn.SecurityPoliciesClient{}).NewListByProfilePager,
					NewFunc:        armcdn.NewSecurityPoliciesClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/securityPolicies",
					SkipFetch:      true,
				},
			},
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, ArmcdnE()...)
}

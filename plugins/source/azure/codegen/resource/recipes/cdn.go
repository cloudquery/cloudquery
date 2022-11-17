package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func CDN() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct:   new(armcdn.Profile),
			Resolver: cdn.ProfilesClient.NewListPager,
			Children: []*resource.Resource{
				{
					Struct:   new(armcdn.Endpoint),
					Resolver: cdn.EndpointsClient.NewListByProfilePager,
					Children: []*resource.Resource{
						{
							Struct:   new(armcdn.CustomDomain),
							Resolver: cdn.CustomDomainsClient.NewListByEndpointPager,
						},
						{
							Struct:   new(armcdn.Route),
							Resolver: cdn.RoutesClient.NewListByEndpointPager,
						},
					},
				},
				{
					Struct:   new(armcdn.RuleSet),
					Resolver: cdn.RuleSetsClient.NewListByProfilePager,
					Children: []*resource.Resource{
						{
							Struct:   new(armcdn.Rule),
							Resolver: cdn.RulesClient.NewListByRuleSetPager,
						},
					},
				},
				{
					Struct:   new(armcdn.SecurityPolicy),
					Resolver: cdn.SecurityPoliciesClient.NewListByProfilePager,
				},
			},
		},
	}
}

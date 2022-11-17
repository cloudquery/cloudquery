package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func CDN() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armcdn.Profile),
			Resolver: &resource.FuncParams{
				Func: cdn.ProfilesClient.NewListPager,
			},
			Children: []*resource.Resource{
				{
					Struct: new(armcdn.Endpoint),
					Resolver: &resource.FuncParams{
						Func:   cdn.EndpointsClient.NewListByProfilePager,
						Params: []string{"id.ResourceGroupName", "*profile.Name"},
					},
					Children: []*resource.Resource{
						{
							Struct: new(armcdn.CustomDomain),
							Resolver: &resource.FuncParams{
								Func:   cdn.CustomDomainsClient.NewListByEndpointPager,
								Params: []string{"id.ResourceGroupName", "*profile.Name", "*endpoint.Name"},
							},
						},
						{
							Struct: new(armcdn.Route),
							Resolver: &resource.FuncParams{
								Func:   cdn.RoutesClient.NewListByEndpointPager,
								Params: []string{"id.ResourceGroupName", "*profile.Name", "*endpoint.Name"},
							},
						},
					},
				},
				{
					Struct: new(armcdn.RuleSet),
					Resolver: &resource.FuncParams{
						Func:   cdn.RuleSetsClient.NewListByProfilePager,
						Params: []string{"id.ResourceGroupName", "*profile.Name"},
					},
					Children: []*resource.Resource{
						{
							Struct: new(armcdn.Rule),
							Resolver: &resource.FuncParams{
								Func:   cdn.RulesClient.NewListByRuleSetPager,
								Params: []string{"id.ResourceGroupName", "*profile.Name", "*ruleSet.Name"},
							},
							IgnoreInTestColumns: []string{"properties_actions", "properties_conditions"},
						},
					},
				},
				{
					Struct: new(armcdn.SecurityPolicy),
					Resolver: &resource.FuncParams{
						Func:   cdn.SecurityPoliciesClient.NewListByProfilePager,
						Params: []string{"id.ResourceGroupName", "*profile.Name"},
					},
				},
			},
		},
	}
}

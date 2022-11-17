package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/authorization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Authorization() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armauthorization.RoleAssignment),
			Resolver: &resource.FuncParams{
				Func: authorization.RoleAssignmentsClient.NewListForSubscriptionPager,
			},
		},
		{
			Struct: new(armauthorization.RoleDefinition),
			Resolver: &resource.FuncParams{
				Func:   authorization.RoleDefinitionsClient.NewListPager,
				Params: []string{"c.ScopeSubscription()"},
			},
		},
	}
}

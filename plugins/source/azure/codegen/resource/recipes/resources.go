package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/links"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/policy"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/resources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Resource() []*resource.Resource {
	return []*resource.Resource{
		{
			Service:    "resource",
			SubService: "groups",
			Struct:     new(armresources.ResourceGroup),
			Resolver: &resource.FuncParams{
				Func: resources.ResourceGroupsClient.NewListPager,
			},
		},
		{
			Service:    "resource",
			SubService: "links",
			Struct:     new(armlinks.ResourceLink),
			Resolver: &resource.FuncParams{
				Func: links.ResourceLinksClient.NewListAtSubscriptionPager,
			},
		},
		{
			Service:    "resource",
			SubService: "policy_assignments",
			Struct:     new(armpolicy.Assignment),
			Resolver: &resource.FuncParams{
				Func: policy.AssignmentsClient.NewListPager,
			},
		},
	}
}

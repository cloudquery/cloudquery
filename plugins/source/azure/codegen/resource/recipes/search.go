package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/search"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Search() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armsearch.Service),
			Resolver: &resource.FuncParams{
				Func:   search.ServicesClient.NewListBySubscriptionPager,
				Params: []string{"nil"},
			},
		},
	}
}

package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/frontdoor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func FrontDoor() []*resource.Resource {
	return []*resource.Resource{
		{
			SubService: "doors",
			Name:       "azure_front_doors",
			Struct:     new(armfrontdoor.FrontDoor),
			Resolver: &resource.FuncParams{
				Func: frontdoor.FrontDoorsClient.NewListPager,
			},
		},
	}
}

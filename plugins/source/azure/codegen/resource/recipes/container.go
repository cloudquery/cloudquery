package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/containerregistry"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/containerservice"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func Container() []*resource.Resource {
	return []*resource.Resource{
		{
			Service: "container",
			Struct:  new(armcontainerregistry.Registry),
			Resolver: &resource.FuncParams{
				Func: containerregistry.RegistriesClient.NewListPager,
			},
			Children: []*resource.Resource{
				{
					Struct: new(armcontainerregistry.Replication),
					Resolver: &resource.FuncParams{
						Func:   containerregistry.ReplicationsClient.NewListPager,
						Params: []string{"id.ResourceGroupName", "*registry.Name"},
					},
				},
			},
		},
		{
			Service: "container",
			Struct:  new(armcontainerservice.ManagedCluster),
			Resolver: &resource.FuncParams{
				Func: containerservice.ManagedClustersClient.NewListPager,
			},
		},
	}
}

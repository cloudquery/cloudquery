package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

func ContainerServiceResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "managed_clusters",
			Struct: &armcontainerservice.ManagedCluster{},
			ResponseStruct: &armcontainerservice.ManagedClustersClientListResponse{},
			Client: &armcontainerservice.ManagedClustersClient{},
			ListFunc: (&armcontainerservice.ManagedClustersClient{}).NewListPager,
			NewFunc: armcontainerservice.NewManagedClustersClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "containerservice/armcontainerservice"
		r.Service = "armcontainerservice"
		r.Template = "list"
	}

	return resources
}
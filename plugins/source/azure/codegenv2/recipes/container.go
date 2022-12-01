package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

func ContainerResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "registries",
			Struct: &armcontainerregistry.Registry{},
			ResponseStruct: &armcontainerregistry.RegistriesClientListResponse{},
			Client: &armcontainerregistry.RegistriesClient{},
			ListFunc: (&armcontainerregistry.RegistriesClient{}).NewListPager,
			NewFunc: armcontainerregistry.NewRegistriesClient,
			OutputField: "Value",
		},
		// {
		// 	SubService: "replications",
		// 	Struct: &armcontainerregistry.Replication{},
		// 	ResponseStruct: &armcontainerregistry.ReplicationsClientListResponse{},
		// 	Client: &armcontainerregistry.ReplicationsClient{},
		// 	ListFunc: (&armcontainerregistry.ReplicationsClient{}).NewListPager,
		// 	NewFunc: armcontainerregistry.NewReplicationsClient,
		// 	OutputField: "Value",
		// },
	}

	for _, r := range resources {
		r.ImportPath = "containerregistry/armcontainerregistry"
		r.Service = "armcontainerregistry"
		r.Template = "list"
	}

	return resources
}
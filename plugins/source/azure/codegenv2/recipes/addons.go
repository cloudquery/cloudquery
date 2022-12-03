package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/addons/armaddons"
)

func AddonsResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "private_link_policies",
			Struct: &armaddons.OperationListValue{},
			ResponseStruct: &armaddons.OperationsClientListResponse{},
			Client: &armaddons.OperationsClient{},
			ListFunc: (&armaddons.OperationsClient{}).List,
			NewFunc: armaddons.NewOperationsClient,
		},
	}

	for _, r := range resources {
		r.ImportPath = "addons/armaddons"
		r.Service = "armaddons"
		r.Template = "list"
	}

	return resources
}
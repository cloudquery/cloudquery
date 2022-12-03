package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

func AdvisorResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "operation_entities",
			Struct: &armadvisor.OperationEntity{},
			ResponseStruct: &armadvisor.OperationsClientListResponse{},
			Client: &armadvisor.OperationsClient{},
			ListFunc: (&armadvisor.OperationsClient{}).NewListPager,
			NewFunc: armadvisor.NewOperationsClient,
		},
		{
			SubService: "config_data",
			Struct: &armadvisor.ConfigData{},
			ResponseStruct: &armadvisor.ConfigurationsClientListBySubscriptionResponse{},
			Client: &armadvisor.ConfigurationsClient{},
			ListFunc: (&armadvisor.ConfigurationsClient{}).NewListBySubscriptionPager,
			NewFunc: armadvisor.NewConfigurationsClient,
		},
	}

	for _, r := range resources {
		r.ImportPath = "advisor/armadvisor"
		r.Service = "armadvisor"
		r.Template = "list"
	}

	return resources
}
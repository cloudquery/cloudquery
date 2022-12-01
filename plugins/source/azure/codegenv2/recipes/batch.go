package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"

func BatchResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "accounts",
			Struct: &armbatch.Account{},
			ResponseStruct: &armbatch.AccountClientListResponse{},
			Client: &armbatch.AccountClient{},
			ListFunc: (&armbatch.AccountClient{}).NewListPager,
			NewFunc: armbatch.NewAccountClient,
			OutputField: "Value",
		},
		// {
		// 	SubService: "applications",
		// 	Struct: &armbatch.Application{},
		// 	ResponseStruct: &armbatch.ApplicationClientListResponse{},
		// 	Client: &armbatch.ApplicationClient{},
		// 	ListFunc: (&armbatch.ApplicationClient{}).NewListPager,
		// 	NewFunc: armbatch.NewApplicationClient,
		// 	OutputField: "Value",
		// },
	}

	for _, r := range resources {
		r.ImportPath = "batch/armbatch"
		r.Service = "armbatch"
		r.Template = "list"
	}

	return resources
}
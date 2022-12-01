package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"

func CDNResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "profiles",
			Struct: &armcdn.Profile{},
			ResponseStruct: &armcdn.ProfilesClientListResponse{},
			Client: &armcdn.ProfilesClient{},
			ListFunc: (&armcdn.ProfilesClient{}).NewListPager,
			NewFunc: armcdn.NewProfilesClient,
			OutputField: "Value",
			Relations: []string{"Endpoints()"},
		},
		{
			SubService: "endpoints",
			ChildTable: true,
			Struct: &armcdn.Endpoint{},
			ResponseStruct: &armcdn.EndpointsClientListByProfileResponse{},
			Client: &armcdn.EndpointsClient{},
			ListFunc: (&armcdn.EndpointsClient{}).NewListByProfilePager,
			NewFunc: armcdn.NewEndpointsClient,
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
		r.ImportPath = "cdn/armcdn"
		r.Service = "armcdn"
		r.Template = "list"
	}

	return resources
}
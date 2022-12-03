package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

func AadResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "private_link_policies",
			Struct: &armaad.PrivateLinkPolicy{},
			ResponseStruct: &armaad.PrivateLinkForAzureAdClientListResponse{},
			Client: &armaad.PrivateLinkForAzureAdClient{},
			ListFunc: (&armaad.PrivateLinkForAzureAdClient{}).NewListPager,
			NewFunc: armaad.NewPrivateLinkForAzureAdClient,
		},
	}

	for _, r := range resources {
		r.ImportPath = "aad/armaad"
		r.Service = "armaad"
		r.Template = "list"
	}

	return resources
}
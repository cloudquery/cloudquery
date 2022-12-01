package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

func FrontDoorResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "doors",
			Struct: &armfrontdoor.FrontDoor{},
			ResponseStruct: &armfrontdoor.FrontDoorsClientListResponse{},
			Client: &armfrontdoor.FrontDoorsClient{},
			ListFunc: (&armfrontdoor.FrontDoorsClient{}).NewListPager,
			NewFunc: armfrontdoor.NewFrontDoorsClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "frontdoor/armfrontdoor"
		r.Service = "armfrontdoor"
		r.Template = "list"
	}

	return resources
}
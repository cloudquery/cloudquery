package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

func IotSecurityResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "devices",
			Struct: &armiotsecurity.DeviceModel{},
			ResponseStruct: &armiotsecurity.DevicesClientListResponse{},
			Client: &armiotsecurity.DevicesClient{},
			ListFunc: (&armiotsecurity.DevicesClient{}).NewListPager,
			NewFunc: armiotsecurity.NewDevicesClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "iotsecurity/armiotsecurity"
		r.Service = "armiotsecurity"
		r.Template = "list"
	}

	return resources
}
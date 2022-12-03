package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

func MonitorResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "servers",
			Struct: &armmonitor.ActivityLogAlertResource{},
			ResponseStruct: &armmonitor.ActivityLogAlertsClientListByResourceGroupResponse{},
			Client: &armmonitor.ActivityLogAlertsClient{},
			ListFunc: (&armmonitor.ActivityLogAlertsClient{}).NewListBySubscriptionIDPager,
			NewFunc: armmariadb.NewServersClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "mariadb/armmariadb"
		r.Service = "armmariadb"
		r.Template = "list"
	}

	return resources
}
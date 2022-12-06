// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"

func Armcustomerinsights() []Table {
	tables := []Table{
		{
			Name:           "hub",
			Struct:         &armcustomerinsights.Hub{},
			ResponseStruct: &armcustomerinsights.HubsClientListResponse{},
			Client:         &armcustomerinsights.HubsClient{},
			ListFunc:       (&armcustomerinsights.HubsClient{}).NewListPager,
			NewFunc:        armcustomerinsights.NewHubsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.CustomerInsights/hubs",
		},
	}

	for i := range tables {
		tables[i].Service = "armcustomerinsights"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcustomerinsights()...)
}

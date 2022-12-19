// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"

func init() {
	tables := []Table{
		{
			Service:        "armcustomerinsights",
			Name:           "hubs",
			Struct:         &armcustomerinsights.Hub{},
			ResponseStruct: &armcustomerinsights.HubsClientListResponse{},
			Client:         &armcustomerinsights.HubsClient{},
			ListFunc:       (&armcustomerinsights.HubsClient{}).NewListPager,
			NewFunc:        armcustomerinsights.NewHubsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.CustomerInsights/hubs",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_customerinsights)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}

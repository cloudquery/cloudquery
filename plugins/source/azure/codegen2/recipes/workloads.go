// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"

func init() {
	tables := []Table{
		{
			Service:        "armworkloads",
			Name:           "monitors",
			Struct:         &armworkloads.Monitor{},
			ResponseStruct: &armworkloads.MonitorsClientListResponse{},
			Client:         &armworkloads.MonitorsClient{},
			ListFunc:       (&armworkloads.MonitorsClient{}).NewListPager,
			NewFunc:        armworkloads.NewMonitorsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/monitors",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Workloads)`,
		},
	}
	Tables = append(Tables, tables...)
}

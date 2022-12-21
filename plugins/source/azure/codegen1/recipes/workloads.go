// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"

func Armworkloads() []*Table {
	tables := []*Table{
		{
			NewFunc:        armworkloads.NewMonitorsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/monitors",
			Namespace:      "microsoft.workloads",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_workloads)`,
			Pager:          `NewListPager`,
			ResponseStruct: "MonitorsClientListResponse",
		},
		{
			NewFunc:        armworkloads.NewSKUsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/skus",
			Namespace:      "microsoft.workloads",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_workloads)`,
			Pager:          `NewListPager`,
			ResponseStruct: "SKUsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armworkloads())
}

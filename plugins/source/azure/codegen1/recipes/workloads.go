// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"

func Armworkloads() []*Table {
	tables := []*Table{
		{
			NewFunc: armworkloads.NewSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/skus",
		},
		{
			NewFunc: armworkloads.NewMonitorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/monitors",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armworkloads())
}

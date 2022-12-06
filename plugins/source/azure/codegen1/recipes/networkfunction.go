// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction"

func Armnetworkfunction() []*Table {
	tables := []*Table{
		{
			NewFunc: armnetworkfunction.NewAzureTrafficCollectorsBySubscriptionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkFunction/azureTrafficCollectors",
		},
		{
			NewFunc: armnetworkfunction.NewAzureTrafficCollectorsByResourceGroupClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkFunction/azureTrafficCollectors",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armnetworkfunction())
}
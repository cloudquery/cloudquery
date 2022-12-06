// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction"

func Armnetworkfunction() []Table {
	tables := []Table{
		{
      Name: "azure_traffic_collector",
      Struct: &armnetworkfunction.AzureTrafficCollector{},
      ResponseStruct: &armnetworkfunction.AzureTrafficCollectorsByResourceGroupClientListResponse{},
      Client: &armnetworkfunction.AzureTrafficCollectorsByResourceGroupClient{},
      ListFunc: (&armnetworkfunction.AzureTrafficCollectorsByResourceGroupClient{}).NewListPager,
			NewFunc: armnetworkfunction.NewAzureTrafficCollectorsByResourceGroupClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkFunction/azureTrafficCollectors",
		},
		{
      Name: "azure_traffic_collector",
      Struct: &armnetworkfunction.AzureTrafficCollector{},
      ResponseStruct: &armnetworkfunction.AzureTrafficCollectorsBySubscriptionClientListResponse{},
      Client: &armnetworkfunction.AzureTrafficCollectorsBySubscriptionClient{},
      ListFunc: (&armnetworkfunction.AzureTrafficCollectorsBySubscriptionClient{}).NewListPager,
			NewFunc: armnetworkfunction.NewAzureTrafficCollectorsBySubscriptionClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkFunction/azureTrafficCollectors",
		},
	}

	for i := range tables {
		tables[i].Service = "armnetworkfunction"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armnetworkfunction()...)
}
// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"

func Armhybridnetwork() []Table {
	tables := []Table{
		{
      Name: "network_function_vendor",
      Struct: &armhybridnetwork.NetworkFunctionVendor{},
      ResponseStruct: &armhybridnetwork.NetworkFunctionVendorsClientListResponse{},
      Client: &armhybridnetwork.NetworkFunctionVendorsClient{},
      ListFunc: (&armhybridnetwork.NetworkFunctionVendorsClient{}).NewListPager,
			NewFunc: armhybridnetwork.NewNetworkFunctionVendorsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/networkFunctionVendors",
		},
	}

	for i := range tables {
		tables[i].Service = "armhybridnetwork"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armhybridnetwork()...)
}
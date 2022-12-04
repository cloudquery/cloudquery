// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"

func Armhybridnetwork() []Table {
	tables := []Table{
		{
      Name: "preview_subscription",
      Struct: &armhybridnetwork.PreviewSubscription{},
      ResponseStruct: &armhybridnetwork.VendorSKUPreviewClientListResponse{},
      Client: &armhybridnetwork.VendorSKUPreviewClient{},
      ListFunc: (&armhybridnetwork.VendorSKUPreviewClient{}).NewListPager,
			NewFunc: armhybridnetwork.NewVendorSKUPreviewClient,
		},
		{
      Name: "vendor_sku",
      Struct: &armhybridnetwork.VendorSKU{},
      ResponseStruct: &armhybridnetwork.VendorSKUsClientListResponse{},
      Client: &armhybridnetwork.VendorSKUsClient{},
      ListFunc: (&armhybridnetwork.VendorSKUsClient{}).NewListPager,
			NewFunc: armhybridnetwork.NewVendorSKUsClient,
		},
		{
      Name: "operation",
      Struct: &armhybridnetwork.Operation{},
      ResponseStruct: &armhybridnetwork.OperationsClientListResponse{},
      Client: &armhybridnetwork.OperationsClient{},
      ListFunc: (&armhybridnetwork.OperationsClient{}).NewListPager,
			NewFunc: armhybridnetwork.NewOperationsClient,
		},
		{
      Name: "role_instance",
      Struct: &armhybridnetwork.RoleInstance{},
      ResponseStruct: &armhybridnetwork.RoleInstancesClientListResponse{},
      Client: &armhybridnetwork.RoleInstancesClient{},
      ListFunc: (&armhybridnetwork.RoleInstancesClient{}).NewListPager,
			NewFunc: armhybridnetwork.NewRoleInstancesClient,
		},
		{
      Name: "vendor_network_function",
      Struct: &armhybridnetwork.VendorNetworkFunction{},
      ResponseStruct: &armhybridnetwork.VendorNetworkFunctionsClientListResponse{},
      Client: &armhybridnetwork.VendorNetworkFunctionsClient{},
      ListFunc: (&armhybridnetwork.VendorNetworkFunctionsClient{}).NewListPager,
			NewFunc: armhybridnetwork.NewVendorNetworkFunctionsClient,
		},
		{
      Name: "network_function_vendor",
      Struct: &armhybridnetwork.NetworkFunctionVendor{},
      ResponseStruct: &armhybridnetwork.NetworkFunctionVendorsClientListResponse{},
      Client: &armhybridnetwork.NetworkFunctionVendorsClient{},
      ListFunc: (&armhybridnetwork.NetworkFunctionVendorsClient{}).NewListPager,
			NewFunc: armhybridnetwork.NewNetworkFunctionVendorsClient,
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
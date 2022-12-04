// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"

func Armhybridcompute() []Table {
	tables := []Table{
		{
      Name: "machine_extension",
      Struct: &armhybridcompute.MachineExtension{},
      ResponseStruct: &armhybridcompute.MachineExtensionsClientListResponse{},
      Client: &armhybridcompute.MachineExtensionsClient{},
      ListFunc: (&armhybridcompute.MachineExtensionsClient{}).NewListPager,
			NewFunc: armhybridcompute.NewMachineExtensionsClient,
		},
		{
      Name: "operation_value",
      Struct: &armhybridcompute.OperationValue{},
      ResponseStruct: &armhybridcompute.OperationsClientListResponse{},
      Client: &armhybridcompute.OperationsClient{},
      ListFunc: (&armhybridcompute.OperationsClient{}).NewListPager,
			NewFunc: armhybridcompute.NewOperationsClient,
		},
		{
      Name: "private_link_scope",
      Struct: &armhybridcompute.PrivateLinkScope{},
      ResponseStruct: &armhybridcompute.PrivateLinkScopesClientListResponse{},
      Client: &armhybridcompute.PrivateLinkScopesClient{},
      ListFunc: (&armhybridcompute.PrivateLinkScopesClient{}).NewListPager,
			NewFunc: armhybridcompute.NewPrivateLinkScopesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armhybridcompute"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armhybridcompute()...)
}
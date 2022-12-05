// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementpartner/armmanagementpartner"

func Armmanagementpartner() []Table {
	tables := []Table{
		{
      Name: "operation_response",
      Struct: &armmanagementpartner.OperationResponse{},
      ResponseStruct: &armmanagementpartner.OperationClientListResponse{},
      Client: &armmanagementpartner.OperationClient{},
      ListFunc: (&armmanagementpartner.OperationClient{}).NewListPager,
			NewFunc: armmanagementpartner.NewOperationClient,
			URL: "/providers/Microsoft.ManagementPartner/operations",
		},
	}

	for i := range tables {
		tables[i].Service = "armmanagementpartner"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armmanagementpartner()...)
}
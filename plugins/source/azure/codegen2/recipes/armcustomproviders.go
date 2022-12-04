// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders"

func Armcustomproviders() []Table {
	tables := []Table{
		{
      Name: "resource_provider_operation",
      Struct: &armcustomproviders.ResourceProviderOperation{},
      ResponseStruct: &armcustomproviders.OperationsClientListResponse{},
      Client: &armcustomproviders.OperationsClient{},
      ListFunc: (&armcustomproviders.OperationsClient{}).NewListPager,
			NewFunc: armcustomproviders.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armcustomproviders"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armcustomproviders()...)
}
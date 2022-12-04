// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"

func Armdeviceprovisioningservices() []Table {
	tables := []Table{
		{
      Name: "operation",
      Struct: &armdeviceprovisioningservices.Operation{},
      ResponseStruct: &armdeviceprovisioningservices.OperationsClientListResponse{},
      Client: &armdeviceprovisioningservices.OperationsClient{},
      ListFunc: (&armdeviceprovisioningservices.OperationsClient{}).NewListPager,
			NewFunc: armdeviceprovisioningservices.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armdeviceprovisioningservices"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdeviceprovisioningservices()...)
}
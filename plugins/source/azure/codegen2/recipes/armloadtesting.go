// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting"

func Armloadtesting() []Table {
	tables := []Table{
		{
      Name: "operation",
      Struct: &armloadtesting.Operation{},
      ResponseStruct: &armloadtesting.OperationsClientListResponse{},
      Client: &armloadtesting.OperationsClient{},
      ListFunc: (&armloadtesting.OperationsClient{}).NewListPager,
			NewFunc: armloadtesting.NewOperationsClient,
		},
		{
      Name: "quota_resource",
      Struct: &armloadtesting.QuotaResource{},
      ResponseStruct: &armloadtesting.QuotasClientListResponse{},
      Client: &armloadtesting.QuotasClient{},
      ListFunc: (&armloadtesting.QuotasClient{}).NewListPager,
			NewFunc: armloadtesting.NewQuotasClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armloadtesting"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armloadtesting()...)
}
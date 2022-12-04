// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"

func Armagrifood() []Table {
	tables := []Table{
		{
      Name: "farm_beats_extension",
      Struct: &armagrifood.FarmBeatsExtension{},
      ResponseStruct: &armagrifood.FarmBeatsExtensionsClientListResponse{},
      Client: &armagrifood.FarmBeatsExtensionsClient{},
      ListFunc: (&armagrifood.FarmBeatsExtensionsClient{}).NewListPager,
			NewFunc: armagrifood.NewFarmBeatsExtensionsClient,
		},
		{
      Name: "operation",
      Struct: &armagrifood.Operation{},
      ResponseStruct: &armagrifood.OperationsClientListResponse{},
      Client: &armagrifood.OperationsClient{},
      ListFunc: (&armagrifood.OperationsClient{}).NewListPager,
			NewFunc: armagrifood.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armagrifood"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armagrifood()...)
}
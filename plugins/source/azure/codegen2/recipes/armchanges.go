// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armchanges"

func Armchanges() []Table {
	tables := []Table{
		{
      Name: "change_resource_result",
      Struct: &armchanges.ChangeResourceResult{},
      ResponseStruct: &armchanges.ClientListResponse{},
      Client: &armchanges.Client{},
      ListFunc: (&armchanges.Client{}).NewListPager,
			NewFunc: armchanges.NewClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armchanges"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armchanges()...)
}
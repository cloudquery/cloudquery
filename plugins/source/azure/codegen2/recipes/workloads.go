// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"

func Armworkloads() []Table {
	tables := []Table{
		{
      Name: "sku_definition",
      Struct: &armworkloads.SKUDefinition{},
      ResponseStruct: &armworkloads.SKUsClientListResponse{},
      Client: &armworkloads.SKUsClient{},
      ListFunc: (&armworkloads.SKUsClient{}).NewListPager,
			NewFunc: armworkloads.NewSKUsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/skus",
		},
		{
      Name: "monitor",
      Struct: &armworkloads.Monitor{},
      ResponseStruct: &armworkloads.MonitorsClientListResponse{},
      Client: &armworkloads.MonitorsClient{},
      ListFunc: (&armworkloads.MonitorsClient{}).NewListPager,
			NewFunc: armworkloads.NewMonitorsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/monitors",
		},
	}

	for i := range tables {
		tables[i].Service = "armworkloads"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armworkloads()...)
}
// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"

func Armpowerbidedicated() []Table {
	tables := []Table{
		{
      Name: "dedicated_capacity",
      Struct: &armpowerbidedicated.DedicatedCapacity{},
      ResponseStruct: &armpowerbidedicated.CapacitiesClientListResponse{},
      Client: &armpowerbidedicated.CapacitiesClient{},
      ListFunc: (&armpowerbidedicated.CapacitiesClient{}).NewListPager,
			NewFunc: armpowerbidedicated.NewCapacitiesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/capacities",
		},
	}

	for i := range tables {
		tables[i].Service = "armpowerbidedicated"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armpowerbidedicated()...)
}
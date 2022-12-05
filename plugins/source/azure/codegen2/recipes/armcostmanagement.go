// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"

func Armcostmanagement() []Table {
	tables := []Table{
		{
      Name: "view",
      Struct: &armcostmanagement.View{},
      ResponseStruct: &armcostmanagement.ViewsClientListResponse{},
      Client: &armcostmanagement.ViewsClient{},
      ListFunc: (&armcostmanagement.ViewsClient{}).NewListPager,
			NewFunc: armcostmanagement.NewViewsClient,
			URL: "/providers/Microsoft.CostManagement/views",
		},
		{
      Name: "dimension",
      Struct: &armcostmanagement.Dimension{},
      ResponseStruct: &armcostmanagement.DimensionsClientListResponse{},
      Client: &armcostmanagement.DimensionsClient{},
      ListFunc: (&armcostmanagement.DimensionsClient{}).NewListPager,
			NewFunc: armcostmanagement.NewDimensionsClient,
			URL: "/{scope}/providers/Microsoft.CostManagement/dimensions",
		},
	}

	for i := range tables {
		tables[i].Service = "armcostmanagement"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armcostmanagement()...)
}
// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"

func init() {
	tables := []Table{
		{
			Service:        "armcostmanagement",
			Name:           "views",
			Struct:         &armcostmanagement.View{},
			ResponseStruct: &armcostmanagement.ViewsClientListResponse{},
			Client:         &armcostmanagement.ViewsClient{},
			ListFunc:       (&armcostmanagement.ViewsClient{}).NewListPager,
			NewFunc:        armcostmanagement.NewViewsClient,
			URL:            "/providers/Microsoft.CostManagement/views",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_costmanagement)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}

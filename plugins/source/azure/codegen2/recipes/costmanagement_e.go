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
			ListFunc:       (&armcostmanagement.ViewsClient{}).NewListByScopePager,
			NewFunc:        armcostmanagement.NewViewsClient,
			URL:            "/providers/Microsoft.CostManagement/views",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_costmanagement)`,
			ExtraColumns:   DefaultExtraColumns,
			SkipFetch:      true,
			SkipMock:       true,
			Relations: []*Table{
				{
					Service:        "armcostmanagement",
					Name:           "view_queries",
					Struct:         &armcostmanagement.QueryResult{},
					ResponseStruct: &armcostmanagement.QueryClientUsageResponse{},
					Client:         &armcostmanagement.QueryClient{},
					ListFunc:       (&armcostmanagement.QueryClient{}).Usage,
					NewFunc:        armcostmanagement.NewQueryClient,
					URL:            "/providers/Microsoft.CostManagement/query",
					SkipFetch:      true,
					SkipMock:       true,
				},
			},
		},
	}
	Tables = append(Tables, tables...)
}

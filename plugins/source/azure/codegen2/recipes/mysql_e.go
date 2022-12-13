package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"

func init() {
	tables := []Table{
		{
			Service:        "armmysql",
			Name:           "servers",
			Struct:         &armmysql.Server{},
			ResponseStruct: &armmysql.ServersClientListResponse{},
			Client:         &armmysql.ServersClient{},
			ListFunc:       (&armmysql.ServersClient{}).NewListPager,
			NewFunc:        armmysql.NewServersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/servers",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DBforMySQL)`,
			ExtraColumns:   DefaultExtraColumns,
			Relations: []*Table{
				{
					Service:        "armmysql",
					Name:           "server_configurations",
					Struct:         &armmysql.Configuration{},
					ResponseStruct: &armmysql.ConfigurationsClientListByServerResponse{},
					Client:         &armmysql.ConfigurationsClient{},
					ListFunc:       (&armmysql.ConfigurationsClient{}).NewListByServerPager,
					NewFunc:        armmysql.NewConfigurationsClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/configurations",
					SkipFetch:      true,
				},
			},
		},
	}
	Tables = append(Tables, tables...)
}

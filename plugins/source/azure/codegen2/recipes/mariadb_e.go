package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"

func init() {
	tables := []Table{
		{
			Service:        "armmariadb",
			Name:           "servers",
			Struct:         &armmariadb.Server{},
			ResponseStruct: &armmariadb.ServersClientListResponse{},
			Client:         &armmariadb.ServersClient{},
			ListFunc:       (&armmariadb.ServersClient{}).NewListPager,
			NewFunc:        armmariadb.NewServersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMariaDB/servers",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DBforMariaDB)`,
			ExtraColumns:   DefaultExtraColumns,
			Relations: []*Table{
				{
					Service:        "armmariadb",
					Name:           "server_configurations",
					Struct:         &armmariadb.Configuration{},
					ResponseStruct: &armmariadb.ConfigurationsClientListByServerResponse{},
					Client:         &armmariadb.ConfigurationsClient{},
					ListFunc:       (&armmariadb.ConfigurationsClient{}).NewListByServerPager,
					NewFunc:        armmariadb.NewConfigurationsClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/configurations",
					SkipFetch:      true,
				},
			},
		},
	}
	Tables = append(Tables, tables...)
}

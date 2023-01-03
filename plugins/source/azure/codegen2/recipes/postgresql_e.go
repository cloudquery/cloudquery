package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

func PostgresqlE() []Table {
	tables := []Table{
		{
			Service:        "armpostgresql",
			Name:           "servers",
			Struct:         &armpostgresql.Server{},
			ResponseStruct: &armpostgresql.ServersClientListResponse{},
			Client:         &armpostgresql.ServersClient{},
			ListFunc:       (&armpostgresql.ServersClient{}).NewListPager,
			NewFunc:        armpostgresql.NewServersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/servers",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_dbforpostgresql)`,
			ExtraColumns:   DefaultExtraColumns,
			Relations: []*Table{
				{
					Service:        "armpostgresql",
					Name:           "server_configurations",
					Struct:         &armpostgresql.Configuration{},
					ResponseStruct: &armpostgresql.ConfigurationsClientListByServerResponse{},
					Client:         &armpostgresql.ConfigurationsClient{},
					ListFunc:       (&armpostgresql.ConfigurationsClient{}).NewListByServerPager,
					NewFunc:        armpostgresql.NewConfigurationsClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.DBforPostgreSQL/servers/test string/configurations",
					SkipFetch:      true,
				},
			},
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, PostgresqlE()...)
}

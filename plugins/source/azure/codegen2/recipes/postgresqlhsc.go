// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"

func init() {
	tables := []Table{
		{
			Service:        "armpostgresqlhsc",
			Name:           "server_groups",
			Struct:         &armpostgresqlhsc.ServerGroup{},
			ResponseStruct: &armpostgresqlhsc.ServerGroupsClientListResponse{},
			Client:         &armpostgresqlhsc.ServerGroupsClient{},
			ListFunc:       (&armpostgresqlhsc.ServerGroupsClient{}).NewListPager,
			NewFunc:        armpostgresqlhsc.NewServerGroupsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DBForPostgreSql/serverGroupsv2",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DBForPostgreSql)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}

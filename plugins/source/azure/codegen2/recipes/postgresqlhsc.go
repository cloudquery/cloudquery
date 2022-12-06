// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"

func Armpostgresqlhsc() []Table {
	tables := []Table{
		{
			Name:           "server_group",
			Struct:         &armpostgresqlhsc.ServerGroup{},
			ResponseStruct: &armpostgresqlhsc.ServerGroupsClientListResponse{},
			Client:         &armpostgresqlhsc.ServerGroupsClient{},
			ListFunc:       (&armpostgresqlhsc.ServerGroupsClient{}).NewListPager,
			NewFunc:        armpostgresqlhsc.NewServerGroupsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DBForPostgreSql/serverGroupsv2",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.DBForPostgreSql")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armpostgresqlhsc"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armpostgresqlhsc()...)
}

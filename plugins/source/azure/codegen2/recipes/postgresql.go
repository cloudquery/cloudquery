// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"

func Armpostgresql() []Table {
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
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DBforPostgreSQL)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armpostgresql()...)
}

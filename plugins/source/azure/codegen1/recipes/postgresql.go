// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"

func Armpostgresql() []*Table {
	tables := []*Table{
		{
			NewFunc:        armpostgresql.NewServersClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/servers",
			Namespace:      "microsoft.dbforpostgresql",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_dbforpostgresql)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ServersClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armpostgresql())
}

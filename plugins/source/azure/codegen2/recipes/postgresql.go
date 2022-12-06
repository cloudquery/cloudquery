// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"

func Armpostgresql() []Table {
	tables := []Table{
		{
			Name:           "server",
			Struct:         &armpostgresql.Server{},
			ResponseStruct: &armpostgresql.ServersClientListResponse{},
			Client:         &armpostgresql.ServersClient{},
			ListFunc:       (&armpostgresql.ServersClient{}).NewListPager,
			NewFunc:        armpostgresql.NewServersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/servers",
		},
	}

	for i := range tables {
		tables[i].Service = "armpostgresql"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armpostgresql()...)
}

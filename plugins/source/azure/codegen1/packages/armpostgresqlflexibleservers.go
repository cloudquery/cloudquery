// Code generated by codegen; DO NOT EDIT.
package packages

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers"

func Armpostgresqlflexibleservers() []*Table {
	tables := []*Table{
		{
			NewFunc: armpostgresqlflexibleservers.NewServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/flexibleServers",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armpostgresqlflexibleservers())
}
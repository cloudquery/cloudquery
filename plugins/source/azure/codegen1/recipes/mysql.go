// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"

func Armmysql() []*Table {
	tables := []*Table{
		{
			NewFunc: armmysql.NewServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/servers",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmysql())
}

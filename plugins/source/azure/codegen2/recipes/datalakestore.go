// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"

func Armdatalakestore() []Table {
	tables := []Table{
		{
			Name:           "accounts",
			Struct:         &armdatalakestore.AccountBasic{},
			ResponseStruct: &armdatalakestore.AccountsClientListResponse{},
			Client:         &armdatalakestore.AccountsClient{},
			ListFunc:       (&armdatalakestore.AccountsClient{}).NewListPager,
			NewFunc:        armdatalakestore.NewAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/accounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.DataLakeStore")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armdatalakestore"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdatalakestore()...)
}

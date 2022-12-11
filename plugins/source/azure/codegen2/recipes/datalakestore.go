// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"

func init() {
	tables := []Table{
		{
			Service:        "armdatalakestore",
			Name:           "accounts",
			Struct:         &armdatalakestore.AccountBasic{},
			ResponseStruct: &armdatalakestore.AccountsClientListResponse{},
			Client:         &armdatalakestore.AccountsClient{},
			ListFunc:       (&armdatalakestore.AccountsClient{}).NewListPager,
			NewFunc:        armdatalakestore.NewAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/accounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DataLakeStore)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}

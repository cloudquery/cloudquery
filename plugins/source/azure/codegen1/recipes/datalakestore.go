// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"

func Armdatalakestore() []*Table {
	tables := []*Table{
		{
			NewFunc:        armdatalakestore.NewAccountsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/accounts",
			Namespace:      "Microsoft.DataLakeStore",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DataLakeStore)`,
			Pager:          `NewListPager`,
			ResponseStruct: "AccountsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdatalakestore())
}

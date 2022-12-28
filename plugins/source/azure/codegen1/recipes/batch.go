// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"

func Armbatch() []*Table {
	tables := []*Table{
		{
			NewFunc:        armbatch.NewAccountClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Batch/batchAccounts",
			Namespace:      "microsoft.batch",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_batch)`,
			Pager:          `NewListPager`,
			ResponseStruct: "AccountClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armbatch())
}

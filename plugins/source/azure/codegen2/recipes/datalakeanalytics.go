// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"

func init() {
	tables := []Table{
		{
			Service:        "armdatalakeanalytics",
			Name:           "accounts",
			Struct:         &armdatalakeanalytics.AccountBasic{},
			ResponseStruct: &armdatalakeanalytics.AccountsClientListResponse{},
			Client:         &armdatalakeanalytics.AccountsClient{},
			ListFunc:       (&armdatalakeanalytics.AccountsClient{}).NewListPager,
			NewFunc:        armdatalakeanalytics.NewAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/accounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DataLakeAnalytics)`,
		},
	}
	Tables = append(Tables, tables...)
}

// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"

func Armdatalakeanalytics() []*Table {
	tables := []*Table{
		{
			NewFunc:   armdatalakeanalytics.NewAccountsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/accounts",
			Namespace: "Microsoft.DataLakeAnalytics",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DataLakeAnalytics)`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdatalakeanalytics())
}

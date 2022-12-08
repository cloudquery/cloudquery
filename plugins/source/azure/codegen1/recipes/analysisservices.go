// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"

func Armanalysisservices() []*Table {
	tables := []*Table{
		{
			NewFunc:   armanalysisservices.NewServersClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/servers",
			Namespace: "Microsoft.AnalysisServices",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_AnalysisServices)`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armanalysisservices())
}

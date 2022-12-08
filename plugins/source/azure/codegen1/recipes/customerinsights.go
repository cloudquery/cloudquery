// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"

func Armcustomerinsights() []*Table {
	tables := []*Table{
		{
			NewFunc:   armcustomerinsights.NewHubsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.CustomerInsights/hubs",
			Namespace: "Microsoft.CustomerInsights",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_CustomerInsights)`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcustomerinsights())
}

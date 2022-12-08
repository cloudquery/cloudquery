// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"

func Armaad() []*Table {
	tables := []*Table{
		{
			NewFunc:   armaad.NewPrivateLinkForAzureAdClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad",
			URL:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.aadiam/privateLinkForAzureAd",
			Namespace: "microsoft.aadiam",
			Multiplex: `client.SubscriptionResourceGroupMultiplexRegisteredNamespace(client.Namespacemicrosoft_aadiam)`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armaad())
}

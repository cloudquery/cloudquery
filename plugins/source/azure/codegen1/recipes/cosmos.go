// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"

func Armcosmos() []*Table {
	tables := []*Table{
		{
			NewFunc:   armcosmos.NewDatabaseAccountsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/databaseAccounts",
			Namespace: "Microsoft.DocumentDB",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DocumentDB)`,
		},
		{
			NewFunc:   armcosmos.NewLocationsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations",
			Namespace: "Microsoft.DocumentDB",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DocumentDB)`,
		},
		{
			NewFunc:   armcosmos.NewRestorableDatabaseAccountsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/restorableDatabaseAccounts",
			Namespace: "Microsoft.DocumentDB",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DocumentDB)`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcosmos())
}

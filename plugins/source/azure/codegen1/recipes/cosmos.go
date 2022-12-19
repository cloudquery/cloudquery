// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"

func Armcosmos() []*Table {
	tables := []*Table{
		{
			NewFunc:        armcosmos.NewLocationsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations",
			Namespace:      "microsoft.documentdb",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_documentdb)`,
			Pager:          `NewListPager`,
			ResponseStruct: "LocationsClientListResponse",
		},
		{
			NewFunc:        armcosmos.NewRestorableDatabaseAccountsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/restorableDatabaseAccounts",
			Namespace:      "microsoft.documentdb",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_documentdb)`,
			Pager:          `NewListPager`,
			ResponseStruct: "RestorableDatabaseAccountsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcosmos())
}

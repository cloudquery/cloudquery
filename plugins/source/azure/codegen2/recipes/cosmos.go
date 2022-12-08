// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"

func Armcosmos() []Table {
	tables := []Table{
		{
			Name:           "database_accounts",
			Struct:         &armcosmos.DatabaseAccountGetResults{},
			ResponseStruct: &armcosmos.DatabaseAccountsClientListResponse{},
			Client:         &armcosmos.DatabaseAccountsClient{},
			ListFunc:       (&armcosmos.DatabaseAccountsClient{}).NewListPager,
			NewFunc:        armcosmos.NewDatabaseAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/databaseAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DocumentDB)`,
		},
		{
			Name:           "locations",
			Struct:         &armcosmos.LocationGetResult{},
			ResponseStruct: &armcosmos.LocationsClientListResponse{},
			Client:         &armcosmos.LocationsClient{},
			ListFunc:       (&armcosmos.LocationsClient{}).NewListPager,
			NewFunc:        armcosmos.NewLocationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DocumentDB)`,
		},
		{
			Name:           "restorable_database_accounts",
			Struct:         &armcosmos.RestorableDatabaseAccountGetResult{},
			ResponseStruct: &armcosmos.RestorableDatabaseAccountsClientListResponse{},
			Client:         &armcosmos.RestorableDatabaseAccountsClient{},
			ListFunc:       (&armcosmos.RestorableDatabaseAccountsClient{}).NewListPager,
			NewFunc:        armcosmos.NewRestorableDatabaseAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/restorableDatabaseAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DocumentDB)`,
		},
	}

	for i := range tables {
		tables[i].Service = "armcosmos"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcosmos()...)
}

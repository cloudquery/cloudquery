// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"

func Armstorage() []Table {
	tables := []Table{
		{
			Service:        "armstorage",
			Name:           "accounts",
			Struct:         &armstorage.Account{},
			ResponseStruct: &armstorage.AccountsClientListResponse{},
			Client:         &armstorage.AccountsClient{},
			ListFunc:       (&armstorage.AccountsClient{}).NewListPager,
			NewFunc:        armstorage.NewAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/storageAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Storage)`,
		},
		{
			Service:        "armstorage",
			Name:           "deleted_accounts",
			Struct:         &armstorage.DeletedAccount{},
			ResponseStruct: &armstorage.DeletedAccountsClientListResponse{},
			Client:         &armstorage.DeletedAccountsClient{},
			ListFunc:       (&armstorage.DeletedAccountsClient{}).NewListPager,
			NewFunc:        armstorage.NewDeletedAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/deletedAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Storage)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armstorage()...)
}

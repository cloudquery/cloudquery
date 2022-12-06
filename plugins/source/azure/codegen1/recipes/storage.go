// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"

func Armstorage() []*Table {
	tables := []*Table{
		{
			NewFunc: armstorage.NewDeletedAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/deletedAccounts",
		},
		{
			NewFunc: armstorage.NewSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/skus",
		},
		{
			NewFunc: armstorage.NewAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/storageAccounts",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armstorage())
}
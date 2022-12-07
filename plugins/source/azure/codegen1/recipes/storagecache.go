// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache"

func Armstoragecache() []*Table {
	tables := []*Table{
		{
			NewFunc:   armstoragecache.NewCachesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/caches",
			Namespace: "Microsoft.StorageCache",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.StorageCache")`,
		},
		{
			NewFunc:   armstoragecache.NewSKUsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/skus",
			Namespace: "Microsoft.StorageCache",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.StorageCache")`,
		},
		{
			NewFunc:   armstoragecache.NewUsageModelsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/usageModels",
			Namespace: "Microsoft.StorageCache",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.StorageCache")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armstoragecache())
}

// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache"

func init() {
	tables := []Table{
		{
			Service:        "armstoragecache",
			Name:           "caches",
			Struct:         &armstoragecache.Cache{},
			ResponseStruct: &armstoragecache.CachesClientListResponse{},
			Client:         &armstoragecache.CachesClient{},
			ListFunc:       (&armstoragecache.CachesClient{}).NewListPager,
			NewFunc:        armstoragecache.NewCachesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/caches",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_StorageCache)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}

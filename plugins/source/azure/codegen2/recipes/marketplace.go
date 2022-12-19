// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"

func init() {
	tables := []Table{
		{
			Service:        "armmarketplace",
			Name:           "private_store",
			Struct:         &armmarketplace.PrivateStore{},
			ResponseStruct: &armmarketplace.PrivateStoreClientListResponse{},
			Client:         &armmarketplace.PrivateStoreClient{},
			ListFunc:       (&armmarketplace.PrivateStoreClient{}).NewListPager,
			NewFunc:        armmarketplace.NewPrivateStoreClient,
			URL:            "/providers/Microsoft.Marketplace/privateStores",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_marketplace)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}

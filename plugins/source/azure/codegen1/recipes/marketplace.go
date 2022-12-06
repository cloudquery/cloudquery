// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"

func Armmarketplace() []*Table {
	tables := []*Table{
		{
			NewFunc:   armmarketplace.NewPrivateStoreClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace",
			URL:       "/providers/Microsoft.Marketplace/privateStores",
			Namespace: "Microsoft.Marketplace",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Marketplace")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmarketplace())
}

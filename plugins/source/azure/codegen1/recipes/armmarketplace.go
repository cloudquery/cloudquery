// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"

func Armmarketplace() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmarketplace.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace",
		},
		{
			NewFunc: armmarketplace.NewPrivateStoreCollectionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace",
		},
		{
			NewFunc: armmarketplace.NewPrivateStoreClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace",
		},
		{
			NewFunc: armmarketplace.NewPrivateStoreCollectionOfferClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmarketplace())
}
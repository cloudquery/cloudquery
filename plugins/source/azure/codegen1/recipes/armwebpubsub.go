// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"

func Armwebpubsub() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armwebpubsub.NewHubsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub",
		},
		{
			NewFunc: armwebpubsub.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub",
		},
		{
			NewFunc: armwebpubsub.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub",
		},
		{
			NewFunc: armwebpubsub.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub",
		},
		{
			NewFunc: armwebpubsub.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub",
		},
		{
			NewFunc: armwebpubsub.NewSharedPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub",
		},
		{
			NewFunc: armwebpubsub.NewUsagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armwebpubsub())
}
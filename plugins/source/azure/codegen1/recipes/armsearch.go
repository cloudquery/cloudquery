// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"

func Armsearch() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armsearch.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch",
		},
		{
			NewFunc: armsearch.NewSharedPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch",
		},
		{
			NewFunc: armsearch.NewAdminKeysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch",
		},
		{
			NewFunc: armsearch.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch",
		},
		{
			NewFunc: armsearch.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch",
		},
		{
			NewFunc: armsearch.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch",
		},
		{
			NewFunc: armsearch.NewQueryKeysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armsearch())
}
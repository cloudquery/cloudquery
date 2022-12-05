// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"

func Armagrifood() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armagrifood.NewExtensionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood",
			URL: "",
		},
		{
			NewFunc: armagrifood.NewFarmBeatsExtensionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood",
			URL: "/providers/Microsoft.AgFoodPlatform/farmBeatsExtensionDefinitions",
		},
		{
			NewFunc: armagrifood.NewFarmBeatsModelsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood",
			URL: "",
		},
		{
			NewFunc: armagrifood.NewLocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood",
			URL: "",
		},
		{
			NewFunc: armagrifood.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood",
			URL: "",
		},
		{
			NewFunc: armagrifood.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armagrifood())
}
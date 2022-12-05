// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks"

func Armpowerbiprivatelinks() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armpowerbiprivatelinks.NewPowerBIResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks",
			URL: "",
		},
		{
			NewFunc: armpowerbiprivatelinks.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks",
			URL: "",
		},
		{
			NewFunc: armpowerbiprivatelinks.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks",
			URL: "",
		},
		{
			NewFunc: armpowerbiprivatelinks.NewPrivateLinkServiceResourceOperationResultsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks",
			URL: "",
		},
		{
			NewFunc: armpowerbiprivatelinks.NewPrivateLinkServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks",
			URL: "",
		},
		{
			NewFunc: armpowerbiprivatelinks.NewPrivateLinkServicesForPowerBIClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armpowerbiprivatelinks())
}
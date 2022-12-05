// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"

func Armdigitaltwins() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdigitaltwins.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DigitalTwins/digitalTwinsInstances",
		},
		{
			NewFunc: armdigitaltwins.NewEndpointClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints",
		},
		{
			NewFunc: armdigitaltwins.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/privateEndpointConnections",
		},
		{
			NewFunc: armdigitaltwins.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/privateLinkResources",
		},
		{
			NewFunc: armdigitaltwins.NewTimeSeriesDatabaseConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/timeSeriesDatabaseConnections",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdigitaltwins())
}
// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"

func Armtrafficmanager() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armtrafficmanager.NewHeatMapClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager",
		},
		{
			NewFunc: armtrafficmanager.NewProfilesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager",
		},
		{
			NewFunc: armtrafficmanager.NewEndpointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager",
		},
		{
			NewFunc: armtrafficmanager.NewUserMetricsKeysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager",
		},
		{
			NewFunc: armtrafficmanager.NewGeographicHierarchiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armtrafficmanager())
}
// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"

func Armapplicationinsights() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armapplicationinsights.NewAPIKeysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ApiKeys",
		},
		{
			NewFunc: armapplicationinsights.NewAnalyticsItemsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}",
		},
		{
			NewFunc: armapplicationinsights.NewAnnotationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations",
		},
		{
			NewFunc: armapplicationinsights.NewComponentAvailableFeaturesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "",
		},
		{
			NewFunc: armapplicationinsights.NewComponentCurrentBillingFeaturesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "",
		},
		{
			NewFunc: armapplicationinsights.NewComponentFeatureCapabilitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "",
		},
		{
			NewFunc: armapplicationinsights.NewComponentQuotaStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "",
		},
		{
			NewFunc: armapplicationinsights.NewComponentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/components",
		},
		{
			NewFunc: armapplicationinsights.NewExportConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration",
		},
		{
			NewFunc: armapplicationinsights.NewFavoritesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites",
		},
		{
			NewFunc: armapplicationinsights.NewMyWorkbooksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "",
		},
		{
			NewFunc: armapplicationinsights.NewProactiveDetectionConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs",
		},
		{
			NewFunc: armapplicationinsights.NewWebTestLocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/syntheticmonitorlocations",
		},
		{
			NewFunc: armapplicationinsights.NewWebTestsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/webtests",
		},
		{
			NewFunc: armapplicationinsights.NewWorkItemConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs",
		},
		{
			NewFunc: armapplicationinsights.NewWorkbooksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armapplicationinsights())
}
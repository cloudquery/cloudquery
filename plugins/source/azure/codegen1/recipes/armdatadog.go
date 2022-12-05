// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"

func Armdatadog() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdatadog.NewMarketplaceAgreementsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/agreements",
		},
		{
			NewFunc: armdatadog.NewMonitorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/monitors",
		},
		{
			NewFunc: armdatadog.NewSingleSignOnConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/singleSignOnConfigurations",
		},
		{
			NewFunc: armdatadog.NewTagRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/tagRules",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdatadog())
}
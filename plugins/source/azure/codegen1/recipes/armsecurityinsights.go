// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"

func Armsecurityinsights() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armsecurityinsights.NewWatchlistsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewDataConnectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewIncidentRelationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewSentinelOnboardingStatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewThreatIntelligenceIndicatorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewAlertRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewActionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewAlertRuleTemplatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewIncidentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewThreatIntelligenceIndicatorClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewWatchlistItemsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewAutomationRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewBookmarksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewIncidentCommentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
		{
			NewFunc: armsecurityinsights.NewThreatIntelligenceIndicatorMetricsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armsecurityinsights())
}
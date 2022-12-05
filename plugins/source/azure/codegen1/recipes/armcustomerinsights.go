// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"

func Armcustomerinsights() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcustomerinsights.NewAuthorizationPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewConnectorMappingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewConnectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewHubsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.CustomerInsights/hubs",
		},
		{
			NewFunc: armcustomerinsights.NewImagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewInteractionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewKpiClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewPredictionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewProfilesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewRelationshipLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewRelationshipsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewRoleAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewRolesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewViewsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
		{
			NewFunc: armcustomerinsights.NewWidgetTypesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armcustomerinsights())
}
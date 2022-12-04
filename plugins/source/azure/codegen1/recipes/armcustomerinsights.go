// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"

func Armcustomerinsights() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcustomerinsights.NewAuthorizationPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewInteractionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewProfilesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewRoleAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewWidgetTypesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewConnectorMappingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewConnectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewHubsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewRolesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewImagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewKpiClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewPredictionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewRelationshipLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewRelationshipsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
		{
			NewFunc: armcustomerinsights.NewViewsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armcustomerinsights())
}
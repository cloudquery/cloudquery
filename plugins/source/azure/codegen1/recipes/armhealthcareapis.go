// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"

func Armhealthcareapis() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armhealthcareapis.NewDicomServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
		{
			NewFunc: armhealthcareapis.NewFhirDestinationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
		{
			NewFunc: armhealthcareapis.NewFhirServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
		{
			NewFunc: armhealthcareapis.NewIotConnectorFhirDestinationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
		{
			NewFunc: armhealthcareapis.NewIotConnectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
		{
			NewFunc: armhealthcareapis.NewOperationResultsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
		{
			NewFunc: armhealthcareapis.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
		{
			NewFunc: armhealthcareapis.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
		{
			NewFunc: armhealthcareapis.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.HealthcareApis/services",
		},
		{
			NewFunc: armhealthcareapis.NewWorkspacePrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
		{
			NewFunc: armhealthcareapis.NewWorkspacePrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
		{
			NewFunc: armhealthcareapis.NewWorkspacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armhealthcareapis())
}
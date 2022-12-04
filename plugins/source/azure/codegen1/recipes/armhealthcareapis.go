// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"

func Armhealthcareapis() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armhealthcareapis.NewFhirDestinationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewDicomServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewFhirServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewIotConnectorFhirDestinationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewWorkspacePrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewOperationResultsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewIotConnectorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewWorkspacePrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
		{
			NewFunc: armhealthcareapis.NewWorkspacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armhealthcareapis())
}
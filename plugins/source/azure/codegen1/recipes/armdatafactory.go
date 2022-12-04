// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"

func Armdatafactory() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdatafactory.NewDataFlowsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewManagedPrivateEndpointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewPipelineRunsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewPrivateEndPointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewExposureControlClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewLinkedServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewActivityRunsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewPrivateEndpointConnectionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewTriggerRunsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewIntegrationRuntimesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewPipelinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewFactoriesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewGlobalParametersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewIntegrationRuntimeObjectMetadataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewDataFlowDebugSessionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewDatasetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewTriggersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewIntegrationRuntimeNodesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewManagedVirtualNetworksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
		{
			NewFunc: armdatafactory.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdatafactory())
}
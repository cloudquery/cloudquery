// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning"

func Armmachinelearning() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmachinelearning.NewComputeClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes",
		},
		{
			NewFunc: armmachinelearning.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/privateEndpointConnections",
		},
		{
			NewFunc: armmachinelearning.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/privateLinkResources",
		},
		{
			NewFunc: armmachinelearning.NewQuotasClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/quotas",
		},
		{
			NewFunc: armmachinelearning.NewUsagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/usages",
		},
		{
			NewFunc: armmachinelearning.NewVirtualMachineSizesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/vmSizes",
		},
		{
			NewFunc: armmachinelearning.NewWorkspaceConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections",
		},
		{
			NewFunc: armmachinelearning.NewWorkspaceFeaturesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/features",
		},
		{
			NewFunc: armmachinelearning.NewWorkspaceSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/workspaces/skus",
		},
		{
			NewFunc: armmachinelearning.NewWorkspacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmachinelearning())
}
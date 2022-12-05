// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"

func Armkusto() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armkusto.NewAttachedDatabaseConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "",
		},
		{
			NewFunc: armkusto.NewClusterPrincipalAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/principalAssignments",
		},
		{
			NewFunc: armkusto.NewClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/clusters",
		},
		{
			NewFunc: armkusto.NewDataConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "",
		},
		{
			NewFunc: armkusto.NewDatabasePrincipalAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/principalAssignments",
		},
		{
			NewFunc: armkusto.NewDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "",
		},
		{
			NewFunc: armkusto.NewManagedPrivateEndpointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/managedPrivateEndpoints",
		},
		{
			NewFunc: armkusto.NewOperationsResultsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "",
		},
		{
			NewFunc: armkusto.NewOperationsResultsLocationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "",
		},
		{
			NewFunc: armkusto.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/privateEndpointConnections",
		},
		{
			NewFunc: armkusto.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/privateLinkResources",
		},
		{
			NewFunc: armkusto.NewScriptsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armkusto())
}
// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"

func Armdatabricks() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdatabricks.NewWorkspacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks",
		},
		{
			NewFunc: armdatabricks.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks",
		},
		{
			NewFunc: armdatabricks.NewOutboundNetworkDependenciesEndpointsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks",
		},
		{
			NewFunc: armdatabricks.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks",
		},
		{
			NewFunc: armdatabricks.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks",
		},
		{
			NewFunc: armdatabricks.NewVNetPeeringClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdatabricks())
}
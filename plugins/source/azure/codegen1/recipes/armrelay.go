// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"

func Armrelay() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armrelay.NewHybridConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
			URL: "",
		},
		{
			NewFunc: armrelay.NewNamespacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Relay/namespaces",
		},
		{
			NewFunc: armrelay.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/privateEndpointConnections",
		},
		{
			NewFunc: armrelay.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/privateLinkResources",
		},
		{
			NewFunc: armrelay.NewWCFRelaysClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armrelay())
}
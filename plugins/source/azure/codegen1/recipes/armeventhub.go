// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"

func Armeventhub() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armeventhub.NewDisasterRecoveryConfigsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
		},
		{
			NewFunc: armeventhub.NewConfigurationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
		},
		{
			NewFunc: armeventhub.NewNamespacesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
		},
		{
			NewFunc: armeventhub.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
		},
		{
			NewFunc: armeventhub.NewConsumerGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
		},
		{
			NewFunc: armeventhub.NewClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
		},
		{
			NewFunc: armeventhub.NewEventHubsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
		},
		{
			NewFunc: armeventhub.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
		},
		{
			NewFunc: armeventhub.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
		},
		{
			NewFunc: armeventhub.NewSchemaRegistryClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armeventhub())
}
// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"

func Armhdinsight() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armhdinsight.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
		{
			NewFunc: armhdinsight.NewVirtualMachinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
		{
			NewFunc: armhdinsight.NewScriptActionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
		{
			NewFunc: armhdinsight.NewApplicationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
		{
			NewFunc: armhdinsight.NewClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
		{
			NewFunc: armhdinsight.NewConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
		{
			NewFunc: armhdinsight.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
		{
			NewFunc: armhdinsight.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
		{
			NewFunc: armhdinsight.NewScriptExecutionHistoryClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
		{
			NewFunc: armhdinsight.NewExtensionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
		{
			NewFunc: armhdinsight.NewLocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armhdinsight())
}
// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"

func Armavs() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armavs.NewAddonsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewCloudLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewPrivateCloudsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewScriptPackagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewWorkloadNetworksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewVirtualMachinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewPlacementPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewAuthorizationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewDatastoresClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewGlobalReachConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewScriptCmdletsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewScriptExecutionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewHcxEnterpriseSitesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
		{
			NewFunc: armavs.NewLocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armavs())
}
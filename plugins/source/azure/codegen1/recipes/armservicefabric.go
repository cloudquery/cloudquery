// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"

func Armservicefabric() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armservicefabric.NewApplicationTypesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric",
		},
		{
			NewFunc: armservicefabric.NewApplicationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric",
		},
		{
			NewFunc: armservicefabric.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric",
		},
		{
			NewFunc: armservicefabric.NewApplicationTypeVersionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric",
		},
		{
			NewFunc: armservicefabric.NewClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric",
		},
		{
			NewFunc: armservicefabric.NewClusterVersionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric",
		},
		{
			NewFunc: armservicefabric.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armservicefabric())
}
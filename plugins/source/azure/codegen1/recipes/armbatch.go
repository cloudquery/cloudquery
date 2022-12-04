// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"

func Armbatch() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armbatch.NewApplicationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
		},
		{
			NewFunc: armbatch.NewCertificateClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
		},
		{
			NewFunc: armbatch.NewPoolClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
		},
		{
			NewFunc: armbatch.NewLocationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
		},
		{
			NewFunc: armbatch.NewAccountClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
		},
		{
			NewFunc: armbatch.NewApplicationPackageClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
		},
		{
			NewFunc: armbatch.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
		},
		{
			NewFunc: armbatch.NewPrivateEndpointConnectionClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
		},
		{
			NewFunc: armbatch.NewPrivateLinkResourceClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armbatch())
}
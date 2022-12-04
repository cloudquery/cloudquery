// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"

func Armdevcenter() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdevcenter.NewDevCentersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewEnvironmentTypesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewGalleriesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewDevBoxDefinitionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewImageVersionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewProjectsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewUsagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewImagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewPoolsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewProjectAllowedEnvironmentTypesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewCatalogsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewNetworkConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewOperationStatusesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewProjectEnvironmentTypesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewSchedulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
		{
			NewFunc: armdevcenter.NewAttachedNetworksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdevcenter())
}
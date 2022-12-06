// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"

func Armcompute() []*Table {
	tables := []*Table{
		{
			NewFunc: armcompute.NewCloudServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices",
		},
		{
			NewFunc: armcompute.NewAvailabilitySetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets",
		},
		{
			NewFunc: armcompute.NewDisksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks",
		},
		{
			NewFunc: armcompute.NewRestorePointCollectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections",
		},
		{
			NewFunc: armcompute.NewSnapshotsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots",
		},
		{
			NewFunc: armcompute.NewDiskEncryptionSetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskEncryptionSets",
		},
		{
			NewFunc: armcompute.NewGalleriesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/galleries",
		},
		{
			NewFunc: armcompute.NewVirtualMachineScaleSetsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets",
		},
		{
			NewFunc: armcompute.NewVirtualMachinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines",
		},
		{
			NewFunc: armcompute.NewImagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/images",
		},
		{
			NewFunc: armcompute.NewDiskAccessesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskAccesses",
		},
		{
			NewFunc: armcompute.NewResourceSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/skus",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcompute())
}

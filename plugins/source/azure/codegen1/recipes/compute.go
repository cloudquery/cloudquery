// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"

func Armcompute() []*Table {
	tables := []*Table{
		{
			NewFunc:   armcompute.NewImagesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/images",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewVirtualMachinesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionResourceGroupMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewVirtualMachineScaleSetsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionResourceGroupMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewSnapshotsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewCloudServicesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionResourceGroupMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewAvailabilitySetsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionResourceGroupMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewDisksClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewGalleriesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/galleries",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewDiskAccessesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskAccesses",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewDiskEncryptionSetsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskEncryptionSets",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewResourceSKUsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/skus",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
		{
			NewFunc:   armcompute.NewRestorePointCollectionsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute",
			URL:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections",
			Namespace: "Microsoft.Compute",
			Multiplex: `client.SubscriptionResourceGroupMultiplexRegisteredNamespace("Microsoft.Compute")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcompute())
}

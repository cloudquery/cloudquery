// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"

func Armconnectedvmware() []*Table {
	tables := []*Table{
		{
			NewFunc: armconnectedvmware.NewVirtualMachinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines",
		},
		{
			NewFunc: armconnectedvmware.NewClustersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/clusters",
		},
		{
			NewFunc: armconnectedvmware.NewVirtualNetworksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks",
		},
		{
			NewFunc: armconnectedvmware.NewVirtualMachineTemplatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates",
		},
		{
			NewFunc: armconnectedvmware.NewDatastoresClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/datastores",
		},
		{
			NewFunc: armconnectedvmware.NewHostsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/hosts",
		},
		{
			NewFunc: armconnectedvmware.NewResourcePoolsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/resourcePools",
		},
		{
			NewFunc: armconnectedvmware.NewVCentersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/vcenters",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armconnectedvmware())
}

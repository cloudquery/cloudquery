// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"

func Armconnectedvmware() []*Table {
	tables := []*Table{
		{
			NewFunc:   armconnectedvmware.NewVirtualMachinesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines",
			Namespace: "Microsoft.ConnectedVMwarevSphere",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.ConnectedVMwarevSphere")`,
		},
		{
			NewFunc:   armconnectedvmware.NewHostsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/hosts",
			Namespace: "Microsoft.ConnectedVMwarevSphere",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.ConnectedVMwarevSphere")`,
		},
		{
			NewFunc:   armconnectedvmware.NewDatastoresClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/datastores",
			Namespace: "Microsoft.ConnectedVMwarevSphere",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.ConnectedVMwarevSphere")`,
		},
		{
			NewFunc:   armconnectedvmware.NewVirtualNetworksClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks",
			Namespace: "Microsoft.ConnectedVMwarevSphere",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.ConnectedVMwarevSphere")`,
		},
		{
			NewFunc:   armconnectedvmware.NewVirtualMachineTemplatesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates",
			Namespace: "Microsoft.ConnectedVMwarevSphere",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.ConnectedVMwarevSphere")`,
		},
		{
			NewFunc:   armconnectedvmware.NewClustersClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/clusters",
			Namespace: "Microsoft.ConnectedVMwarevSphere",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.ConnectedVMwarevSphere")`,
		},
		{
			NewFunc:   armconnectedvmware.NewVCentersClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/vcenters",
			Namespace: "Microsoft.ConnectedVMwarevSphere",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.ConnectedVMwarevSphere")`,
		},
		{
			NewFunc:   armconnectedvmware.NewResourcePoolsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/resourcePools",
			Namespace: "Microsoft.ConnectedVMwarevSphere",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.ConnectedVMwarevSphere")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armconnectedvmware())
}

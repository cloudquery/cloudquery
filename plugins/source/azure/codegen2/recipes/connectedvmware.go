// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"

func init() {
	tables := []Table{
		{
			Service:        "armconnectedvmware",
			Name:           "clusters",
			Struct:         &armconnectedvmware.Cluster{},
			ResponseStruct: &armconnectedvmware.ClustersClientListResponse{},
			Client:         &armconnectedvmware.ClustersClient{},
			ListFunc:       (&armconnectedvmware.ClustersClient{}).NewListPager,
			NewFunc:        armconnectedvmware.NewClustersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/clusters",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ConnectedVMwarevSphere)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armconnectedvmware",
			Name:           "datastores",
			Struct:         &armconnectedvmware.Datastore{},
			ResponseStruct: &armconnectedvmware.DatastoresClientListResponse{},
			Client:         &armconnectedvmware.DatastoresClient{},
			ListFunc:       (&armconnectedvmware.DatastoresClient{}).NewListPager,
			NewFunc:        armconnectedvmware.NewDatastoresClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/datastores",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ConnectedVMwarevSphere)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armconnectedvmware",
			Name:           "hosts",
			Struct:         &armconnectedvmware.Host{},
			ResponseStruct: &armconnectedvmware.HostsClientListResponse{},
			Client:         &armconnectedvmware.HostsClient{},
			ListFunc:       (&armconnectedvmware.HostsClient{}).NewListPager,
			NewFunc:        armconnectedvmware.NewHostsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/hosts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ConnectedVMwarevSphere)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armconnectedvmware",
			Name:           "resource_pools",
			Struct:         &armconnectedvmware.ResourcePool{},
			ResponseStruct: &armconnectedvmware.ResourcePoolsClientListResponse{},
			Client:         &armconnectedvmware.ResourcePoolsClient{},
			ListFunc:       (&armconnectedvmware.ResourcePoolsClient{}).NewListPager,
			NewFunc:        armconnectedvmware.NewResourcePoolsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/resourcePools",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ConnectedVMwarevSphere)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armconnectedvmware",
			Name:           "v_centers",
			Struct:         &armconnectedvmware.VCenter{},
			ResponseStruct: &armconnectedvmware.VCentersClientListResponse{},
			Client:         &armconnectedvmware.VCentersClient{},
			ListFunc:       (&armconnectedvmware.VCentersClient{}).NewListPager,
			NewFunc:        armconnectedvmware.NewVCentersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/vcenters",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ConnectedVMwarevSphere)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armconnectedvmware",
			Name:           "virtual_machine_templates",
			Struct:         &armconnectedvmware.VirtualMachineTemplate{},
			ResponseStruct: &armconnectedvmware.VirtualMachineTemplatesClientListResponse{},
			Client:         &armconnectedvmware.VirtualMachineTemplatesClient{},
			ListFunc:       (&armconnectedvmware.VirtualMachineTemplatesClient{}).NewListPager,
			NewFunc:        armconnectedvmware.NewVirtualMachineTemplatesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ConnectedVMwarevSphere)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armconnectedvmware",
			Name:           "virtual_machines",
			Struct:         &armconnectedvmware.VirtualMachine{},
			ResponseStruct: &armconnectedvmware.VirtualMachinesClientListResponse{},
			Client:         &armconnectedvmware.VirtualMachinesClient{},
			ListFunc:       (&armconnectedvmware.VirtualMachinesClient{}).NewListPager,
			NewFunc:        armconnectedvmware.NewVirtualMachinesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ConnectedVMwarevSphere)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armconnectedvmware",
			Name:           "virtual_networks",
			Struct:         &armconnectedvmware.VirtualNetwork{},
			ResponseStruct: &armconnectedvmware.VirtualNetworksClientListResponse{},
			Client:         &armconnectedvmware.VirtualNetworksClient{},
			ListFunc:       (&armconnectedvmware.VirtualNetworksClient{}).NewListPager,
			NewFunc:        armconnectedvmware.NewVirtualNetworksClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ConnectedVMwarevSphere)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}

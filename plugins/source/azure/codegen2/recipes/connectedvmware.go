// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"

func Armconnectedvmware() []Table {
	tables := []Table{
		{
      Name: "v_center",
      Struct: &armconnectedvmware.VCenter{},
      ResponseStruct: &armconnectedvmware.VCentersClientListResponse{},
      Client: &armconnectedvmware.VCentersClient{},
      ListFunc: (&armconnectedvmware.VCentersClient{}).NewListPager,
			NewFunc: armconnectedvmware.NewVCentersClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/vcenters",
		},
		{
      Name: "virtual_network",
      Struct: &armconnectedvmware.VirtualNetwork{},
      ResponseStruct: &armconnectedvmware.VirtualNetworksClientListResponse{},
      Client: &armconnectedvmware.VirtualNetworksClient{},
      ListFunc: (&armconnectedvmware.VirtualNetworksClient{}).NewListPager,
			NewFunc: armconnectedvmware.NewVirtualNetworksClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks",
		},
		{
      Name: "host",
      Struct: &armconnectedvmware.Host{},
      ResponseStruct: &armconnectedvmware.HostsClientListResponse{},
      Client: &armconnectedvmware.HostsClient{},
      ListFunc: (&armconnectedvmware.HostsClient{}).NewListPager,
			NewFunc: armconnectedvmware.NewHostsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/hosts",
		},
		{
      Name: "virtual_machine",
      Struct: &armconnectedvmware.VirtualMachine{},
      ResponseStruct: &armconnectedvmware.VirtualMachinesClientListResponse{},
      Client: &armconnectedvmware.VirtualMachinesClient{},
      ListFunc: (&armconnectedvmware.VirtualMachinesClient{}).NewListPager,
			NewFunc: armconnectedvmware.NewVirtualMachinesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines",
		},
		{
      Name: "datastore",
      Struct: &armconnectedvmware.Datastore{},
      ResponseStruct: &armconnectedvmware.DatastoresClientListResponse{},
      Client: &armconnectedvmware.DatastoresClient{},
      ListFunc: (&armconnectedvmware.DatastoresClient{}).NewListPager,
			NewFunc: armconnectedvmware.NewDatastoresClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/datastores",
		},
		{
      Name: "cluster",
      Struct: &armconnectedvmware.Cluster{},
      ResponseStruct: &armconnectedvmware.ClustersClientListResponse{},
      Client: &armconnectedvmware.ClustersClient{},
      ListFunc: (&armconnectedvmware.ClustersClient{}).NewListPager,
			NewFunc: armconnectedvmware.NewClustersClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/clusters",
		},
		{
      Name: "virtual_machine_template",
      Struct: &armconnectedvmware.VirtualMachineTemplate{},
      ResponseStruct: &armconnectedvmware.VirtualMachineTemplatesClientListResponse{},
      Client: &armconnectedvmware.VirtualMachineTemplatesClient{},
      ListFunc: (&armconnectedvmware.VirtualMachineTemplatesClient{}).NewListPager,
			NewFunc: armconnectedvmware.NewVirtualMachineTemplatesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates",
		},
		{
      Name: "resource_pool",
      Struct: &armconnectedvmware.ResourcePool{},
      ResponseStruct: &armconnectedvmware.ResourcePoolsClientListResponse{},
      Client: &armconnectedvmware.ResourcePoolsClient{},
      ListFunc: (&armconnectedvmware.ResourcePoolsClient{}).NewListPager,
			NewFunc: armconnectedvmware.NewResourcePoolsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/resourcePools",
		},
	}

	for i := range tables {
		tables[i].Service = "armconnectedvmware"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armconnectedvmware()...)
}
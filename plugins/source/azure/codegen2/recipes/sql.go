// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"

func Armsql() []Table {
	tables := []Table{
		{
      Name: "virtual_cluster",
      Struct: &armsql.VirtualCluster{},
      ResponseStruct: &armsql.VirtualClustersClientListResponse{},
      Client: &armsql.VirtualClustersClient{},
      ListFunc: (&armsql.VirtualClustersClient{}).NewListPager,
			NewFunc: armsql.NewVirtualClustersClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/virtualClusters",
		},
		{
      Name: "managed_instance",
      Struct: &armsql.ManagedInstance{},
      ResponseStruct: &armsql.ManagedInstancesClientListResponse{},
      Client: &armsql.ManagedInstancesClient{},
      ListFunc: (&armsql.ManagedInstancesClient{}).NewListPager,
			NewFunc: armsql.NewManagedInstancesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/managedInstances",
		},
		{
      Name: "server",
      Struct: &armsql.Server{},
      ResponseStruct: &armsql.ServersClientListResponse{},
      Client: &armsql.ServersClient{},
      ListFunc: (&armsql.ServersClient{}).NewListPager,
			NewFunc: armsql.NewServersClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/servers",
		},
		{
      Name: "instance_pool",
      Struct: &armsql.InstancePool{},
      ResponseStruct: &armsql.InstancePoolsClientListResponse{},
      Client: &armsql.InstancePoolsClient{},
      ListFunc: (&armsql.InstancePoolsClient{}).NewListPager,
			NewFunc: armsql.NewInstancePoolsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/instancePools",
		},
		{
      Name: "deleted_server",
      Struct: &armsql.DeletedServer{},
      ResponseStruct: &armsql.DeletedServersClientListResponse{},
      Client: &armsql.DeletedServersClient{},
      ListFunc: (&armsql.DeletedServersClient{}).NewListPager,
			NewFunc: armsql.NewDeletedServersClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/deletedServers",
		},
	}

	for i := range tables {
		tables[i].Service = "armsql"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armsql()...)
}
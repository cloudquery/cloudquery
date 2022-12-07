// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"

func Armsql() []*Table {
	tables := []*Table{
		{
			NewFunc:   armsql.NewServersClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/servers",
			Namespace: "Microsoft.Sql",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Sql")`,
		},
		{
			NewFunc:   armsql.NewManagedInstancesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/managedInstances",
			Namespace: "Microsoft.Sql",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Sql")`,
		},
		{
			NewFunc:   armsql.NewInstancePoolsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/instancePools",
			Namespace: "Microsoft.Sql",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Sql")`,
		},
		{
			NewFunc:   armsql.NewVirtualClustersClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/virtualClusters",
			Namespace: "Microsoft.Sql",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Sql")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armsql())
}

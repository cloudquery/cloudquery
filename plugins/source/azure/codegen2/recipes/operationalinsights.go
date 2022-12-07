// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"

func Armoperationalinsights() []Table {
	tables := []Table{
		{
			Name:           "workspaces",
			Struct:         &armoperationalinsights.Workspace{},
			ResponseStruct: &armoperationalinsights.WorkspacesClientListResponse{},
			Client:         &armoperationalinsights.WorkspacesClient{},
			ListFunc:       (&armoperationalinsights.WorkspacesClient{}).NewListPager,
			NewFunc:        armoperationalinsights.NewWorkspacesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/workspaces",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.OperationalInsights")`,
		},
		{
			Name:           "clusters",
			Struct:         &armoperationalinsights.Cluster{},
			ResponseStruct: &armoperationalinsights.ClustersClientListResponse{},
			Client:         &armoperationalinsights.ClustersClient{},
			ListFunc:       (&armoperationalinsights.ClustersClient{}).NewListPager,
			NewFunc:        armoperationalinsights.NewClustersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/clusters",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.OperationalInsights")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armoperationalinsights"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armoperationalinsights()...)
}

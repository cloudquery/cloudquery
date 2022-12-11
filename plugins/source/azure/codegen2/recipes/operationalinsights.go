// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"

func init() {
	tables := []Table{
		{
			Service:        "armoperationalinsights",
			Name:           "clusters",
			Struct:         &armoperationalinsights.Cluster{},
			ResponseStruct: &armoperationalinsights.ClustersClientListResponse{},
			Client:         &armoperationalinsights.ClustersClient{},
			ListFunc:       (&armoperationalinsights.ClustersClient{}).NewListPager,
			NewFunc:        armoperationalinsights.NewClustersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/clusters",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_OperationalInsights)`,
		},
		{
			Service:        "armoperationalinsights",
			Name:           "workspaces",
			Struct:         &armoperationalinsights.Workspace{},
			ResponseStruct: &armoperationalinsights.WorkspacesClientListResponse{},
			Client:         &armoperationalinsights.WorkspacesClient{},
			ListFunc:       (&armoperationalinsights.WorkspacesClient{}).NewListPager,
			NewFunc:        armoperationalinsights.NewWorkspacesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/workspaces",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_OperationalInsights)`,
		},
	}
	Tables = append(Tables, tables...)
}

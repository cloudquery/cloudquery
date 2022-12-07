// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"

func Armsynapse() []Table {
	tables := []Table{
		{
			Name:           "private_link_hubs",
			Struct:         &armsynapse.PrivateLinkHub{},
			ResponseStruct: &armsynapse.PrivateLinkHubsClientListResponse{},
			Client:         &armsynapse.PrivateLinkHubsClient{},
			ListFunc:       (&armsynapse.PrivateLinkHubsClient{}).NewListPager,
			NewFunc:        armsynapse.NewPrivateLinkHubsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/privateLinkHubs",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Synapse")`,
		},
		{
			Name:           "workspaces",
			Struct:         &armsynapse.Workspace{},
			ResponseStruct: &armsynapse.WorkspacesClientListResponse{},
			Client:         &armsynapse.WorkspacesClient{},
			ListFunc:       (&armsynapse.WorkspacesClient{}).NewListPager,
			NewFunc:        armsynapse.NewWorkspacesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/workspaces",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Synapse")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armsynapse"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armsynapse()...)
}

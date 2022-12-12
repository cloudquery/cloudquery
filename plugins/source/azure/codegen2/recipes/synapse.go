// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"

func init() {
	tables := []Table{
		{
			Service:        "armsynapse",
			Name:           "private_link_hubs",
			Struct:         &armsynapse.PrivateLinkHub{},
			ResponseStruct: &armsynapse.PrivateLinkHubsClientListResponse{},
			Client:         &armsynapse.PrivateLinkHubsClient{},
			ListFunc:       (&armsynapse.PrivateLinkHubsClient{}).NewListPager,
			NewFunc:        armsynapse.NewPrivateLinkHubsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/privateLinkHubs",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Synapse)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsynapse",
			Name:           "workspaces",
			Struct:         &armsynapse.Workspace{},
			ResponseStruct: &armsynapse.WorkspacesClientListResponse{},
			Client:         &armsynapse.WorkspacesClient{},
			ListFunc:       (&armsynapse.WorkspacesClient{}).NewListPager,
			NewFunc:        armsynapse.NewWorkspacesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/workspaces",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Synapse)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}

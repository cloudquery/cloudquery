// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

func Armmonitor() []Table {
	tables := []Table{
		{
			Service:        "armmonitor",
			Name:           "log_profiles",
			Struct:         &armmonitor.LogProfileResource{},
			ResponseStruct: &armmonitor.LogProfilesClientListResponse{},
			Client:         &armmonitor.LogProfilesClient{},
			ListFunc:       (&armmonitor.LogProfilesClient{}).NewListPager,
			NewFunc:        armmonitor.NewLogProfilesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Insights)`,
		},
		{
			Service:        "armmonitor",
			Name:           "private_link_scopes",
			Struct:         &armmonitor.AzureMonitorPrivateLinkScope{},
			ResponseStruct: &armmonitor.PrivateLinkScopesClientListResponse{},
			Client:         &armmonitor.PrivateLinkScopesClient{},
			ListFunc:       (&armmonitor.PrivateLinkScopesClient{}).NewListPager,
			NewFunc:        armmonitor.NewPrivateLinkScopesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/microsoft.insights/privateLinkScopes",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_insights)`,
		},
		{
			Service:        "armmonitor",
			Name:           "tenant_activity_logs",
			Struct:         &armmonitor.EventData{},
			ResponseStruct: &armmonitor.TenantActivityLogsClientListResponse{},
			Client:         &armmonitor.TenantActivityLogsClient{},
			ListFunc:       (&armmonitor.TenantActivityLogsClient{}).NewListPager,
			NewFunc:        armmonitor.NewTenantActivityLogsClient,
			URL:            "/providers/Microsoft.Insights/eventtypes/management/values",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Insights)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armmonitor()...)
}

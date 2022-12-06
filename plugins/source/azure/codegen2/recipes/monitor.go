// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

func Armmonitor() []Table {
	tables := []Table{
		{
			Name:           "azure_monitor_private_link_scope",
			Struct:         &armmonitor.AzureMonitorPrivateLinkScope{},
			ResponseStruct: &armmonitor.PrivateLinkScopesClientListResponse{},
			Client:         &armmonitor.PrivateLinkScopesClient{},
			ListFunc:       (&armmonitor.PrivateLinkScopesClient{}).NewListPager,
			NewFunc:        armmonitor.NewPrivateLinkScopesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/microsoft.insights/privateLinkScopes",
		},
		{
			Name:           "log_profile_resource",
			Struct:         &armmonitor.LogProfileResource{},
			ResponseStruct: &armmonitor.LogProfilesClientListResponse{},
			Client:         &armmonitor.LogProfilesClient{},
			ListFunc:       (&armmonitor.LogProfilesClient{}).NewListPager,
			NewFunc:        armmonitor.NewLogProfilesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles",
		},
		{
			Name:           "event_data",
			Struct:         &armmonitor.EventData{},
			ResponseStruct: &armmonitor.TenantActivityLogsClientListResponse{},
			Client:         &armmonitor.TenantActivityLogsClient{},
			ListFunc:       (&armmonitor.TenantActivityLogsClient{}).NewListPager,
			NewFunc:        armmonitor.NewTenantActivityLogsClient,
			URL:            "/providers/Microsoft.Insights/eventtypes/management/values",
		},
	}

	for i := range tables {
		tables[i].Service = "armmonitor"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmonitor()...)
}

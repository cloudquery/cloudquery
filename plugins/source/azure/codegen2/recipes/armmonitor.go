// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

func Armmonitor() []Table {
	tables := []Table{
		{
      Name: "diagnostic_settings_resource",
      Struct: &armmonitor.DiagnosticSettingsResource{},
      ResponseStruct: &armmonitor.DiagnosticSettingsClientListResponse{},
      Client: &armmonitor.DiagnosticSettingsClient{},
      ListFunc: (&armmonitor.DiagnosticSettingsClient{}).NewListPager,
			NewFunc: armmonitor.NewDiagnosticSettingsClient,
		},
		{
      Name: "diagnostic_settings_category_resource",
      Struct: &armmonitor.DiagnosticSettingsCategoryResource{},
      ResponseStruct: &armmonitor.DiagnosticSettingsCategoryClientListResponse{},
      Client: &armmonitor.DiagnosticSettingsCategoryClient{},
      ListFunc: (&armmonitor.DiagnosticSettingsCategoryClient{}).NewListPager,
			NewFunc: armmonitor.NewDiagnosticSettingsCategoryClient,
		},
		{
      Name: "azure_monitor_private_link_scope",
      Struct: &armmonitor.AzureMonitorPrivateLinkScope{},
      ResponseStruct: &armmonitor.PrivateLinkScopesClientListResponse{},
      Client: &armmonitor.PrivateLinkScopesClient{},
      ListFunc: (&armmonitor.PrivateLinkScopesClient{}).NewListPager,
			NewFunc: armmonitor.NewPrivateLinkScopesClient,
		},
		{
      Name: "event_data",
      Struct: &armmonitor.EventData{},
      ResponseStruct: &armmonitor.ActivityLogsClientListResponse{},
      Client: &armmonitor.ActivityLogsClient{},
      ListFunc: (&armmonitor.ActivityLogsClient{}).NewListPager,
			NewFunc: armmonitor.NewActivityLogsClient,
		},
		{
      Name: "single_metric_baseline",
      Struct: &armmonitor.SingleMetricBaseline{},
      ResponseStruct: &armmonitor.BaselinesClientListResponse{},
      Client: &armmonitor.BaselinesClient{},
      ListFunc: (&armmonitor.BaselinesClient{}).NewListPager,
			NewFunc: armmonitor.NewBaselinesClient,
		},
		{
      Name: "metric_definition",
      Struct: &armmonitor.MetricDefinition{},
      ResponseStruct: &armmonitor.MetricDefinitionsClientListResponse{},
      Client: &armmonitor.MetricDefinitionsClient{},
      ListFunc: (&armmonitor.MetricDefinitionsClient{}).NewListPager,
			NewFunc: armmonitor.NewMetricDefinitionsClient,
		},
		{
      Name: "event_data",
      Struct: &armmonitor.EventData{},
      ResponseStruct: &armmonitor.TenantActivityLogsClientListResponse{},
      Client: &armmonitor.TenantActivityLogsClient{},
      ListFunc: (&armmonitor.TenantActivityLogsClient{}).NewListPager,
			NewFunc: armmonitor.NewTenantActivityLogsClient,
		},
		{
      Name: "localizable_string",
      Struct: &armmonitor.LocalizableString{},
      ResponseStruct: &armmonitor.EventCategoriesClientListResponse{},
      Client: &armmonitor.EventCategoriesClient{},
      ListFunc: (&armmonitor.EventCategoriesClient{}).NewListPager,
			NewFunc: armmonitor.NewEventCategoriesClient,
		},
		{
      Name: "metric_namespace",
      Struct: &armmonitor.MetricNamespace{},
      ResponseStruct: &armmonitor.MetricNamespacesClientListResponse{},
      Client: &armmonitor.MetricNamespacesClient{},
      ListFunc: (&armmonitor.MetricNamespacesClient{}).NewListPager,
			NewFunc: armmonitor.NewMetricNamespacesClient,
		},
		{
      Name: "log_profile_resource",
      Struct: &armmonitor.LogProfileResource{},
      ResponseStruct: &armmonitor.LogProfilesClientListResponse{},
      Client: &armmonitor.LogProfilesClient{},
      ListFunc: (&armmonitor.LogProfilesClient{}).NewListPager,
			NewFunc: armmonitor.NewLogProfilesClient,
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
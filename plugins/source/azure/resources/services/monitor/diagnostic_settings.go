// Auto generated code - DO NOT EDIT.

package monitor

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
)

func diagnosticSettings() *schema.Table {
	return &schema.Table{
		Name:        "azure_monitor_diagnostic_settings",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights#DiagnosticSettingsResource`,
		Resolver:    fetchMonitorDiagnosticSettings,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "monitor_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "storage_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageAccountID"),
			},
			{
				Name:     "service_bus_rule_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceBusRuleID"),
			},
			{
				Name:     "event_hub_authorization_rule_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventHubAuthorizationRuleID"),
			},
			{
				Name:     "event_hub_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventHubName"),
			},
			{
				Name:     "metrics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metrics"),
			},
			{
				Name:     "logs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Logs"),
			},
			{
				Name:     "workspace_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkspaceID"),
			},
			{
				Name:     "log_analytics_destination_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogAnalyticsDestinationType"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "resource_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceURI"),
			},
		},
	}
}

func isResourceTypeNotSupported(err error) bool {
	var azureErr *azure.RequestError
	if errors.As(err, &azureErr) {
		return azureErr.ServiceError != nil && azureErr.ServiceError.Code == "ResourceTypeNotSupported"
	}
	return false
}

// diagnosticSettingResource is a custom copy of insights.DiagnosticSettingsResource with extra ResourceURI field
type diagnosticSettingResource struct {
	insights.DiagnosticSettingsResource
	ResourceURI string
}

func fetchMonitorDiagnosticSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Monitor.DiagnosticSettings

	resource := parent.Item.(resources.GenericResourceExpanded)
	response, err := svc.List(ctx, *resource.ID)
	if err != nil {
		if isResourceTypeNotSupported(err) {
			return nil
		}
		return err
	}
	if response.Value == nil {
		return nil
	}
	for _, v := range *response.Value {
		res <- diagnosticSettingResource{
			DiagnosticSettingsResource: v,
			ResourceURI:                *resource.ID,
		}
	}
	return nil
}

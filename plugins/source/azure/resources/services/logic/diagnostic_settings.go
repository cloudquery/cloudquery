// Auto generated code - DO NOT EDIT.

package logic

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
)

func diagnosticSettings() *schema.Table {
	return &schema.Table{
		Name:     "azure_logic_diagnostic_settings",
		Resolver: fetchLogicDiagnosticSettings,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
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
		},
	}
}

func fetchLogicDiagnosticSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Logic.DiagnosticSettings

	workflow := parent.Item.(logic.Workflow)
	response, err := svc.List(ctx, *workflow.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}

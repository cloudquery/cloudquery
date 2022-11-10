// Auto generated code - DO NOT EDIT.

package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ActivityLogs() *schema.Table {
	return &schema.Table{
		Name:        "azure_monitor_activity_logs",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights#EventData`,
		Resolver:    fetchMonitorActivityLogs,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "authorization",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Authorization"),
			},
			{
				Name:     "claims",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Claims"),
			},
			{
				Name:     "caller",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Caller"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
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
				Name:     "event_data_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventDataID"),
			},
			{
				Name:     "correlation_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CorrelationID"),
			},
			{
				Name:     "event_name",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EventName"),
			},
			{
				Name:     "category",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Category"),
			},
			{
				Name:     "http_request",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HTTPRequest"),
			},
			{
				Name:     "level",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Level"),
			},
			{
				Name:     "resource_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceGroupName"),
			},
			{
				Name:     "resource_provider_name",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourceProviderName"),
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceID"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "operation_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OperationID"),
			},
			{
				Name:     "operation_name",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OperationName"),
			},
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "sub_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SubStatus"),
			},
			{
				Name:     "event_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EventTimestamp"),
			},
			{
				Name:     "submission_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SubmissionTimestamp"),
			},
			{
				Name:     "tenant_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TenantID"),
			},
		},
	}
}

func fetchMonitorActivityLogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Monitor.ActivityLogs

	const fetchWindow = 24 * time.Hour
	now := time.Now().UTC()
	past := now.Add(-fetchWindow)
	filter := fmt.Sprintf("eventTimestamp ge '%s' and eventTimestamp le '%s'", past.Format(time.RFC3339Nano), now.Format(time.RFC3339Nano))
	response, err := svc.List(ctx, filter, "")

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}

package monitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ActivityLogAlerts() *schema.Table {
	return &schema.Table{
		Name:      "azure_monitor_activity_log_alerts",
		Resolver:  fetchActivityLogAlerts,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_monitor_activity_log_alerts", client.Namespacemicrosoft_insights),
		Transform: transformers.TransformWithStruct(&armmonitor.ActivityLogAlertResource{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchActivityLogAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmonitor.NewActivityLogAlertsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListBySubscriptionIDPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

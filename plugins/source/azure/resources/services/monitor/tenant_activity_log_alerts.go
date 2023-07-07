package monitor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func TenantActivityLogAlerts() *schema.Table {
	return &schema.Table{
		Name:                 "azure_monitor_tenant_activity_log_alerts",
		Resolver:             fetchTenantActivityLogAlerts,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/monitor/activity-log-alerts/list-by-subscription-id?tabs=HTTP#activitylogalertresource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_monitor_tenant_activity_log_alerts", client.Namespacemicrosoft_insights),
		Transform:            transformers.TransformWithStruct(&armmonitor.ActivityLogAlertResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},
	}
}

func fetchTenantActivityLogAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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

package monitor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AutoscaleSettings() *schema.Table {
	return &schema.Table{
		Name:                 "azure_monitor_autoscale_settings",
		Resolver:             fetchAutoscaleSetting,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/monitor/autoscale-settings/list-by-subscription?tabs=HTTP#autoscalesettingresource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_monitor_autoscale_settings", client.Namespacemicrosoft_insights),
		Transform:            transformers.TransformWithStruct(&armmonitor.AutoscaleSettingResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchAutoscaleSetting(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmonitor.NewAutoscaleSettingsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListBySubscriptionPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

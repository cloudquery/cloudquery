package monitor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SubscriptionDiagnosticSettings() *schema.Table {
	return &schema.Table{
		Name:                 "azure_monitor_subscription_diagnostic_settings",
		Resolver:             fetchSubscriptionDiagnosticSettings,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/monitor/subscription-diagnostic-settings/list?tabs=HTTP#subscriptiondiagnosticsettingsresource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_monitor_subscription_diagnostic_settings", client.Namespacemicrosoft_insights),
		Transform:            transformers.TransformWithStruct(&armmonitor.DiagnosticSettingsResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchSubscriptionDiagnosticSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmonitor.NewDiagnosticSettingsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager("/subscriptions/"+cl.SubscriptionId, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, ds := range p.Value {
			res <- ds
		}
	}
	return nil
}

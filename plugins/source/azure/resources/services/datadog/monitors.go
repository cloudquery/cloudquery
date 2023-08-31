package datadog

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Monitors() *schema.Table {
	return &schema.Table{
		Name:                 "azure_datadog_monitors",
		Resolver:             fetchMonitors,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/datadog/monitors/list?tabs=HTTP#datadogmonitorresource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_datadog_monitors", client.Namespacemicrosoft_datadog),
		Transform:            transformers.TransformWithStruct(&armdatadog.MonitorResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchMonitors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdatadog.NewMonitorsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

// Code generated by codegen; DO NOT EDIT.

package armelastic

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func MonitoredResource() *schema.Table {
	return &schema.Table{
		Name:      "azure_armelastic_monitored_resource",
		Resolver:  fetchMonitoredResource,
		Multiplex: client.SubscriptionResourceGroupMultiplex,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "reason_for_logs_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReasonForLogsStatus"),
			},
			{
				Name:     "sending_logs",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SendingLogs"),
			},
		},
	}
}

func fetchMonitoredResource(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armelastic.NewMonitoredResourcesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(cl.ResourceGroup, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

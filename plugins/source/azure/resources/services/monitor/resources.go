// Auto generated code - DO NOT EDIT.

package monitor

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Resources() *schema.Table {
	return &schema.Table{
		Name:      "azure_monitor_resources",
		Resolver:  fetchMonitorResources,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			diagnosticSettings(),
		},
	}
}

func fetchMonitorResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Monitor.Resources

	// Add subscription id as the first entry
	res <- struct {
		ID string
	}{ID: "/subscriptions/" + meta.(*client.Client).SubscriptionId}
	response, err := svc.List(ctx, "", "", nil)

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

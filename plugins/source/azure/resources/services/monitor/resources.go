// Auto generated code - DO NOT EDIT.

package monitor

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
)

func Resources() *schema.Table {
	return &schema.Table{
		Name:        "azure_monitor_resources",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources#GenericResourceExpanded`,
		Resolver:    fetchMonitorResources,
		Multiplex:   client.SubscriptionMultiplex,
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
	subscriptionId := "/" + client.ScopeSubscription(meta.(*client.Client).SubscriptionId)
	res <- resources.GenericResourceExpanded{ID: &subscriptionId}
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

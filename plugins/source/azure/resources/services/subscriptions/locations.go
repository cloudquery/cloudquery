// Auto generated code - DO NOT EDIT.

package subscriptions

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:      "azure_subscriptions_locations",
		Resolver:  fetchSubscriptionsLocations,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
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
				Name:     "regional_display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegionalDisplayName"),
			},
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscriptionID"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchSubscriptionsLocations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions.Locations
	pager := svc.NewListLocationsPager(meta.(*client.Client).Services().Subscriptions.SubscriptionID, nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, v := range nextResult.Value {
			res <- v
		}
	}
	return nil
}

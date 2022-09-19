// Auto generated code - DO NOT EDIT.

package resources

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Links() *schema.Table {
	return &schema.Table{
		Name:      "azure_resources_links",
		Resolver:  fetchResourcesLinks,
		Multiplex: client.SubscriptionMultiplex,
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
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "properties_source_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.SourceID"),
			},
			{
				Name:     "properties_target_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.TargetID"),
			},
			{
				Name:     "properties_notes",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Notes"),
			},
		},
	}
}

func fetchResourcesLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Resources.Links

	response, err := svc.ListAtSubscription(ctx, "")

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

// Auto generated code - DO NOT EDIT.

package resources

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
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
				Name:     "source_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceID"),
			},
			{
				Name:     "target_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetID"),
			},
			{
				Name:     "notes",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Notes"),
			},
		},
	}
}

func fetchResourcesLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Resources.Links

	response, err := svc.ListAtSubscription(ctx, "")

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

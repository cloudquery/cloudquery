package resources

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Links() *schema.Table {
	return &schema.Table{
		Name:      "azure_resources_links",
		Resolver:  fetchLinks,
		Multiplex: client.SubscriptionMultiplex,
		Transform: transformers.TransformWithStruct(&armlinks.ResourceLink{}),
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
	}
}

func fetchLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armlinks.NewResourceLinksClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListAtSubscriptionPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

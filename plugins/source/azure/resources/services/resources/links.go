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
		Name:        "azure_resources_links",
		Resolver:    fetchLinks,
		Description: "https://learn.microsoft.com/en-us/rest/api/resources/resource-links/list-at-subscription#resourcelink",
		Multiplex:   client.SubscriptionMultiplex,
		Transform:   transformers.TransformWithStruct(&armlinks.ResourceLink{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
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

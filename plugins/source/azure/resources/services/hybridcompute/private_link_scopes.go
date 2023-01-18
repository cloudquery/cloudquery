package hybridcompute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PrivateLinkScopes() *schema.Table {
	return &schema.Table{
		Name:        "azure_hybridcompute_private_link_scopes",
		Resolver:    fetchPrivateLinkScopes,
		Description: "https://learn.microsoft.com/en-us/rest/api/hybridcompute/private-link-scopes/list?tabs=HTTP#hybridcomputeprivatelinkscope",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_hybridcompute_private_link_scopes", client.Namespacemicrosoft_hybridcompute),
		Transform:   transformers.TransformWithStruct(&armhybridcompute.PrivateLinkScope{}),
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
		},
	}
}

func fetchPrivateLinkScopes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armhybridcompute.NewPrivateLinkScopesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

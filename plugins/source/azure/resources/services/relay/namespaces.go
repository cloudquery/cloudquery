package relay

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:        "azure_relay_namespaces",
		Resolver:    fetchNamespaces,
		Description: "https://learn.microsoft.com/en-us/rest/api/relay/namespaces/list?tabs=HTTP#relaynamespace",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_relay_namespaces", client.Namespacemicrosoft_relay),
		Transform:   transformers.TransformWithStruct(&armrelay.Namespace{}),
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

func fetchNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armrelay.NewNamespacesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

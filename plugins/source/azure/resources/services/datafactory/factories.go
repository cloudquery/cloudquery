package datafactory

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Factories() *schema.Table {
	return &schema.Table{
		Name:      "azure_datafactory_factories",
		Resolver:  fetchFactories,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_datafactory_factories", client.Namespacemicrosoft_datafactory),
		Transform: transformers.TransformWithStruct(&armdatafactory.Factory{}, transformers.WithSkipFields("ETag")),
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
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ETag"),
			},
		},
	}
}

func fetchFactories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdatafactory.NewFactoriesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

package saas

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Resources() *schema.Table {
	return &schema.Table{
		Name:      "azure_saas_resources",
		Resolver:  fetchResources,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_saas_resources", client.Namespacemicrosoft_saas),
		Transform: transformers.TransformWithStruct(&armsaas.Resource{}),
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

func fetchResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsaas.NewResourcesClient(cl.Creds, cl.Options)
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

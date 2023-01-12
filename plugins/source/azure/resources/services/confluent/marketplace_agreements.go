package confluent

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func MarketplaceAgreements() *schema.Table {
	return &schema.Table{
		Name:      "azure_confluent_marketplace_agreements",
		Resolver:  fetchMarketplaceAgreements,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_confluent_marketplace_agreements", client.Namespacemicrosoft_confluent),
		Transform: transformers.TransformWithStruct(&armconfluent.AgreementResource{}),
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

func fetchMarketplaceAgreements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconfluent.NewMarketplaceAgreementsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

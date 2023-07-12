package confluent

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func MarketplaceAgreements() *schema.Table {
	return &schema.Table{
		Name:                 "azure_confluent_marketplace_agreements",
		Resolver:             fetchMarketplaceAgreements,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/confluent/marketplace-agreements/list?tabs=HTTP#confluentagreementresource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_confluent_marketplace_agreements", client.Namespacemicrosoft_confluent),
		Transform:            transformers.TransformWithStruct(&armconfluent.AgreementResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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

package datadog

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func MarketplaceAgreements() *schema.Table {
	return &schema.Table{
		Name:        "azure_datadog_marketplace_agreements",
		Resolver:    fetchMarketplaceAgreements,
		Description: "https://learn.microsoft.com/en-us/rest/api/datadog/marketplace-agreements/list?tabs=HTTP#datadogagreementresource",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_datadog_marketplace_agreements", client.Namespacemicrosoft_datadog),
		Transform:   transformers.TransformWithStruct(&armdatadog.AgreementResource{}),
		Columns: []schema.Column{
			client.SubscriptionID,
			client.IDColumn,
		},
	}
}

func fetchMarketplaceAgreements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdatadog.NewMarketplaceAgreementsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

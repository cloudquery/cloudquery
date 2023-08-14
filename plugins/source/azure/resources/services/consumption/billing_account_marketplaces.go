package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BillingAccountMarketplaces() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_billing_account_marketplaces",
		Resolver:             fetchBillingAccountMarketplaces,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/marketplaces/list?tabs=HTTP#marketplace",
		Multiplex:            client.BillingAccountMultiplex,
		Transform:            transformers.TransformWithStruct(&armconsumption.Marketplace{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingAccountMarketplaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewMarketplacesClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(*cl.BillingAccount.ID, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

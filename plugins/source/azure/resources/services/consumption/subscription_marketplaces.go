package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SubscriptionMarketplaces() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_subscription_marketplaces",
		Resolver:             fetchSubscriptionMarketplaces,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/budgets/list?tabs=HTTP#budget",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_consumption_subscription_marketplaces", client.Namespacemicrosoft_consumption),
		Transform:            transformers.TransformWithStruct(&armconsumption.Marketplace{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchSubscriptionMarketplaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewMarketplacesClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	scope := "subscriptions/" + cl.SubscriptionId
	pager := svc.NewListPager(scope, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

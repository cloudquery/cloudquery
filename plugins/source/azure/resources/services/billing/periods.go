package billing

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Periods() *schema.Table {
	return &schema.Table{
		Name:      "azure_billing_periods",
		Resolver:  fetchPeriods,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_billing_periods", client.Namespacemicrosoft_billing),
		Transform: transformers.TransformWithStruct(&armbilling.Period{}),
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

func fetchPeriods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armbilling.NewPeriodsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

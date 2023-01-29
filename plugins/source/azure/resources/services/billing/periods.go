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
		Name:        "azure_billing_periods",
		Resolver:    fetchPeriods,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling@v0.5.0#Period",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_billing_periods", client.Namespacemicrosoft_billing),
		Transform:   transformers.TransformWithStruct(&armbilling.Period{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
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

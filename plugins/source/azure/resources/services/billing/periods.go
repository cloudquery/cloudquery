package billing

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Periods() *schema.Table {
	return &schema.Table{
		Name:                 "azure_billing_periods",
		Resolver:             fetchPeriods,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling@v0.5.0#Period",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_billing_periods", client.Namespacemicrosoft_billing),
		Transform:            transformers.TransformWithStruct(&armbilling.Period{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPeriods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	// we already fetch all billing periods during initialization, so no need to fetch them again
	cl := meta.(*client.Client)
	res <- cl.BillingPeriods[cl.SubscriptionId]
	return nil
}

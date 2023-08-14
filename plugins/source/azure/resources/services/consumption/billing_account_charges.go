package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BillingAccountCharges() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_billing_account_charges",
		Resolver:             fetchBillingAccountCharges,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/charges/list?tabs=HTTP#modernchargesummary",
		Multiplex:            client.BillingAccountMultiplex,
		Transform:            transformers.TransformWithStruct(&armconsumption.ModernChargeSummary{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingAccountCharges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewChargesClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	resp, err := svc.List(ctx, *cl.BillingAccount.ID, nil)
	if err != nil {
		return err
	}
	for _, v := range resp.Value {
		res <- v.(*armconsumption.ModernChargeSummary)
	}
	return nil
}

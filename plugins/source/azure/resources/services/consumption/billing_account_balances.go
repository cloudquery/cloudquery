package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BillingAccountBalances() *schema.Table {
	return &schema.Table{
		Name:        "azure_consumption_billing_account_balances",
		Resolver:    fetchBillingAccountBalances,
		Description: "https://learn.microsoft.com/en-us/rest/api/consumption/balances/get-by-billing-account?tabs=HTTP#balance",
		Multiplex:   client.BillingAccountMultiplex,
		Transform:   transformers.TransformWithStruct(&armconsumption.Balance{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingAccountBalances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewBalancesClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	resp, err := svc.GetByBillingAccount(ctx, *cl.BillingAccount.ID, nil)
	if err != nil {
		return err
	}
	res <- resp.Balance
	return nil
}

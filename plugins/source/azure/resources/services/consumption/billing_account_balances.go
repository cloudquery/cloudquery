package consumption

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BillingAccountBalances() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_billing_account_balances",
		Resolver:             fetchBillingAccountBalances,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/balances/get-by-billing-account?tabs=HTTP#balance",
		Multiplex:            client.BillingAccountMultiplex,
		Transform:            transformers.TransformWithStruct(&armconsumption.Balance{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingAccountBalances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewBalancesClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	resp, err := svc.GetByBillingAccount(ctx, *cl.BillingAccount.Name, nil)
	if err != nil {
		var respError *azcore.ResponseError
		// If there's no data a 404 error is returned so we ignore it
		if errors.As(err, &respError) && respError.StatusCode == 404 {
			cl.Logger().Debug().Msg("No data for billing account balances")
			return nil
		}
		return err
	}
	res <- resp.Balance
	return nil
}

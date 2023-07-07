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

func BillingAccountLots() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_billing_account_lots",
		Resolver:             fetchBillingAccountLots,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/lots/list-by-billing-account?tabs=HTTP#lotsummary",
		Multiplex:            client.BillingAccountMultiplex,
		Transform:            transformers.TransformWithStruct(&armconsumption.LotSummary{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingAccountLots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewLotsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListByBillingAccountPager(*cl.BillingAccount.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			var respError *azcore.ResponseError
			// If there's no data a 404 error is returned so we ignore it
			if errors.As(err, &respError) && respError.StatusCode == 404 {
				cl.Logger().Debug().Msg("No data for billing account lots")
				return nil
			}
			return err
		}
		res <- p.Value
	}
	return nil
}

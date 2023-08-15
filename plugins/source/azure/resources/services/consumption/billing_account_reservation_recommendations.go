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

func BillingAccountReservationRecommendations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_billing_account_reservation_recommendations",
		Resolver:             fetchBillingAccountReservationRecommendations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/reservation-recommendations/list?tabs=HTTP#legacyreservationrecommendation",
		Multiplex:            client.LegacyBillingAccountMultiplex,
		Transform:            transformers.TransformWithStruct(&armconsumption.LegacyReservationRecommendation{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingAccountReservationRecommendations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewReservationRecommendationsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(*cl.BillingAccount.ID, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			var respError *azcore.ResponseError
			// If there's no data a 204 error is returned so we ignore it
			if errors.As(err, &respError) && respError.StatusCode == 204 {
				cl.Logger().Debug().Msg("No data for billing profile reservation recommendations")
				return nil
			}
			return err
		}
		for _, v := range p.Value {
			res <- v.(*armconsumption.LegacyReservationRecommendation)
		}
	}
	return nil
}

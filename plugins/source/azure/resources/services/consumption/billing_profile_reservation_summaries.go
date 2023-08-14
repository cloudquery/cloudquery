package consumption

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BillingProfileReservationSummaries() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_billing_profile_reservation_summaries",
		Resolver:             fetchBillingProfileReservationSummaries,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/reservations-summaries/list?tabs=HTTP#reservationsummary",
		Multiplex:            client.BillingAccountProfileMultiplex,
		Transform:            transformers.TransformWithStruct(&armconsumption.ReservationSummary{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingProfileReservationSummaries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewReservationsSummariesClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	today := to.Ptr(time.Now().UTC().AddDate(0, 0, 1).Format("2006-01-02"))
	pager := svc.NewListPager(*cl.BillingProfile.ID, armconsumption.DatagrainDailyGrain, &armconsumption.ReservationsSummariesClientListOptions{StartDate: to.Ptr("2000-01-01"), EndDate: today})
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

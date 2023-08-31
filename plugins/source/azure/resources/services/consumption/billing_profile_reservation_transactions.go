package consumption

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BillingProfileReservationTransactions() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_billing_profile_reservation_transactions",
		Resolver:             fetchBillingProfileReservationTransactions,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/reservation-transactions/list?tabs=HTTP#reservationtransaction",
		Multiplex:            client.BillingAccountProfileMultiplex,
		Transform:            transformers.TransformWithStruct(&armconsumption.ReservationTransaction{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingProfileReservationTransactions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewReservationTransactionsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	end := now.AddDate(0, 0, 1).Format("2006-01-02")
	start := now.AddDate(-1, 0, 0).Format("2006-01-02")
	filter := to.Ptr(fmt.Sprintf("properties/eventDate ge %s AND properties/eventDate le %s", start, end))
	pager := svc.NewListByBillingProfilePager(*cl.BillingAccount.Name, *cl.BillingProfile.Name, &armconsumption.ReservationTransactionsClientListByBillingProfileOptions{Filter: filter})
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

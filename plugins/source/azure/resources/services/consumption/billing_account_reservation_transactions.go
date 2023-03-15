package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BillingAccountReservationTransactions() *schema.Table {
	return &schema.Table{
		Name:        "azure_consumption_billing_account_reservation_transactions",
		Resolver:    fetchBillingAccountReservationTransactions,
		Description: "https://learn.microsoft.com/en-us/rest/api/consumption/reservation-transactions/list?tabs=HTTP#reservationtransaction",
		Multiplex:   client.BillingAccountMultiplex,
		Transform:   transformers.TransformWithStruct(&armconsumption.ReservationTransaction{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingAccountReservationTransactions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewReservationTransactionsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(*cl.BillingAccount.ID, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

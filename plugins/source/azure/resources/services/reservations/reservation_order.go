package reservations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ReservationOrder() *schema.Table {
	return &schema.Table{
		Name:        "azure_reservations_reservation_order",
		Resolver:    fetchReservationOrder,
		Description: "https://learn.microsoft.com/en-us/rest/api/reserved-vm-instances/reservation-order/get?tabs=HTTP#reservationorderresponse",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_reservations_reservation_order", client.Namespacemicrosoft_capacity),
		Transform:   transformers.TransformWithStruct(&armreservations.ReservationOrderResponse{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchReservationOrder(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armreservations.NewReservationOrderClient(cl.Creds, cl.Options)
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

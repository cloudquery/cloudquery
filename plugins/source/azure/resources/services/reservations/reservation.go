package reservations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Reservation() *schema.Table {
	return &schema.Table{
		Name:                 "azure_reservations_reservation",
		Resolver:             fetchReservation,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/reserved-vm-instances/reservation/list-all?tabs=HTTP#reservationresponse",
		Transform:            transformers.TransformWithStruct(&armreservations.ReservationResponse{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchReservation(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armreservations.NewReservationClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListAllPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

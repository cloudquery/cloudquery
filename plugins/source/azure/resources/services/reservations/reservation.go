package reservations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Reservation() *schema.Table {
	return &schema.Table{
		Name:        "azure_reservations_reservation",
		Resolver:    fetchReservation,
		Description: "https://learn.microsoft.com/en-us/rest/api/reserved-vm-instances/reservation/get?tabs=HTTP#reservationresponse",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_reservations_reservation", client.Namespacemicrosoft_capacity),
		Transform:   transformers.TransformWithStruct(&armreservations.ReservationResponse{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
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

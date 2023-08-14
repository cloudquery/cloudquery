package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func capacityReservations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_capacity_reservations",
		Resolver:             fetchCapacityReservations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/capacity-reservations/list-by-capacity-reservation-group?tabs=HTTP#capacityreservation",
		Transform:            transformers.TransformWithStruct(&armcompute.CapacityReservation{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchCapacityReservations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	capacityGroup := parent.Item.(*armcompute.CapacityReservationGroup)
	svc, err := armcompute.NewCapacityReservationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	resourceGroup, err := client.ParseResourceGroup(*capacityGroup.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListByCapacityReservationGroupPager(resourceGroup, *capacityGroup.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

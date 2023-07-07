package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func CapacityReservationGroups() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_capacity_reservation_groups",
		Resolver:             fetchCapacityReservationGroups,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/capacity-reservation-groups/list-by-resource-group?tabs=HTTP#capacityreservationgroup",
		Multiplex:            client.SubscriptionResourceGroupMultiplexRegisteredNamespace("azure_compute_capacity_reservation_groups", client.Namespacemicrosoft_compute),
		Transform:            transformers.TransformWithStruct(&armcompute.CapacityReservationGroup{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations:            []*schema.Table{capacityReservations()},
	}
}

func fetchCapacityReservationGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewCapacityReservationGroupsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	expand := to.Ptr(armcompute.ExpandTypesForGetCapacityReservationGroupsVirtualMachinesRef + "," + armcompute.ExpandTypesForGetCapacityReservationGroupsVirtualMachineScaleSetVMsRef)
	pager := svc.NewListByResourceGroupPager(cl.ResourceGroup, &armcompute.CapacityReservationGroupsClientListByResourceGroupOptions{Expand: expand})
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

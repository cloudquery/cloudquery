package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchVirtualMachineInstanceViews(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewVirtualMachinesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	item := parent.Item.(*armcompute.VirtualMachine)
	group, err := client.ParseResourceGroup(*item.ID)
	if err != nil {
		return err
	}
	instanceView, err := svc.InstanceView(ctx, group, *item.Name, nil)
	if err != nil {
		return err
	}
	// views := []any{instanceView}
	res <- instanceView.VirtualMachineInstanceView
	return nil
}

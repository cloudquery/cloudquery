package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func getInstanceView(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewVirtualMachinesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	item := resource.Item.(*armcompute.VirtualMachine)
	group, err := client.ParseResourceGroup(*item.ID)
	if err != nil {
		return err
	}
	instanceView, err := svc.InstanceView(ctx, group, *item.Name, nil)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, instanceView.VirtualMachineInstanceView)
}

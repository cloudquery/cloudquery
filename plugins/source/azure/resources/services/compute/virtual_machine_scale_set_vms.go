package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func VirtualMachineScaleSetsVMs() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_virtual_machine_scale_set_vms",
		Resolver:             fetchVirtualMachineScaleSetsVMs,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-scale-set-vms/list?tabs=HTTP#virtualmachinescalesetvm",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_virtual_machine_scale_sets_vms", client.Namespacemicrosoft_compute),
		Transform:            transformers.TransformWithStruct(&armcompute.VirtualMachineScaleSetVM{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVirtualMachineScaleSetsVMs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	scaleSet := parent.Item.(*armcompute.VirtualMachineScaleSet)
	svc, err := armcompute.NewVirtualMachineScaleSetVMsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*scaleSet.ID)
	if err != nil {
		return err
	}
	expand := "instanceView"
	pager := svc.NewListPager(group, *scaleSet.Name, &armcompute.VirtualMachineScaleSetVMsClientListOptions{Expand: &expand})
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

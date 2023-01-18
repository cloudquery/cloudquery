package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VirtualMachineScaleSetsVMs() *schema.Table {
	return &schema.Table{
		Name:        "azure_compute_virtual_machine_scale_set_vms",
		Resolver:    fetchVirtualMachineScaleSetsVMs,
		Description: "https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-scale-set-vms/list?tabs=HTTP#virtualmachinescalesetvm",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_compute_virtual_machine_scale_sets_vms", client.Namespacemicrosoft_compute),
		Transform:   transformers.TransformWithStruct(&armcompute.VirtualMachineScaleSetVM{}),
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
	pager := svc.NewListPager(group, *scaleSet.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

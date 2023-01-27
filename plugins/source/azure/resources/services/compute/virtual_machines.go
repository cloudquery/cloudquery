package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VirtualMachines() *schema.Table {
	return &schema.Table{
		Name:        "azure_compute_virtual_machines",
		Resolver:    fetchVirtualMachines,
		Description: "https://learn.microsoft.com/en-us/rest/api/compute/virtual-machines/list?tabs=HTTP#virtualmachine",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_compute_virtual_machines", client.Namespacemicrosoft_compute),
		Transform:   transformers.TransformWithStruct(&armcompute.VirtualMachine{}),
		Columns: []schema.Column{
			client.SubscriptionID,
			{
				Name:     "instance_view",
				Type:     schema.TypeJSON,
				Resolver: getInstanceView,
			},
			client.IDColumn,
		},
		Relations: []*schema.Table{
			VirtualMachineExtensions(),
		},
	}
}

func fetchVirtualMachines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewVirtualMachinesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VirtualMachineScaleSets() *schema.Table {
	return &schema.Table{
		Name:      "azure_compute_virtual_machine_scale_sets",
		Resolver:  fetchVirtualMachineScaleSets,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_compute_virtual_machine_scale_sets", client.Namespacemicrosoft_compute),
		Transform: transformers.TransformWithStruct(&armcompute.VirtualMachineScaleSet{}),
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
		Relations: []*schema.Table{
			VirtualMachineScaleSetsVMs(),
		},
	}
}

func fetchVirtualMachineScaleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewVirtualMachineScaleSetsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func VirtualMachineScaleSets() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_virtual_machine_scale_sets",
		Resolver:             fetchVirtualMachineScaleSets,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-scale-sets/list?tabs=HTTP#virtualmachinescaleset",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_virtual_machine_scale_sets", client.Namespacemicrosoft_compute),
		Transform:            transformers.TransformWithStruct(&armcompute.VirtualMachineScaleSet{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations: []*schema.Table{
			VirtualMachineScaleSetsVMs(),
			virtualMachineScaleSetsNetworkInterfaces(),
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

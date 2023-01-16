package compute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VirtualMachineExtensions() *schema.Table {
	return &schema.Table{
		Name:      "azure_compute_virtual_machine_extensions",
		Resolver:  fetchVirtualMachineExtnsions,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_compute_virtual_machine_extensions", client.Namespacemicrosoft_compute),
		Transform: transformers.TransformWithStruct(&armcompute.VirtualMachineExtension{}),
		Columns: []schema.Column{
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

func fetchVirtualMachineExtnsions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armcompute.VirtualMachine)
	cl := meta.(*client.Client)
	svc, err := armcompute.NewVirtualMachineExtensionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	resp, err := svc.List(ctx, group, *p.Name, nil)
	if err != nil {
		return err
	}
	res <- resp.Value
	return nil
}

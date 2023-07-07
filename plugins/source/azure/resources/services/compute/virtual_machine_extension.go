package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func VirtualMachineExtensions() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_virtual_machine_extensions",
		Resolver:             fetchVirtualMachineExtnsions,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-extensions/list?tabs=HTTP#virtualmachineextension",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_virtual_machine_extensions", client.Namespacemicrosoft_compute),
		Transform:            transformers.TransformWithStruct(&armcompute.VirtualMachineExtension{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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

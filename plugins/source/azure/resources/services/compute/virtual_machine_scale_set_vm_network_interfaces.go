package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func virtualMachineScaleSetsNetworkInterfaces() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_virtual_machine_scale_set_network_interfaces",
		Resolver:             fetchVirtualMachineScaleSetsNetworkInterfaces,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/network-interface-in-vm-ss/list-virtual-machine-scale-set-network-interfaces?tabs=HTTP#networkinterface",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_virtual_machine_scale_set_network_interfaces", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.Interface{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVirtualMachineScaleSetsNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	scaleSet := parent.Item.(*armcompute.VirtualMachineScaleSet)
	svc, err := armnetwork.NewInterfacesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*scaleSet.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListVirtualMachineScaleSetNetworkInterfacesPager(group, *scaleSet.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

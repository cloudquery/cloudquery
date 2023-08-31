package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func VirtualNetworks() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_virtual_networks",
		Resolver:             fetchVirtualNetworks,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/virtual-networks/list-all?tabs=HTTP#virtualnetwork",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_virtual_networks", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.VirtualNetwork{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations: []*schema.Table{
			virtualNetworkSubnets(),
		},
	}
}

func fetchVirtualNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewVirtualNetworksClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

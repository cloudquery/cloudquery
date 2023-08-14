package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func virtualNetworkSubnets() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_virtual_network_subnets",
		Resolver:             fetchVirtualNetworkSubnets,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/subnets/list?tabs=HTTP#subnet",
		Transform:            transformers.TransformWithStruct(&armnetwork.Subnet{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVirtualNetworkSubnets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armnetwork.VirtualNetwork)
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewSubnetsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(group, *p.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

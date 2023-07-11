package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BastionHosts() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_bastion_hosts",
		Resolver:             fetchBastionHosts,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/bastion-hosts/list?tabs=HTTP#bastionhost",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_bastion_hosts", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.BastionHost{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchBastionHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewBastionHostsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func IpGroups() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_ip_groups",
		Resolver:    fetchIpGroups,
		Description: "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/ip-groups/list?tabs=HTTP#ipgroup",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_ip_groups", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.IPGroup{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchIpGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewIPGroupsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

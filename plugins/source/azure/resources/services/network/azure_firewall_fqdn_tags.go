package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AzureFirewallFqdnTags() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_azure_firewall_fqdn_tags",
		Resolver:             fetchAzureFirewallFqdnTags,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/firewall/azure-firewall-fqdn-tags/list-all?tabs=HTTP#azurefirewallfqdntag",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_azure_firewall_fqdn_tags", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.AzureFirewallFqdnTag{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionIDPK},
	}
}

func fetchAzureFirewallFqdnTags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewAzureFirewallFqdnTagsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

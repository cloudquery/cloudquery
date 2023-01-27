package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AzureFirewalls() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_azure_firewalls",
		Resolver:    fetchAzureFirewalls,
		Description: "https://learn.microsoft.com/en-us/rest/api/firewall/azure-firewalls/list?tabs=HTTP#azurefirewall",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_azure_firewalls", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.AzureFirewall{}),
		Columns: []schema.Column{
			client.SubscriptionID,
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

func fetchAzureFirewalls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewAzureFirewallsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func FirewallPolicies() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_firewall_policies",
		Resolver:    fetchFirewallPolicies,
		Description: "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/firewall-policies/list-all?tabs=HTTP#firewallpolicy",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_firewall_policies", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.FirewallPolicy{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
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

func fetchFirewallPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewFirewallPoliciesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WebApplicationFirewallPolicies() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_web_application_firewall_policies",
		Resolver:    fetchWebApplicationFirewallPolicies,
		Description: "https://learn.microsoft.com/en-us/rest/api/application-gateway/web-application-firewall-policies/list?tabs=HTTP#webapplicationfirewallpolicy",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_web_application_firewall_policies", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.WebApplicationFirewallPolicy{}),
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

func fetchWebApplicationFirewallPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewWebApplicationFirewallPoliciesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

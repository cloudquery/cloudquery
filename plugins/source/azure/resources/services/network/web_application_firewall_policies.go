package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func WebApplicationFirewallPolicies() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_web_application_firewall_policies",
		Resolver:             fetchWebApplicationFirewallPolicies,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/application-gateway/web-application-firewall-policies/list?tabs=HTTP#webapplicationfirewallpolicy",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_web_application_firewall_policies", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.WebApplicationFirewallPolicy{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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

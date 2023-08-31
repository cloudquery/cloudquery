package dnsresolver

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DnsForwardingRulesets() *schema.Table {
	return &schema.Table{
		Name:                 "azure_dnsresolver_dns_forwarding_rulesets",
		Resolver:             fetchDnsForwardingRulesets,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/dns/dnsresolver/dns-forwarding-rulesets/list?tabs=HTTP#dnsforwardingruleset",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_dnsresolver_dns_forwarding_rulesets", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armdnsresolver.DNSForwardingRuleset{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchDnsForwardingRulesets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdnsresolver.NewDNSForwardingRulesetsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

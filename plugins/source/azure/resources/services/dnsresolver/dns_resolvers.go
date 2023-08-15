package dnsresolver

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DnsResolvers() *schema.Table {
	return &schema.Table{
		Name:                 "azure_dnsresolver_dns_resolvers",
		Resolver:             fetchDnsResolvers,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/dns/dnsresolver/dns-resolvers/list?tabs=HTTP#dnsresolver",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_dnsresolver_dns_resolvers", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armdnsresolver.DNSResolver{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchDnsResolvers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdnsresolver.NewDNSResolversClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

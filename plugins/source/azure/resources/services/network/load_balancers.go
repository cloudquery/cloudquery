package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_load_balancers",
		Resolver:             fetchLoadBalancers,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/load-balancer/load-balancers/list?tabs=HTTP#loadbalancer",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_load_balancers", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.LoadBalancer{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchLoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewLoadBalancersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

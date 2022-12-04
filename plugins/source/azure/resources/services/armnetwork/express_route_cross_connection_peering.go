// Code generated by codegen; DO NOT EDIT.

package armnetwork

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ExpressRouteCrossConnectionPeering() *schema.Table {
	return &schema.Table{
		Name:      "azure_armnetwork_express_route_cross_connection_peering",
		Resolver:  fetchExpressRouteCrossConnectionPeering,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
		},
	}
}

func fetchExpressRouteCrossConnectionPeering(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmnetworkExpressRouteCrossConnectionPeering
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

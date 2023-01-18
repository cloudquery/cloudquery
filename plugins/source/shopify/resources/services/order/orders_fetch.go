package order

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchOrders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	cursor := ""
	for {
		ret, cur, err := cl.Services.GetOrders(ctx, cursor)
		if err != nil {
			return err
		}
		res <- ret.Orders

		if len(ret.Orders) < ret.PageSize || cur == "" {
			return nil
		}

		cursor = cur
	}
}

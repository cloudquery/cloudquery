package product

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	cursor := ""
	for {
		ret, cur, err := cl.Services.GetProducts(ctx, cursor)
		if err != nil {
			return err
		}
		res <- ret.Products

		if len(ret.Products) < ret.PageSize || cur == "" {
			return nil
		}

		cursor = cur
	}
}

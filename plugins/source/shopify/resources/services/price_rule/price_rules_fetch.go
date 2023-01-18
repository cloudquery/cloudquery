package price_rule

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchPriceRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	cursor := ""
	for {
		ret, cur, err := cl.Services.GetPriceRules(ctx, cursor)
		if err != nil {
			return err
		}
		res <- ret.PriceRules

		if len(ret.PriceRules) < ret.PageSize || cur == "" {
			return nil
		}

		cursor = cur
	}
}

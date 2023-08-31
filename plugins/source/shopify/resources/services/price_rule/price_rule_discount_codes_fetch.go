package price_rule

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchPriceRuleDiscountCodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(shopify.PriceRule)

	cursor := ""
	for {
		ret, cur, err := cl.Services.GetDiscountCodes(ctx, p.ID, cursor)
		if err != nil {
			return err
		}
		res <- ret.DiscountCodes

		if len(ret.DiscountCodes) < ret.PageSize || cur == "" {
			return nil
		}

		cursor = cur
	}
}

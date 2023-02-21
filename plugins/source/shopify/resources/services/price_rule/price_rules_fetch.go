package price_rule

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchPriceRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	const key = "price_rules"

	p := url.Values{}
	min := time.Time{}

	if cl.Backend != nil {
		value, err := cl.Backend.Get(ctx, key, cl.ID())
		if err != nil {
			return fmt.Errorf("failed to retrieve state from backend: %w", err)
		}
		if value != "" {
			min, err = time.Parse(time.RFC3339, value)
			if err != nil {
				return fmt.Errorf("retrieved invalid state value: %q %w", value, err)
			}
			p.Set("updated_at_min", min.Format(time.RFC3339))
		}
	}

	cursor := ""
	for {
		ret, cur, err := cl.Services.GetPriceRules(ctx, cursor, p)
		if err != nil {
			return err
		}

		for i := range ret.PriceRules {
			if ret.PriceRules[i].UpdatedAt.After(min) {
				min = ret.PriceRules[i].UpdatedAt
			}
		}

		res <- ret.PriceRules

		if len(ret.PriceRules) < ret.PageSize || cur == "" {
			break
		}

		cursor = cur
	}

	if cl.Backend != nil {
		if err := cl.Backend.Set(ctx, key, cl.ID(), min.Format(time.RFC3339)); err != nil {
			return fmt.Errorf("failed to store state to backend: %w", err)
		}
	}

	return nil
}

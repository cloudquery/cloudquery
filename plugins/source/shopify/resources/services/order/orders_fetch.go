package order

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchOrders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	const key = "orders"

	p := url.Values{}
	p.Set("status", "any")
	min := time.Time{}

	if cl.Backend != nil {
		value, err := cl.Backend.GetKey(ctx, key)
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
		ret, cur, err := cl.Services.GetOrders(ctx, cursor, p)
		if err != nil {
			return err
		}

		for i := range ret.Orders {
			ts := ret.Orders[i].CreatedAt
			if ret.Orders[i].UpdatedAt != nil {
				ts = *ret.Orders[i].UpdatedAt
			}
			if ts.After(min) {
				min = ts
			}
		}

		res <- ret.Orders

		if len(ret.Orders) < ret.PageSize || cur == "" {
			break
		}

		cursor = cur
		p = nil
	}

	if cl.Backend != nil {
		if err := cl.Backend.SetKey(ctx, key, min.Format(time.RFC3339)); err != nil {
			return fmt.Errorf("failed to store state to backend: %w", err)
		}
	}

	return nil
}

package checkout

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAbandonedCheckouts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	cursor := ""
	for {
		ret, cur, err := cl.Services.GetAbandonedCheckouts(ctx, cursor)
		if err != nil {
			return err
		}
		res <- ret.Checkouts

		if len(ret.Checkouts) < ret.PageSize || cur == "" {
			return nil
		}

		cursor = cur
	}
}

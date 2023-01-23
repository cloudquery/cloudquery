package domain

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	const key = "domains"
	pg, err := cl.GetPaginator(ctx, key)
	if err != nil {
		return err
	}

	for {
		list, p, err := cl.Services.ListDomains(ctx, &pg)
		if err != nil {
			return err
		}
		res <- list

		if p.Next == nil {
			break
		}
		pg.Next = p.Next

		if err := cl.SavePaginator(ctx, key, pg); err != nil {
			return err
		}
	}
	return nil
}

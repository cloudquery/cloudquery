package team

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchTeams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	const key = "teams"
	pg, err := cl.GetPaginator(ctx, key)
	if err != nil {
		return err
	}

	for {
		list, p, err := cl.Services.ListTeams(ctx, &pg)
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

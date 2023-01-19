package team

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchTeamMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	team := parent.Item.(vercel.Team)
	cl := meta.(*client.Client)

	const key = "team_members"
	pg, err := cl.GetPaginator(ctx, key, team.ID)
	if err != nil {
		return err
	}

	for {
		list, p, err := cl.Services.ListTeamMembers(ctx, team.ID, &pg)
		if err != nil {
			return err
		}
		res <- list

		if p.Next == nil {
			break
		}
		pg.Next = p.Next

		if err := cl.SavePaginator(ctx, key, pg, team.ID); err != nil {
			return err
		}
	}
	return nil
}

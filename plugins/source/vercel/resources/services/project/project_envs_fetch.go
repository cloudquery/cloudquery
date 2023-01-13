package project

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchProjectEnvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	proj := parent.Item.(vercel.Project)
	cl := meta.(*client.Client)

	const key = "project_envs"
	pg, err := cl.GetPaginator(ctx, key, proj.ID)
	if err != nil {
		return err
	}

	for {
		list, p, err := cl.Services.ListProjectEnvs(ctx, proj.ID, &pg)
		if err != nil {
			return err
		}
		res <- list

		if p.Next == nil {
			break
		}
		pg.Next = p.Next

		if err := cl.SavePaginator(ctx, key, pg, proj.ID); err != nil {
			return err
		}
	}
	return nil
}

package deployment

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDeploymentChecks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	d := parent.Item.(vercel.Deployment)
	cl := meta.(*client.Client)

	const key = "deployment_checks"
	pg, err := cl.GetPaginator(ctx, key, d.UID)
	if err != nil {
		return err
	}

	for {
		list, p, err := cl.Services.ListDeploymentChecks(ctx, d.UID, &pg)
		if err != nil {
			return err
		}
		res <- list

		if p.Next == nil {
			break
		}
		pg.Next = p.Next

		if err := cl.SavePaginator(ctx, key, pg, d.UID); err != nil {
			return err
		}
	}
	return nil
}

package domain

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDomainRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	dom := parent.Item.(vercel.Domain)
	cl := meta.(*client.Client)

	const key = "deployment_checks"
	pg, err := cl.GetPaginator(ctx, key, dom.ID)
	if err != nil {
		return err
	}

	for {
		list, p, err := cl.Services.ListDomainRecords(ctx, dom.Name, &pg)
		if err != nil {
			return err
		}
		res <- list

		if p.Next == nil {
			break
		}
		pg.Next = p.Next

		if err := cl.SavePaginator(ctx, key, pg, dom.ID); err != nil {
			return err
		}
	}
	return nil
}

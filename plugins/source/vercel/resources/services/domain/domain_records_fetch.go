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

	var pg vercel.Paginator

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
	}
	return nil
}

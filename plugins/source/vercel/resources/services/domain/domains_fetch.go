package domain

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/domain/model"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)

	var until *int64

	for {
		var (
			list struct {
				Domains    []model.Domain         `json:"domains"`
				Pagination client.VercelPaginator `json:"pagination"`
			}
		)

		err := cl.Services.Request(ctx, model.DomainsURL, until, &list)
		if err != nil {
			return err
		}
		res <- list.Domains

		if list.Pagination.Next == nil {
			break
		}

		until = list.Pagination.Next
	}
	return nil
}

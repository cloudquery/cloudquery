package domain

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/domain/model"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDomainRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	dom := parent.Item.(model.Domain)
	u := fmt.Sprintf(model.DomainRecordsURL, dom.Name)

	cl := meta.(*client.Client)

	var until *int64

	for {
		var (
			list struct {
				Records    []model.DomainRecord   `json:"records"`
				Pagination client.VercelPaginator `json:"pagination"`
			}
		)

		err := cl.Services.Request(ctx, u, until, &list)
		if err != nil {
			return err
		}
		res <- list.Records

		if list.Pagination.Next == nil {
			break
		}

		until = list.Pagination.Next
	}
	return nil
}

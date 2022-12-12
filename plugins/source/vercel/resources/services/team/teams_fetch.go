package team

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/team/model"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchTeams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)

	var until *int64

	for {
		var (
			list struct {
				Teams      []model.Team           `json:"teams"`
				Pagination client.VercelPaginator `json:"pagination"`
			}
		)

		err := cl.Services.Request(ctx, model.TeamsURL, until, &list)
		if err != nil {
			return err
		}
		res <- list.Teams

		if list.Pagination.Next == nil {
			break
		}

		until = list.Pagination.Next
	}
	return nil
}

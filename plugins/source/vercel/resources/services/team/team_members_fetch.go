package team

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/team/model"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchTeamMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	team := parent.Item.(model.Team)
	u := fmt.Sprintf(model.TeamMembersURL, team.ID)

	var until *int64

	for {
		var (
			list struct {
				TeamMembers []model.TeamMember     `json:"members"`
				Pagination  client.VercelPaginator `json:"pagination"`
			}
		)

		err := cl.Services.Request(ctx, u, until, &list)
		if err != nil {
			return err
		}
		res <- list.TeamMembers

		if list.Pagination.Next == nil {
			break
		}

		until = list.Pagination.Next
	}
	return nil
}

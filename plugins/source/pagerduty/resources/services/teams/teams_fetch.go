package teams

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchTeams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListTeamsWithContext(ctx, pagerduty.ListTeamOptions{
			Limit:  client.MaxPaginationLimit,
			Offset: offset,
		})
		if err != nil {
			return err
		}

		if len(response.Teams) == 0 {
			return nil
		}

		res <- response.Teams

		offset += uint(len(response.Teams))
		more = response.More
	}

	return nil
}

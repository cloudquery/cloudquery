package users

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListUsersWithContext(ctx, pagerduty.ListUsersOptions{
			Limit:   client.MaxPaginationLimit,
			Offset:  offset,
			TeamIDs: cqClient.Spec.TeamIds,
		})
		if err != nil {
			return err
		}

		if len(response.Users) == 0 {
			return nil
		}

		res <- response.Users

		offset += uint(len(response.Users))
		more = response.More
	}

	return nil
}

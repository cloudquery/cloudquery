package users

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchUsersUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	users, err := meta.(*client.Client).Slack.GetUsersContext(ctx)
	if err != nil {
		return err
	}
	res <- users
	return nil
}

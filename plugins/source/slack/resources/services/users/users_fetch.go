package users

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		users, err := c.Slack.GetUsersContext(ctx)
		if err != nil {
			return err
		}
		res <- users
		return nil
	}
	return c.RetryOnError(ctx, "slack_users", f)
}

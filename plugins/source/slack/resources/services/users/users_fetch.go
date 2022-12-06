package users

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	users, err := meta.(*client.Client).Slack.GetUsersContext(ctx)
	if err != nil {
		return err
	}
	res <- users
	return nil
}

func fetchUserPresences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	userID := parent.Item.(slack.User).ID
	identities, err := meta.(*client.Client).Slack.GetUserPresenceContext(ctx, userID)
	if err != nil {
		return err
	}
	res <- identities
	return nil
}

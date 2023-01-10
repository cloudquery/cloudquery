package users

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchUserPresences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		userID := parent.Item.(slack.User).ID
		identities, err := c.Slack.GetUserPresenceContext(ctx, userID)
		if err != nil {
			return err
		}
		res <- identities
		return nil
	}
	return c.RetryOnError(ctx, "slack_user_presences", f)
}

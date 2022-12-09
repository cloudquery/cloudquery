package conversations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchConversations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	teamID := c.TeamID
	params := &slack.GetConversationsParameters{
		TeamID: teamID,
		Types:  []string{"public_channel", "private_channel", "mpim", "im"},
		Limit:  1000,
	}
	for {
		conversations, nextCursor, err := c.Slack.GetConversationsContext(ctx, params)
		c.Logger().Debug().Err(err).Interface("conversations", conversations).Str("next_cursor", nextCursor).Msg("Got conversations")
		if err != nil {
			return err
		}
		res <- conversations
		if nextCursor == "" {
			break
		}
		params.Cursor = nextCursor
	}
	return nil
}

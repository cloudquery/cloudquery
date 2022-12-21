package conversations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/services/conversations/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchConversationMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		channelID := parent.Item.(slack.Channel).ID
		params := &slack.GetUsersInConversationParameters{
			ChannelID: channelID,
			Limit:     1000,
		}
		for {
			ids, nextCursor, err := c.Slack.GetUsersInConversationContext(ctx, params)
			if err != nil {
				return err
			}
			for _, id := range ids {
				res <- models.ConversationMember{
					UserID:    id,
					ChannelID: channelID,
				}
			}
			if nextCursor == "" {
				break
			}
			params.Cursor = nextCursor
		}
		return nil
	}
	return c.RetryOnError(ctx, "slack_conversation_members", f)
}

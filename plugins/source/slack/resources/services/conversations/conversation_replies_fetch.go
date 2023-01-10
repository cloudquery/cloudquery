package conversations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchConversationReplies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		params := &slack.GetConversationRepliesParameters{
			ChannelID:          parent.Get("channel_id").String(),
			Timestamp:          parent.Item.(slack.Msg).Timestamp,
			Limit:              1000,
			IncludeAllMetadata: true,
		}
		if parent.Item.(slack.Msg).LatestReply == "" {
			// optimization: return early if parent is known to not have any replies
			return nil
		}

		for {
			resp, hasMore, nextCursor, err := c.Slack.GetConversationRepliesContext(ctx, params)
			if err != nil {
				return err
			}
			for _, m := range resp {
				res <- m.Msg
			}
			if !hasMore {
				break
			}
			params.Cursor = nextCursor
		}
		return nil
	}
	return c.RetryOnError(ctx, "slack_conversation_replies", f)
}

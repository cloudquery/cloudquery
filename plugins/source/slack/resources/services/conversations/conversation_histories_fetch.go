package conversations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchConversationHistories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		channel := parent.Item.(slack.Channel)
		params := &slack.GetConversationHistoryParameters{
			ChannelID:          channel.ID,
			Limit:              1000,
			IncludeAllMetadata: true,
		}
		for {
			resp, err := c.Slack.GetConversationHistoryContext(ctx, params)
			if err != nil {
				if isNotInChannel(err) {
					// we do not expect to fetch conversation histories from channels the bot is not
					// added to
					return nil
				}
				return err
			}
			if resp.Err() != nil {
				return resp.Err()
			}
			for _, m := range resp.Messages {
				res <- m.Msg
			}
			if !resp.HasMore {
				break
			}
			params.Cursor = resp.ResponseMetaData.NextCursor
		}
		return nil
	}
	return c.RetryOnError(ctx, "slack_conversation_histories", f)
}

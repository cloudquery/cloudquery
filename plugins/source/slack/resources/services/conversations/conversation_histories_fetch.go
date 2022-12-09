package conversations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchConversationHistories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	params := &slack.GetConversationHistoryParameters{
		ChannelID:          parent.Item.(slack.Channel).ID,
		Limit:              1000,
		IncludeAllMetadata: true,
	}
	for {
		resp, err := c.Slack.GetConversationHistoryContext(ctx, params)
		if err != nil {
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

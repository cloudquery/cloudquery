package groups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchGroupConversations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	teamID := meta.(*client.Client).TeamID
	params := &slack.GetConversationsParameters{
		Cursor: "",
		TeamID: teamID,
		Types:  []string{"public_channel", "private_channel", "mpim", "im"},
		Limit:  1000,
	}

	for {
		conversations, nextCursor, err := meta.(*client.Client).Slack.GetConversationsContext(ctx, params)
		if err != nil {
			return err
		}
		if nextCursor == "" {
			break
		}
		res <- conversations
		params.Cursor = nextCursor
	}
	return nil
}

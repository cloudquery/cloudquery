package conversations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchConversationBookmarks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		channelID := parent.Item.(slack.Channel).ID
		bookmarks, err := c.Slack.ListBookmarksContext(ctx, channelID)
		if err != nil {
			if isNotInChannel(err) {
				// we do not expect to fetch bookmarks from channels the bot is not
				// added to
				return nil
			}
			return err
		}
		res <- bookmarks
		return nil
	}
	return c.RetryOnError(ctx, "slack_conversation_bookmarks", f)
}

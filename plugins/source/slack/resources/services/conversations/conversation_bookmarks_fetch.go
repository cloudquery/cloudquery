package conversations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchConversationBookmarks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	f := func() error {
		channelID := parent.Get("id").String()
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
	return c.RetryOnRateLimitError("slack_conversation_bookmarks", f)
}

package conversations

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/slack-go/slack"
)

func ConversationHistories() *schema.Table {
	return &schema.Table{
		Name:        "slack_conversation_histories",
		Description: `https://api.slack.com/methods/conversations.history`,
		Resolver:    fetchConversationHistories,
		Multiplex:   client.TeamMultiplex,
		Transform:   transformers.TransformWithStruct(&slack.Msg{}, transformers.WithSkipFields("Blocks", "Replies")),
		Columns: []schema.Column{
			{
				Name:       "channel_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
			{
				Name:       "team_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("team_id"),
				PrimaryKey: true,
			},
			{
				Name:       "ts",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Timestamp"),
				PrimaryKey: true,
			},
			{
				Name:     "thread_ts",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("ThreadTimestamp"),
			},
		},

		Relations: []*schema.Table{
			ConversationReplies(),
		},
	}
}

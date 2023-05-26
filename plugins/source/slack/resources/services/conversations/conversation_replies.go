package conversations

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/slack-go/slack"
)

func ConversationReplies() *schema.Table {
	return &schema.Table{
		Name:        "slack_conversation_replies",
		Description: `https://api.slack.com/methods/conversations.replies`,
		Resolver:    fetchConversationReplies,
		Multiplex:   client.TeamMultiplex,
		Transform:   transformers.TransformWithStruct(&slack.Msg{}, transformers.WithSkipFields("Blocks")),
		Columns: []schema.Column{
			{
				Name:     "conversation_history_ts",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("ts"),
			},
			{
				Name:       "team_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("team_id"),
				PrimaryKey: true,
			},
			{
				Name:       "channel_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("channel_id"),
				PrimaryKey: true,
			},
			{
				Name:       "ts",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Timestamp"),
				PrimaryKey: true,
			},
		},
	}
}

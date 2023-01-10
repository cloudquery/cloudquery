package conversations

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("ts"),
			},
			{
				Name:     "team_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("team_id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "channel_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("channel_id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "ts",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Timestamp"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

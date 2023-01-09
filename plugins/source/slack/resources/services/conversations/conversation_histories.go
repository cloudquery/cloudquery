package conversations

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "channel_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
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
				Name:     "ts",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Timestamp"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "thread_ts",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThreadTimestamp"),
			},
		},

		Relations: []*schema.Table{
			ConversationReplies(),
		},
	}
}

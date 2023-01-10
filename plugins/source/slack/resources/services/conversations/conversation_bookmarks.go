package conversations

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/slack-go/slack"
)

func ConversationBookmarks() *schema.Table {
	return &schema.Table{
		Name:        "slack_conversation_bookmarks",
		Description: `https://api.slack.com/methods/bookmarks.list`,
		Resolver:    fetchConversationBookmarks,
		Multiplex:   client.TeamMultiplex,
		Transform:   transformers.TransformWithStruct(&slack.Bookmark{}),
		Columns: []schema.Column{
			{
				Name:     "team_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveTeamID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "channel_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ChannelID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "date_created",
				Type:     schema.TypeTimestamp,
				Resolver: client.JSONTimeResolver("Created"),
			},
			{
				Name:     "date_updated",
				Type:     schema.TypeTimestamp,
				Resolver: client.JSONTimeResolver("Updated"),
			},
		},
	}
}

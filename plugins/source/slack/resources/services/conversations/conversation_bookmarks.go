package conversations

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
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
				Name:       "team_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveTeamID,
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:       "channel_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ChannelID"),
				PrimaryKey: true,
			},
			{
				Name:     "date_created",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: client.JSONTimeResolver("Created"),
			},
			{
				Name:     "date_updated",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: client.JSONTimeResolver("Updated"),
			},
		},
	}
}

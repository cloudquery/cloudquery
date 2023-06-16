package conversations

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/services/conversations/models"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Conversations() *schema.Table {
	return &schema.Table{
		Name:        "slack_conversations",
		Description: `https://api.slack.com/methods/conversations.list`,
		Resolver:    fetchConversations,
		Multiplex:   client.TeamMultiplex,
		Transform: transformers.TransformWithStruct(
			&models.Conversation{},
			transformers.WithSkipFields("Members"),
		),
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
				Name:     "created",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: client.JSONTimeResolver("Created"),
			},
		},

		Relations: []*schema.Table{
			ConversationBookmarks(),
			ConversationHistories(),
			ConversationMembers(),
		},
	}
}

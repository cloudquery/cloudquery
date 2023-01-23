package conversations

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/services/conversations/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "created",
				Type:     schema.TypeTimestamp,
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

package conversations

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/services/conversations/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConversationMembers() *schema.Table {
	return &schema.Table{
		Name:        "slack_conversation_members",
		Description: `https://api.slack.com/methods/conversations.members`,
		Resolver:    fetchConversationMembers,
		Multiplex:   client.TeamMultiplex,
		Transform:   transformers.TransformWithStruct(&models.ConversationMember{}),
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
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserID"),
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
		},
	}
}

package conversations

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/services/conversations/models"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
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
				Name:       "team_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveTeamID,
				PrimaryKey: true,
			},
			{
				Name:       "user_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("UserID"),
				PrimaryKey: true,
			},
			{
				Name:       "channel_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ChannelID"),
				PrimaryKey: true,
			},
		},
	}
}

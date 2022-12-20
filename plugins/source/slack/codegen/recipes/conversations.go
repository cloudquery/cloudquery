package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/services/conversations/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func ConversationResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "conversations",
			Description: "https://api.slack.com/methods/conversations.list",
			DataStruct:  &models.Conversation{},
			SkipFields: []string{
				"Members", // part of model, but not returned by API
			},
			PKColumns: []string{"team_id", "id"},
			Relations: []string{
				`ConversationBookmarks()`,
				`ConversationHistories()`,
				`ConversationMembers()`,
			},
			ExtraColumns: []codegen.ColumnDefinition{TeamIDColumn},
		},
		{
			SubService:   "conversation_members",
			Description:  "https://api.slack.com/methods/conversations.members",
			DataStruct:   &models.ConversationMember{},
			PKColumns:    []string{"team_id", "channel_id", "user_id"},
			Relations:    []string{},
			ExtraColumns: []codegen.ColumnDefinition{TeamIDColumn},
		},
		{
			SubService:   "conversation_bookmarks",
			Description:  "https://api.slack.com/methods/bookmarks.list",
			DataStruct:   &slack.Bookmark{},
			PKColumns:    []string{"team_id", "channel_id", "id"},
			ExtraColumns: []codegen.ColumnDefinition{TeamIDColumn},
		},
		{
			SubService:  "conversation_histories",
			Description: "https://api.slack.com/methods/conversations.history",
			DataStruct:  &slack.Msg{},
			PKColumns:   []string{"team_id", "channel_id", "ts"},
			SkipFields: []string{
				"Team",    // empty in API response
				"Blocks",  // empty in API response
				"Replies", // empty in API response
				"Channel", // empty in API response
			},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "channel_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("id")`,
				},
				{
					Name:     "team_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("team_id")`,
				},
			},
			Relations: []string{
				`ConversationReplies()`,
			},
		},
		{
			SubService:  "conversation_replies",
			Description: "https://api.slack.com/methods/conversations.replies",
			DataStruct:  &slack.Msg{},
			SkipFields:  []string{"Team", "Blocks"},
			PKColumns:   []string{"team_id", "channel_id", "ts"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "conversation_history_ts",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("ts")`,
				},
				{
					Name:     "team_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("team_id")`,
				},
				{
					Name:     "channel_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("channel_id")`,
				},
			},
		},
	}
	for _, r := range resources {
		r.Service = "conversations"
		r.Multiplex = `client.TeamMultiplex`
	}
	return resources
}

package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func UserResources() []*Resource {
	resources := []*Resource{
		{
			DataStruct:  &slack.User{},
			Description: "https://api.slack.com/methods/users.list",
			PKColumns:   []string{"id"},
			SkipFields: []string{
				"Has2FA",   // not returned for bot tokens; skipping to avoid confusion
				"Presence", // not returned by this API, but can be found in slack_user_presences table
			},
			Relations: []string{
				`UserPresences()`,
			},
		},
		{
			PKColumns:   []string{"user_id"},
			DataStruct:  &slack.UserPresence{},
			Description: "https://api.slack.com/methods/users.getPresence",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "user_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("id")`,
				},
			},
		},
	}
	for _, r := range resources {
		r.Service = "users"
	}
	return resources
}

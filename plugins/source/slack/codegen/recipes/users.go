package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func UserResources() []*Resource {
	resources := []*Resource{
		{
			DataStruct: &slack.User{},
			PKColumns:  []string{"id"},
			Relations: []string{
				`UserPresences()`,
			},
		},
		{
			PKColumns:  []string{"user_id"},
			DataStruct: &slack.UserPresence{},
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

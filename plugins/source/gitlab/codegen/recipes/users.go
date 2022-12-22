package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/xanzy/go-gitlab"
)

func Users() []*Resource {
	resources := []*Resource{
		{
			Service:    "users",
			SubService: "users",
			PKColumns:  []string{"base_url", "id"},
			Struct:     &gitlab.User{},
			SkipFields: []string{"LastActivityOn"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:          "last_activity_on",
					Type:          schema.TypeJSON,
					Resolver:      `schema.PathResolver("LastActivityOn")`,
					IgnoreInTests: true,
				},
			},
		},
	}

	return resources
}

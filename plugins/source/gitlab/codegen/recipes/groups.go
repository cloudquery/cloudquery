package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/xanzy/go-gitlab"
)

func Groups() []*Resource {
	resources := []*Resource{
		{
			Service:    "groups",
			SubService: "groups",
			PKColumns:  []string{"base_url", "id", "name"},
			Struct:     &gitlab.Group{},
			Relations:  []string{"GroupMembers()"},
		},
		{
			Service:    "groups",
			SubService: "group_members",
			Struct:     &gitlab.GroupMember{},
			PKColumns:  []string{"base_url", "group_id", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "group_id",
					Type:     schema.TypeInt,
					Resolver: `resolveGroupID`,
				},
			},
		},
	}

	return resources
}

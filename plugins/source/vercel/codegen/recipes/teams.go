package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TeamResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &vercel.Team{},
			Service:    "team",
			PKColumns:  []string{"id"},
			Relations:  []string{"TeamMembers()"},
		},
		{
			DataStruct: &vercel.TeamMember{},
			Service:    "team",
			PKColumns:  []string{"team_id", "uid"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "team_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("id")`,
				},
			},
		},
	}
}

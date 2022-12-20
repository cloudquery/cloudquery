package recipes

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TeamResources() []*Resource {
	return []*Resource{
		{
			SubService:  "teams",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.Team{},
			Description: "https://developer.pagerduty.com/api-reference/0138639504311-list-teams",

			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "members",
					Type:     schema.TypeJSON,
					Resolver: `MembersResolver`,
				},
			},

			ListOptionsStructNameOverride: "ListTeamOptions",
			ResponseStructOverride:        "ListTeamResponse",
		},
	}
}

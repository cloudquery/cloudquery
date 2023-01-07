package teams

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_teams",
		Description: `https://developer.pagerduty.com/api-reference/0138639504311-list-teams`,
		Resolver:    fetchTeams,
		Transform:   transformers.TransformWithStruct(&pagerduty.Team{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			{
				Name:     "members",
				Type:     schema.TypeJSON,
				Resolver: MembersResolver,
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
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
		},
	}
}

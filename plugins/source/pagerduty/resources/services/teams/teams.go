package teams

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
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
				Type:     types.ExtensionTypes.JSON,
				Resolver: MembersResolver,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "html_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("HTMLURL"),
			},
		},
	}
}

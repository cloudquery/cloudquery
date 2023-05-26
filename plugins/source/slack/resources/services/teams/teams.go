package teams

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/slack-go/slack"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:        "slack_teams",
		Description: `https://slack.com/api/team.info`,
		Resolver:    fetchTeams,
		Multiplex:   client.TeamMultiplex,
		Transform:   transformers.TransformWithStruct(&slack.TeamInfo{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}

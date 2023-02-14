package teams

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:      "github_teams",
		Resolver:  fetchTeams,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.Team{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns:   []schema.Column{client.OrgColumn},
		Relations: []*schema.Table{members(), repositories()},
	}
}

var teamIDColumn = schema.Column{
	Name:            "team_id",
	Type:            schema.TypeInt,
	Resolver:        client.ResolveParentColumn("ID"),
	CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
}

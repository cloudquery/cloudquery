package teams

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:      "github_team_repositories",
		Resolver:  fetchRepositories,
		Transform: transformers.TransformWithStruct(&github.Repository{}, client.SharedTransformers()...),
		Columns: []schema.Column{
			{
				Name:        "org",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
				Description: `The Github Organization of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "team_id",
				Type:     schema.TypeInt,
				Resolver: client.ResolveParentColumn("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

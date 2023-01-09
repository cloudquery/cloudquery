package repositories

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Alerts() *schema.Table {
	return &schema.Table{
		Name:      "github_repository_dependabot_alerts",
		Resolver:  fetchAlerts,
		Transform: transformers.TransformWithStruct(&github.DependabotAlert{}, client.SharedTransformers()...),
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
				Name:     "repository_id",
				Type:     schema.TypeInt,
				Resolver: client.ResolveParentColumn("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Number"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

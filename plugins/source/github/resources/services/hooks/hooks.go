package hooks

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Hooks() *schema.Table {
	return &schema.Table{
		Name:      "github_hooks",
		Resolver:  fetchHooks,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.Hook{}, client.SharedTransformers()...),
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
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			Deliveries(),
		},
	}
}

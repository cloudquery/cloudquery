package team

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:      "vercel_teams",
		Resolver:  fetchTeams,
		Transform: transformers.TransformWithStruct(&vercel.Team{}, client.SharedTransformers()...),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			TeamMembers(),
		},
	}
}

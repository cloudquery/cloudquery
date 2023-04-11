package project

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:          "vercel_projects",
		Resolver:      fetchProjects,
		Transform:     client.TransformWithStruct(&vercel.Project{}),
		Multiplex:     client.TeamMultiplex,
		IsIncremental: true,
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
			ProjectEnvs(),
		},
	}
}

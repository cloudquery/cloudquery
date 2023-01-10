package projects

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_projects",
		Resolver:  fetchProjectsProjects,
		Transform: transformers.TransformWithStruct(&godo.Project{}),
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
			Resources(),
		},
	}
}

package projects

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
			resources(),
		},
	}
}

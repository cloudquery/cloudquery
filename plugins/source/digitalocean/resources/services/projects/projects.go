package projects

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_projects",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Projects",
		Resolver:    fetchProjectsProjects,
		Transform:   transformers.TransformWithStruct(&godo.Project{}),
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

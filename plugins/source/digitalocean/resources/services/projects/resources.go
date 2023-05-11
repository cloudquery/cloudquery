package projects

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func resources() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_project_resources",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Project-Resources",
		Resolver:    fetchProjectsResources,
		Transform:   transformers.TransformWithStruct(&godo.ProjectResource{}),
		Columns: []schema.Column{
			{
				Name:     "urn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

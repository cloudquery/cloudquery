package projects

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func resources() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_project_resources",
		Resolver:  fetchProjectsResources,
		Transform: transformers.TransformWithStruct(&godo.ProjectResource{}),
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

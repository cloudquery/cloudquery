package projects

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Resources() *schema.Table {
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

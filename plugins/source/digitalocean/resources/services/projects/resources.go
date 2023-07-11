package projects

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "urn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("URN"),
				PrimaryKey: true,
			},
		},
	}
}

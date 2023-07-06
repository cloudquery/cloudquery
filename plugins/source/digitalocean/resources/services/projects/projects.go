package projects

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			resources(),
		},
	}
}

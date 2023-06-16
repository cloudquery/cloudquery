package project

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/v3/schema"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			ProjectEnvs(),
		},
	}
}

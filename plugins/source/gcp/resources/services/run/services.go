package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "gcp_run_services",
		Description: `https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#Service`,
		Resolver:    fetchServices,
		Multiplex:   client.ProjectMultiplexEnabledServices("run.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Service{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

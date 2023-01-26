package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "gcp_run_services",
		Description: `https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#Service`,
		Resolver:    fetchServices,
		Multiplex:   client.ProjectMultiplexEnabledServices("run.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Service{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
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

package cloudscheduler

import (
	pb "cloud.google.com/go/scheduler/apiv1/schedulerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Jobs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudscheduler_jobs",
		Description: `https://cloud.google.com/scheduler/docs/reference/rest/v1/projects.locations.jobs#Job`,
		Resolver:    fetchJobs,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudscheduler.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Job{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

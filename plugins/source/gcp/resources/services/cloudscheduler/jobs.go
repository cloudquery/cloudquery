package cloudscheduler

import (
	pb "cloud.google.com/go/scheduler/apiv1/schedulerpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Jobs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudscheduler_jobs",
		Description: `https://cloud.google.com/scheduler/docs/reference/rest/v1/projects.locations.jobs#Job`,
		Resolver:    fetchJobs,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudscheduler.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Job{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}

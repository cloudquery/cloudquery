package videotranscoder

import (
	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Jobs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_videotranscoder_jobs",
		Description: `https://cloud.google.com/transcoder/docs/reference/rest/v1/projects.locations.jobs`,
		Resolver:    fetchJobs,
		Multiplex:   client.ProjectMultiplexEnabledServices("transcoder.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Job{}, transformers.WithPrimaryKeys("Name")),
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

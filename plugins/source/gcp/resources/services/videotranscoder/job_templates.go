package videotranscoder

import (
	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func JobTemplates() *schema.Table {
	return &schema.Table{
		Name:        "gcp_videotranscoder_job_templates",
		Description: `https://cloud.google.com/transcoder/docs/reference/rest/v1/projects.locations.jobTemplates`,
		Resolver:    fetchJobTemplates,
		Multiplex:   client.ProjectMultiplexEnabledServices("transcoder.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.JobTemplate{}, transformers.WithPrimaryKeys("Name")),
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

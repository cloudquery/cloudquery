package videotranscoder

import (
	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func JobTemplates() *schema.Table {
	return &schema.Table{
		Name:        "gcp_videotranscoder_job_templates",
		Description: `https://cloud.google.com/transcoder/docs/reference/rest/v1/projects.locations.jobTemplates`,
		Resolver:    fetchJobTemplates,
		Multiplex:   client.ProjectMultiplexEnabledServices("transcoder.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.JobTemplate{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
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

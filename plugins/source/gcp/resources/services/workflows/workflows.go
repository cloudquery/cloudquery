package workflows

import (
	pb "cloud.google.com/go/workflows/apiv1/workflowspb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Workflows() *schema.Table {
	return &schema.Table{
		Name:        "gcp_workflows_workflows",
		Description: `https://cloud.google.com/workflows/docs/reference/rest/v1/projects.locations.workflows#resource:-workflow`,
		Resolver:    fetchWorkflows,
		Multiplex:   client.ProjectMultiplexEnabledServices("workflows.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Workflow{}, transformers.WithPrimaryKeys("Name")),
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

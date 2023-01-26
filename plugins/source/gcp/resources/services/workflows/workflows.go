package workflows

import (
	pb "cloud.google.com/go/workflows/apiv1/workflowspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Workflows() *schema.Table {
	return &schema.Table{
		Name:        "gcp_workflows_workflows",
		Description: `https://cloud.google.com/workflows/docs/reference/rest/v1/projects.locations.workflows#resource:-workflow`,
		Resolver:    fetchWorkflows,
		Multiplex:   client.ProjectMultiplexEnabledServices("workflows.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Workflow{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
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

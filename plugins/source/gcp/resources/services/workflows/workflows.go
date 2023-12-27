package workflows

import (
	pb "cloud.google.com/go/workflows/apiv1/workflowspb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Workflows() *schema.Table {
	return &schema.Table{
		Name:        "gcp_workflows_workflows",
		Description: `https://cloud.google.com/workflows/docs/reference/rest/v1/projects.locations.workflows#resource:-workflow`,
		Resolver:    fetchWorkflows,
		Multiplex:   client.ProjectMultiplexEnabledServices("workflows.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Workflow{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			client.ProjectIDColumn(true),
		},
	}
}

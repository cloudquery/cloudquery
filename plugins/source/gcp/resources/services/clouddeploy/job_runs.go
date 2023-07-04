package clouddeploy

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	deploy "cloud.google.com/go/deploy/apiv1"
)

func JobRuns() *schema.Table {
	return &schema.Table{
		Name:        "gcp_clouddeploy_job_runs",
		Description: `https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases.rollouts.jobRuns#JobRun`,
		Resolver:    fetchJobRuns,
		Multiplex:   client.ProjectMultiplexEnabledServices("clouddeploy.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.JobRun{}, transformers.WithPrimaryKeys("Name")),
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

func fetchJobRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListJobRunsRequest{
		Parent: parent.Item.(*pb.Rollout).Name,
	}
	gcpClient, err := deploy.NewCloudDeployClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListJobRuns(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}

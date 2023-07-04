package batch

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/batch/apiv1/batchpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	batch "cloud.google.com/go/batch/apiv1"
)

func TaskGroups() *schema.Table {
	return &schema.Table{
		Name:        "gcp_batch_task_groups",
		Description: `https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#TaskGroup`,
		Resolver:    fetchTaskGroups,
		Multiplex:   client.ProjectMultiplexEnabledServices("batch.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.TaskGroup{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			Tasks(),
		},
	}
}

func fetchTaskGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListJobsRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := batch.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListJobs(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp.TaskGroups
	}
	return nil
}

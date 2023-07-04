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

func Tasks() *schema.Table {
	return &schema.Table{
		Name:        "gcp_batch_tasks",
		Description: `https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs.taskGroups.tasks/list`,
		Resolver:    fetchTasks,
		Multiplex:   client.ProjectMultiplexEnabledServices("batch.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Task{}, transformers.WithPrimaryKeys("Name")),
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

func fetchTasks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListTasksRequest{
		Parent: parent.Item.(*pb.TaskGroup).Name,
	}
	gcpClient, err := batch.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListTasks(ctx, req, c.CallOptions...)
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

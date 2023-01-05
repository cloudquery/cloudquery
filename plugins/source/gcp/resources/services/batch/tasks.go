package batch

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/batch/apiv1/batchpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	batch "cloud.google.com/go/batch/apiv1"
)

func Tasks() *schema.Table {
	return &schema.Table{
		Name:        "gcp_batch_tasks",
		Description: `https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs.taskGroups.tasks/list`,
		Resolver:    fetchTasks,
		Multiplex:   client.ProjectMultiplexEnabledServices("batch.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Task{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
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

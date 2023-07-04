package workflows

import (
	"context"

	workflows "cloud.google.com/go/workflows/apiv1"
	pb "cloud.google.com/go/workflows/apiv1/workflowspb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func fetchWorkflows(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	gcpClient, err := workflows.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := gcpClient.ListWorkflows(ctx, &pb.ListWorkflowsRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}, c.CallOptions...)
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

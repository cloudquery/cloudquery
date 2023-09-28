package vmmigration

import (
	"context"

	vmmigration "cloud.google.com/go/vmmigration/apiv1"
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"google.golang.org/api/iterator"
)

func fetchCutoverJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	parentItem := parent.Item.(*pb.MigratingVm)

	gcpClient, err := vmmigration.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := gcpClient.ListCutoverJobs(ctx, &pb.ListCutoverJobsRequest{
		Parent: parentItem.Name,
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

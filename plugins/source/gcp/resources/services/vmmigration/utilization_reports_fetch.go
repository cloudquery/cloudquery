package vmmigration

import (
	"context"

	vmmigration "cloud.google.com/go/vmmigration/apiv1"
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func fetchUtilizationReports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	parentItem := parent.Item.(*pb.Source)

	gcpClient, err := vmmigration.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := gcpClient.ListUtilizationReports(ctx, &pb.ListUtilizationReportsRequest{
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

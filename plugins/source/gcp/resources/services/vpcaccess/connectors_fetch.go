package vpcaccess

import (
	"context"

	vpcaccess "cloud.google.com/go/vpcaccess/apiv1"
	pb "cloud.google.com/go/vpcaccess/apiv1/vpcaccesspb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func fetchConnectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	parentItem := parent.Item.(*locationpb.Location)

	gcpClient, err := vpcaccess.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := gcpClient.ListConnectors(ctx, &pb.ListConnectorsRequest{
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

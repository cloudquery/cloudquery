package run

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/run/apiv2/runpb"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	v2 "cloud.google.com/go/run/apiv2"
	v1 "google.golang.org/api/run/v1"
)

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListServicesRequest{
		Parent: "projects/" + c.ProjectId + "/locations/" + parent.Item.(*v1.Location).LocationId,
	}
	gcpClient, err := v2.NewServicesClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListServices(ctx, req, c.CallOptions...)
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

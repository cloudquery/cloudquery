package vision

import (
	"context"

	vision "cloud.google.com/go/vision/v2/apiv1"
	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func fetchProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	gcpClient, err := vision.NewProductSearchClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := gcpClient.ListProducts(ctx, &pb.ListProductsRequest{
		Parent:   "projects/" + c.ProjectId + "/locations/" + c.Location,
		PageSize: 100,
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

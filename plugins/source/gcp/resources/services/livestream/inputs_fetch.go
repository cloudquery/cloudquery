package livestream

import (
	"context"

	livestream "cloud.google.com/go/video/livestream/apiv1"
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func fetchInputs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	gcpClient, err := livestream.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := gcpClient.ListInputs(ctx, &pb.ListInputsRequest{
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

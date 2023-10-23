package videotranscoder

import (
	"context"

	transcoder "cloud.google.com/go/video/transcoder/apiv1"
	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"google.golang.org/api/iterator"
)

func fetchJobTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	gcpClient, err := transcoder.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := gcpClient.ListJobTemplates(ctx, &pb.ListJobTemplatesRequest{
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

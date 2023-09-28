package translate

import (
	"context"

	translate "cloud.google.com/go/translate/apiv3"
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"google.golang.org/api/iterator"
)

func fetchGlossaries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	gcpClient, err := translate.NewTranslationClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := gcpClient.ListGlossaries(ctx, &pb.ListGlossariesRequest{
		Parent: "projects/" + c.ProjectId + "/locations/us-central1",
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

package artifactregistry

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"google.golang.org/api/artifactregistry/v1"
)

func fetchLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	gcpClient, err := artifactregistry.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	for {
		output, err := gcpClient.Projects.Locations.List("projects/" + c.ProjectId).PageToken(nextPageToken).Context(ctx).Do()
		if err != nil {
			return err
		}
		res <- output.Locations
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

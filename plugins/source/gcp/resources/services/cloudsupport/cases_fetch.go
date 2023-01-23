package cloudsupport

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	cloudsupport "google.golang.org/api/cloudsupport/v2beta"
)

func fetchCases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	gcpClient, err := cloudsupport.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	for {
		output, err := gcpClient.Cases.List("projects/" + c.ProjectId).PageSize(1000).PageToken(nextPageToken).Context(ctx).Do()
		if err != nil {
			return err
		}
		res <- output.Cases
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

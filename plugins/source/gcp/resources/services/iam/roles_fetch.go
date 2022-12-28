package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iam/v1"
)

func fetchRoles(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	iamClient, err := iam.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	for {
		output, err := iamClient.Projects.Roles.List("projects/" + c.ProjectId).PageSize(1000).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		res <- output.Roles

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

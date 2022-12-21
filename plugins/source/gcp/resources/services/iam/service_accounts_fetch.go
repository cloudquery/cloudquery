package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iam/v1"
)

func fetchServiceAccounts(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	iamClient, err := iam.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	for {
		output, err := iamClient.Projects.ServiceAccounts.List("projects/" + c.ProjectId).PageToken(nextPageToken).PageSize(100).Do()
		if err != nil {
			return err
		}
		res <- output.Accounts

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

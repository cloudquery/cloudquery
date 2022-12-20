package iam

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	iam "google.golang.org/api/iam/v2beta"
)

func fetchDenyPolicies(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	iamClient, err := iam.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	// See https://cloud.google.com/iam/docs/reference/rest/v2beta/policies#Policy for the format of the parent
	parent := fmt.Sprintf("policies/cloudresourcemanager.googleapis.com%%2Fprojects%%2F%s/denypolicies", c.ProjectId)

	for {
		// Not using PageSize here as IAM ignores it and uses 1000 anyway
		output, err := iamClient.Policies.ListPolicies(parent).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		res <- output.Policies

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

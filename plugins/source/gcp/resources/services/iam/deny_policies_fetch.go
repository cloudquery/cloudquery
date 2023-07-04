package iam

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	policies "cloud.google.com/go/iam/apiv2"
	policiespb "cloud.google.com/go/iam/apiv2/iampb"
	"google.golang.org/api/iterator"
)

func fetchDenyPolicies(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	iamClient, err := policies.NewPoliciesClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	iamClient.CallOptions = &policies.PoliciesCallOptions{}

	// See https://cloud.google.com/iam/docs/reference/rest/v2beta/policies#Policy for the format of the parent
	parent := fmt.Sprintf("policies/cloudresourcemanager.googleapis.com%%2Fprojects%%2F%s/denypolicies", c.ProjectId)
	req := &policiespb.ListPoliciesRequest{
		Parent:   parent,
		PageSize: 1000,
	}

	it := iamClient.ListPolicies(ctx, req, c.CallOptions...)

	for {
		policy, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- []*policiespb.Policy{policy}
	}

	return nil
}

package resourcemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func fetchProjectPolicies(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	output, err := c.Services.ResourcemanagerProjectsClient.GetIamPolicy(
		ctx,
		&iampb.GetIamPolicyRequest{
			Resource: "projects/" + c.ProjectId,
		},
	)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- output
	return nil
}

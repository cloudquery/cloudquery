package resourcemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/cloudresourcemanager/v3"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func fetchProjectPolicies(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	p := r.Parent.Item.(*cloudresourcemanager.Project)
	output, err := c.Services.ResourcemanagerProjectsClient.GetIamPolicy(
		ctx,
		&iampb.GetIamPolicyRequest{
			Resource: "projects/" + p.ProjectId,
		},
	)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- output
	return nil
}

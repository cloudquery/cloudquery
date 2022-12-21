package resourcemanager

import (
	"context"

	"cloud.google.com/go/iam/apiv1/iampb"
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func fetchProjectPolicies(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	projectsClient, err := resourcemanager.NewProjectsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	output, err := projectsClient.GetIamPolicy(
		ctx,
		&iampb.GetIamPolicyRequest{
			Resource: "projects/" + c.ProjectId,
		},
	)
	if err != nil {
		return err
	}
	res <- output
	return nil
}

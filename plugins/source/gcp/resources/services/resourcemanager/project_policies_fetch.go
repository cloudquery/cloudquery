package resourcemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/cloudresourcemanager/v3"
)

func fetchProjectPolicies(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	projectsClient, err := pb.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	// We need to use the protobut client to get the current version of the policy struct (v3)
	output, err := projectsClient.Projects.GetIamPolicy("projects/"+c.ProjectId, &pb.GetIamPolicyRequest{}).Context(ctx).Do()
	if err != nil {
		return err
	}
	res <- output
	return nil
}

package resourcemanager

import (
	"context"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &resourcemanagerpb.GetProjectRequest{
		Name: "projects/" + c.ProjectId,
	}
	projectsClient, err := resourcemanager.NewProjectsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	output, err := projectsClient.GetProject(ctx, req, c.CallOptions...)
	if err != nil {
		return err
	}
	res <- output
	return nil
}

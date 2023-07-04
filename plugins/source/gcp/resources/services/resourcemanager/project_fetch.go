package resourcemanager

import (
	"context"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func fetchProjects(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &resourcemanagerpb.GetProjectRequest{
		Name: "projects/" + c.ProjectId,
	}
	projectsClient, err := resourcemanager.NewProjectsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	output, err := projectsClient.GetProject(ctx, req)
	if err != nil {
		return err
	}
	res <- output
	return nil
}

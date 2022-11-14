package resourcemanager

import (
	"context"

	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func fetchProjects(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &resourcemanagerpb.GetProjectRequest{
		Name: "projects/" + c.ProjectId,
	}
	output, err := c.Services.ResourcemanagerProjectsClient.GetProject(ctx, req)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- output
	return nil
}

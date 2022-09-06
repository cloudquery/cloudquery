package resourcemanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	pb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"
)

func fetchProjects(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.GetProjectRequest{
		Name: "projects/" + c.ProjectId,
	}
	output, err := c.Services.ResourcemanagerProjectsClient.GetProject(ctx, req)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- output
	return nil
}

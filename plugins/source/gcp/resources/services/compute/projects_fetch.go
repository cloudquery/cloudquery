package compute

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.GetProjectRequest{
		Project: c.ProjectId,
	}
	computeProjectsClient, err := compute.NewProjectsRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	resp, err := computeProjectsClient.Get(ctx, req)
	if err != nil {
		return err
	}
	res <- resp
	return nil
}

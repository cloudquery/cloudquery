package appengine

import (
	"context"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	appengine "cloud.google.com/go/appengine/apiv1"
)

func fetchApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.GetApplicationRequest{
		Name: "apps/" + c.ProjectId,
	}
	gcpClient, err := appengine.NewApplicationsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	resp, err := gcpClient.GetApplication(ctx, req, c.CallOptions...)
	if err != nil {
		return err
	}
	res <- resp
	return nil
}

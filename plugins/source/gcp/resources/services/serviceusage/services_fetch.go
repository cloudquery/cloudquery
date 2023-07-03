package serviceusage

import (
	"context"

	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	pb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	if len(c.EnabledServices) > 0 {
		for _, svc := range c.EnabledServices[c.ProjectId] {
			res <- svc.(*pb.Service)
		}
		return nil
	}
	req := &pb.ListServicesRequest{
		Parent:   "projects/" + c.ProjectId,
		PageSize: 200,
		Filter:   "state:ENABLED",
	}
	gcpClient, err := serviceusage.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListServices(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}

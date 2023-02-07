package compute

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func Routers() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_routers",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/routers/list#response-body`,
		Resolver:    fetchRouters,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Router{}, append(client.Options(), transformers.WithPrimaryKeys("Id"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
		},
		Relations: []*schema.Table{routerNatMappingInfos()},
	}
}

func fetchRouters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListRoutersRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewRoutersRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.AggregatedList(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp.Value.Routers
	}
	return nil
}

package compute

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	compute "cloud.google.com/go/compute/apiv1"
)

func TargetHttpProxies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_target_http_proxies",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/targetHttpProxies#TargetHttpProxy`,
		Resolver:    fetchTargetHttpProxies,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.TargetHttpProxy{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLink"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchTargetHttpProxies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListTargetHttpProxiesRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewTargetHttpProxiesRESTClient(ctx, c.ClientOptions...)
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

		res <- resp.Value.TargetHttpProxies
	}
	return nil
}

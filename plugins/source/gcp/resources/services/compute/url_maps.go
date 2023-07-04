package compute

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	compute "cloud.google.com/go/compute/apiv1"
)

func UrlMaps() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_url_maps",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/urlMaps#UrlMap`,
		Resolver:    fetchUrlMaps,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.UrlMap{}, transformers.WithPrimaryKeys("SelfLink")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func fetchUrlMaps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListUrlMapsRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewUrlMapsRESTClient(ctx, c.ClientOptions...)
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

		res <- resp.Value.UrlMaps
	}
	return nil
}

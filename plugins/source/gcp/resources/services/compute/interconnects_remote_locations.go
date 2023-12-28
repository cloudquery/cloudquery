package compute

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"

	compute "cloud.google.com/go/compute/apiv1"
)

func InterconnectRemoteLocations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_interconnect_remote_locations",
		Description: " https://cloud.google.com/compute/docs/reference/rest/v1/interconnectRemoteLocations#InterconnectRemoteLocationsGetResponse",
		Resolver:    fetchInterconnectRemoteLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.InterconnectRemoteLocation{}, transformers.WithPrimaryKeys("SelfLink")),
		Columns: []schema.Column{
			client.ProjectIDColumn(false),
		},
	}
}

func fetchInterconnectRemoteLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListInterconnectRemoteLocationsRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewInterconnectRemoteLocationsRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.List(ctx, req, c.CallOptions...)
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

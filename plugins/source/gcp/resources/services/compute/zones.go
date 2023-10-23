package compute

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"

	compute "cloud.google.com/go/compute/apiv1"
)

func Zones() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_zones",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/zones/list#response-body`,
		Resolver:    fetchZones,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Zone{}, transformers.WithPrimaryKeys("SelfLink")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
		Relations: schema.Tables{
			machineTypes(),
			osConfigInventories(),
		},
	}
}

func fetchZones(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	var maxResults = uint32(500)
	req := &pb.ListZonesRequest{
		Project:    c.ProjectId,
		MaxResults: &maxResults,
	}
	gcpClient, err := compute.NewZonesRESTClient(ctx, c.ClientOptions...)
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

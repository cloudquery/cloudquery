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

func Zones() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_zones",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/zones/list#response-body`,
		Resolver:    fetchZones,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Zone{}, append(client.Options(), transformers.WithPrimaryKeys("SelfLink"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
		},
		Relations: []*schema.Table{machineTypes()},
	}
}

func fetchZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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

package compute

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/apache/arrow/go/v15/arrow"
	"google.golang.org/api/iterator"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ExternalVpnGateways() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_external_vpn_gateways",
		Description: "https://cloud.google.com/compute/docs/reference/rest/v1/externalVpnGateways#ExternalVpnGateway",
		Resolver:    fetchExternalVpnGateways,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.ExternalVpnGateway{}, transformers.WithPrimaryKeys("SelfLink")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func fetchExternalVpnGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListExternalVpnGatewaysRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewExternalVpnGatewaysRESTClient(ctx, c.ClientOptions...)
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

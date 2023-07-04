package compute

import (
	"context"
	"strings"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func routerNatMappingInfos() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_router_nat_mapping_infos",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/routers/getNatMappingInfo#response-body`,
		Resolver:    fetchRouterNatMappingInfo,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.VmEndpointNatMappings{}),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func fetchRouterNatMappingInfo(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := parent.Item.(*pb.Router)
	regionParts := strings.Split(*p.Region, "/")
	req := &pb.GetNatMappingInfoRoutersRequest{
		Project: c.ProjectId,
		Region:  regionParts[len(regionParts)-1],
		Router:  *p.Name,
	}
	gcpClient, err := compute.NewRoutersRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.GetNatMappingInfo(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			if strings.Contains(err.Error(), "No Nat mapping information is available in the given router.") {
				return nil
			}
			return err
		}

		res <- resp
	}
	return nil
}

package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeVpnGateways() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_vpn_gateways",
		Resolver:     fetchComputeVpnGateways,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "label_fingerprint",
				Type: schema.TypeString,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "network",
				Type: schema.TypeString,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_vpn_gateway_vpn_interfaces",
				Resolver: fetchComputeVpnGatewayVpnInterfaces,
				Columns: []schema.Column{
					{
						Name:     "vpn_gateway_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: client.ResolveResourceId,
					},
					{
						Name: "interconnect_attachment",
						Type: schema.TypeString,
					},
					{
						Name: "ip_address",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeVpnGateways(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.VpnGateways.AggregatedList(c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var vpnGateways []*compute.VpnGateway
		for _, items := range output.Items {
			vpnGateways = append(vpnGateways, items.VpnGateways...)
		}
		res <- vpnGateways

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchComputeVpnGatewayVpnInterfaces(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.VpnGateway)
	res <- r.VpnInterfaces
	return nil
}

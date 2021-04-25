package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeNetworks() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_networks",
		Resolver:     fetchComputeNetworks,
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
				Name:     "ip_v4_range",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IPv4Range"),
			},
			{
				Name: "auto_create_subnetworks",
				Type: schema.TypeBool,
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
				Name:     "gateway_ip_v4",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GatewayIPv4"),
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
				Name: "mtu",
				Type: schema.TypeBigInt,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "routing_config_routing_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoutingConfig.RoutingMode"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "subnetworks",
				Type: schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_network_peerings",
				Resolver: fetchComputeNetworkPeerings,
				Columns: []schema.Column{
					{
						Name:     "network_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "auto_create_routes",
						Type: schema.TypeBool,
					},
					{
						Name: "exchange_subnet_routes",
						Type: schema.TypeBool,
					},
					{
						Name: "export_custom_routes",
						Type: schema.TypeBool,
					},
					{
						Name: "export_subnet_routes_with_public_ip",
						Type: schema.TypeBool,
					},
					{
						Name: "import_custom_routes",
						Type: schema.TypeBool,
					},
					{
						Name: "import_subnet_routes_with_public_ip",
						Type: schema.TypeBool,
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
						Name: "peer_mtu",
						Type: schema.TypeBigInt,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name: "state_details",
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
func fetchComputeNetworks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	nextPageToken := ""
	c := meta.(*client.Client)
	for {
		call := c.Services.Compute.Networks.List(c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}
		res <- output.Items
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

//"google.golang.org/api/compute/v1"
func fetchComputeNetworkPeerings(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.Network)
	res <- r.Peerings
	return nil
}

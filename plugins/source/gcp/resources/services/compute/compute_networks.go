package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeNetworks() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_networks",
		Description: "Represents a VPC Network resource  Networks connect resources to each other and to the internet",
		Resolver:    fetchComputeNetworks,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "ip_v4_range",
				Description: "Deprecated in favor of subnet mode networks The range of internal addresses that are legal on this network This range is a CIDR specification, for example: 19216800/16 Provided by the client when the network is created",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IPv4Range"),
			},
			{
				Name:        "auto_create_subnetworks",
				Description: "Must be set to create a VPC network If not set, a legacy network is created  When set to true, the VPC network is created in auto mode When set to false, the VPC network is created in custom mode  An auto mode VPC network starts with one subnet per region Each subnet has a predetermined range as described in Auto mode VPC network IP ranges  For custom mode VPC networks, you can add subnets using the subnetworks insert method",
				Type:        schema.TypeBool,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this field when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "gateway_ip_v4",
				Description: "The gateway address for default routing out of the network, selected by GCP",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GatewayIPv4"),
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#network for networks",
				Type:        schema.TypeString,
			},
			{
				Name:        "mtu",
				Description: "Maximum Transmission Unit in bytes The minimum value for this field is 1460 and the maximum value is 1500 bytes",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created",
				Type:        schema.TypeString,
			},
			{
				Name:        "routing_config_routing_mode",
				Description: "The network-wide routing mode to use If set to REGIONAL, this network's Cloud Routers will only advertise routes with subnets of this network in the same region as the router If set to GLOBAL, this network's Cloud Routers will advertise routes with all subnets of this network, across regions",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoutingConfig.RoutingMode"),
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnetworks",
				Description: "Server-defined fully-qualified URLs for all subnetworks in this VPC network",
				Type:        schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_compute_network_peerings",
				Description: "A network peering attached to a network resource The message includes the peering name, peer network, peering state, and a flag indicating whether Google Compute Engine should automatically create routes for the peering",
				Resolver:    fetchComputeNetworkPeerings,
				Columns: []schema.Column{
					{
						Name:        "network_cq_id",
						Description: "Unique ID of gcp_compute_networks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "network_name",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("name"),
					},
					{
						Name:        "auto_create_routes",
						Description: "This field will be deprecated soon Use the exchange_subnet_routes field instead Indicates whether full mesh connectivity is created and managed automatically between peered networks Currently this field should always be true since Google Compute Engine will automatically create and manage subnetwork routes between two networks when peering state is ACTIVE",
						Type:        schema.TypeBool,
					},
					{
						Name:        "exchange_subnet_routes",
						Description: "Indicates whether full mesh connectivity is created and managed automatically between peered networks Currently this field should always be true since Google Compute Engine will automatically create and manage subnetwork routes between two networks when peering state is ACTIVE",
						Type:        schema.TypeBool,
					},
					{
						Name:        "export_custom_routes",
						Description: "Whether to export the custom routes to peer network",
						Type:        schema.TypeBool,
					},
					{
						Name:        "export_subnet_routes_with_public_ip",
						Description: "Whether subnet routes with public IP range are exported The default value is true, all subnet routes are exported The IPv4 special-use ranges (https://enwikipediaorg/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field",
						Type:        schema.TypeBool,
					},
					{
						Name:        "import_custom_routes",
						Description: "Whether to import the custom routes from peer network",
						Type:        schema.TypeBool,
					},
					{
						Name:        "import_subnet_routes_with_public_ip",
						Description: "Whether subnet routes with public IP range are imported The default value is false The IPv4 special-use ranges (https://enwikipediaorg/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field",
						Type:        schema.TypeBool,
					},
					{
						Name:        "name",
						Description: "Name of this peering Provided by the client when the peering is created The name must comply with RFC1035 Specifically, the name must be 1-63 characters long and match regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` The first character must be a lowercase letter, and all the following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash",
						Type:        schema.TypeString,
					},
					{
						Name:        "network",
						Description: "The URL of the peer network It can be either full URL or partial URL The peer network may belong to a different project If the partial URL does not contain project, it is assumed that the peer network is in the same project as the current network",
						Type:        schema.TypeString,
					},
					{
						Name:        "peer_mtu",
						Description: "Maximum Transmission Unit in bytes",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "state",
						Description: "State for the peering, either `ACTIVE` or `INACTIVE` The peering is `ACTIVE` when there's a matching configuration in the peer network",
						Type:        schema.TypeString,
					},
					{
						Name:        "state_details",
						Description: "Details about the current state of the peering",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	nextPageToken := ""
	c := meta.(*client.Client)
	for {
		output, err := c.Services.Compute.Networks.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Items
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchComputeNetworkPeerings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.Network)
	res <- r.Peerings
	return nil
}

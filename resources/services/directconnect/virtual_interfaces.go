package directconnect

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func DirectconnectVirtualInterfaces() *schema.Table {
	return &schema.Table{
		Name:         "aws_directconnect_virtual_interfaces",
		Description:  "Information about a virtual interface. A virtual interface (VLAN) transmits the traffic between the AWS Direct Connect location and the customer network",
		Resolver:     fetchDirectconnectVirtualInterfaces,
		Multiplex:    client.ServiceAccountRegionMultiplexer("directconnect"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "address_family",
				Description: "The address family for the BGP peer.",
				Type:        schema.TypeString,
			},
			{
				Name:        "amazon_address",
				Description: "The IP address assigned to the Amazon interface.",
				Type:        schema.TypeString,
			},
			{
				Name:        "amazon_side_asn",
				Description: "The autonomous system number (ASN) for the Amazon side of the connection.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "asn",
				Description: "The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration",
				Type:        schema.TypeInt,
			},
			{
				Name:        "auth_key",
				Description: "The authentication key for BGP configuration",
				Type:        schema.TypeString,
			},
			{
				Name:        "aws_device_v2",
				Description: "The Direct Connect endpoint on which the virtual interface terminates.",
				Type:        schema.TypeString,
			},
			{
				Name:        "connection_id",
				Description: "The ID of the connection.",
				Type:        schema.TypeString,
			},
			{
				Name:        "customer_address",
				Description: "The IP address assigned to the customer interface.",
				Type:        schema.TypeString,
			},
			{
				Name:        "customer_router_config",
				Description: "The customer router configuration.",
				Type:        schema.TypeString,
			},
			{
				Name:        "direct_connect_gateway_id",
				Description: "The ID of the Direct Connect gateway.",
				Type:        schema.TypeString,
			},
			{
				Name:        "jumbo_frame_capable",
				Description: "Indicates whether jumbo frames (9001 MTU) are supported.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "location",
				Description: "The location of the connection.",
				Type:        schema.TypeString,
			},
			{
				Name:        "mtu",
				Description: "The maximum transmission unit (MTU), in bytes",
				Type:        schema.TypeInt,
			},
			{
				Name:        "owner_account",
				Description: "The ID of the AWS account that owns the virtual interface.",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "The AWS Region where the virtual interface is located.",
				Type:        schema.TypeString,
			},
			{
				Name:        "route_filter_prefixes",
				Description: "The routes to be advertised to the AWS network in this Region",
				Type:        schema.TypeStringArray,
				Resolver:    resolveDirectconnectVirtualInterfaceRouteFilterPrefixes,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the virtual interface.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDirectconnectVirtualInterfaceTags,
			},
			{
				Name:        "virtual_gateway_id",
				Description: "The ID of the virtual private gateway",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The ID of the virtual interface.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualInterfaceId"),
			},
			{
				Name:        "virtual_interface_name",
				Description: "The name of the virtual interface assigned by the customer network",
				Type:        schema.TypeString,
			},
			{
				Name:        "virtual_interface_state",
				Description: "The state of the virtual interface",
				Type:        schema.TypeString,
			},
			{
				Name:        "virtual_interface_type",
				Description: "The type of virtual interface",
				Type:        schema.TypeString,
			},
			{
				Name:        "vlan",
				Description: "The ID of the VLAN.",
				Type:        schema.TypeInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_directconnect_virtual_interface_bgp_peers",
				Description: "Information about a BGP peer. ",
				Resolver:    fetchDirectconnectVirtualInterfaceBgpPeers,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"virtual_interface_cq_id", "bgp_peer_id"}},
				Columns: []schema.Column{
					{
						Name:        "virtual_interface_cq_id",
						Description: "Unique CloudQuery ID of aws_directconnect_virtual_interfaces table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "virtual_interface_id",
						Description: "The ID of the virtual interface.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "address_family",
						Description: "The address family for the BGP peer.",
						Type:        schema.TypeString,
					},
					{
						Name:        "amazon_address",
						Description: "The IP address assigned to the Amazon interface.",
						Type:        schema.TypeString,
					},
					{
						Name:        "asn",
						Description: "The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "auth_key",
						Description: "The authentication key for BGP configuration",
						Type:        schema.TypeString,
					},
					{
						Name:        "aws_device_v2",
						Description: "The Direct Connect endpoint on which the BGP peer terminates.",
						Type:        schema.TypeString,
					},
					{
						Name:        "bgp_peer_id",
						Description: "The ID of the BGP peer.",
						Type:        schema.TypeString,
					},
					{
						Name:        "bgp_peer_state",
						Description: "The state of the BGP peer",
						Type:        schema.TypeString,
					},
					{
						Name:        "bgp_status",
						Description: "The status of the BGP peer",
						Type:        schema.TypeString,
					},
					{
						Name:        "customer_address",
						Description: "The IP address assigned to the customer interface.",
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
func fetchDirectconnectVirtualInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config directconnect.DescribeVirtualInterfacesInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeVirtualInterfaces(ctx, &config, func(options *directconnect.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- output.VirtualInterfaces
	return nil
}
func resolveDirectconnectVirtualInterfaceRouteFilterPrefixes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.VirtualInterface)
	routeFilterPrefixes := make([]*string, len(r.RouteFilterPrefixes))
	for i, prefix := range r.RouteFilterPrefixes {
		routeFilterPrefixes[i] = prefix.Cidr
	}
	return resource.Set("route_filter_prefixes", routeFilterPrefixes)
}
func resolveDirectconnectVirtualInterfaceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.VirtualInterface)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchDirectconnectVirtualInterfaceBgpPeers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	virtualInterface, ok := parent.Item.(types.VirtualInterface)
	if !ok {
		return fmt.Errorf("not a direct connect virtual interface")
	}
	res <- virtualInterface.BgpPeers
	return nil
}

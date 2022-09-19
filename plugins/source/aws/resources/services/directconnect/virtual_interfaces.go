// Code generated by codegen; DO NOT EDIT.

package directconnect

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VirtualInterfaces() *schema.Table {
	return &schema.Table{
		Name:      "aws_directconnect_virtual_interfaces",
		Resolver:  fetchDirectconnectVirtualInterfaces,
		Multiplex: client.ServiceAccountRegionMultiplexer("directconnect"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveVirtualInterfaceARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualInterfaceId"),
			},
			{
				Name:     "address_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AddressFamily"),
			},
			{
				Name:     "amazon_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AmazonAddress"),
			},
			{
				Name:     "amazon_side_asn",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AmazonSideAsn"),
			},
			{
				Name:     "asn",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Asn"),
			},
			{
				Name:     "auth_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthKey"),
			},
			{
				Name:     "aws_device_v_2",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsDeviceV2"),
			},
			{
				Name:     "aws_logical_device_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsLogicalDeviceId"),
			},
			{
				Name:     "bgp_peers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BgpPeers"),
			},
			{
				Name:     "connection_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionId"),
			},
			{
				Name:     "customer_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerAddress"),
			},
			{
				Name:     "customer_router_config",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerRouterConfig"),
			},
			{
				Name:     "direct_connect_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectConnectGatewayId"),
			},
			{
				Name:     "jumbo_frame_capable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("JumboFrameCapable"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "mtu",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Mtu"),
			},
			{
				Name:     "owner_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerAccount"),
			},
			{
				Name:     "route_filter_prefixes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RouteFilterPrefixes"),
			},
			{
				Name:     "site_link_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SiteLinkEnabled"),
			},
			{
				Name:     "virtual_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualGatewayId"),
			},
			{
				Name:     "virtual_interface_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualInterfaceName"),
			},
			{
				Name:     "virtual_interface_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualInterfaceState"),
			},
			{
				Name:     "virtual_interface_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualInterfaceType"),
			},
			{
				Name:     "vlan",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Vlan"),
			},
		},
	}
}

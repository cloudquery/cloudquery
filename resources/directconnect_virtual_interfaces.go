package resources

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
		Resolver:     fetchDirectconnectVirtualInterfaces,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "address_family",
				Type: schema.TypeString,
			},
			{
				Name: "amazon_address",
				Type: schema.TypeString,
			},
			{
				Name: "amazon_side_asn",
				Type: schema.TypeBigInt,
			},
			{
				Name: "asn",
				Type: schema.TypeInt,
			},
			{
				Name: "auth_key",
				Type: schema.TypeString,
			},
			{
				Name: "aws_device_v2",
				Type: schema.TypeString,
			},
			{
				Name: "connection_id",
				Type: schema.TypeString,
			},
			{
				Name: "customer_address",
				Type: schema.TypeString,
			},
			{
				Name: "customer_router_config",
				Type: schema.TypeString,
			},
			{
				Name: "direct_connect_gateway_id",
				Type: schema.TypeString,
			},
			{
				Name: "jumbo_frame_capable",
				Type: schema.TypeBool,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name: "mtu",
				Type: schema.TypeInt,
			},
			{
				Name: "owner_account",
				Type: schema.TypeString,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "route_filter_prefixes",
				Type:     schema.TypeStringArray,
				Resolver: resolveVirtualInterfaceRouteFilterPrefixes,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDirectconnectVirtualInterfaceTags,
			},
			{
				Name: "virtual_gateway_id",
				Type: schema.TypeString,
			},
			{
				Name: "virtual_interface_id",
				Type: schema.TypeString,
			},
			{
				Name: "virtual_interface_name",
				Type: schema.TypeString,
			},
			{
				Name: "virtual_interface_state",
				Type: schema.TypeString,
			},
			{
				Name: "virtual_interface_type",
				Type: schema.TypeString,
			},
			{
				Name: "vlan",
				Type: schema.TypeInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_directconnect_virtual_interface_bgp_peers",
				Resolver: fetchVirtualInterfaceBGPPeers,
				Columns: []schema.Column{
					{
						Name:     "virtual_interface_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "address_family",
						Type: schema.TypeString,
					},
					{
						Name: "amazon_address",
						Type: schema.TypeString,
					},
					{
						Name: "asn",
						Type: schema.TypeInt,
					},
					{
						Name: "auth_key",
						Type: schema.TypeString,
					},
					{
						Name: "aws_device_v2",
						Type: schema.TypeString,
					},
					{
						Name: "bgp_peer_id",
						Type: schema.TypeString,
					},
					{
						Name: "bgp_peer_state",
						Type: schema.TypeString,
					},
					{
						Name: "bgp_status",
						Type: schema.TypeString,
					},
					{
						Name: "customer_address",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func resolveVirtualInterfaceRouteFilterPrefixes(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.VirtualInterface)
	routeFilterPrefixes := make([]*string, len(r.RouteFilterPrefixes))
	for i, prefix := range r.RouteFilterPrefixes {
		routeFilterPrefixes[i] = prefix.Cidr
	}
	return resource.Set("route_filter_prefixes", routeFilterPrefixes)
}

func fetchVirtualInterfaceBGPPeers(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	virtualInterface, ok := parent.Item.(types.VirtualInterface)
	if !ok {
		return fmt.Errorf("not a direct connect virtual interface")
	}
	res <- virtualInterface.BgpPeers
	return nil
}

func resolveDirectconnectVirtualInterfaceTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.VirtualInterface)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}

func fetchDirectconnectVirtualInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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

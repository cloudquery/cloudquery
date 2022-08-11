package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkRouteTables() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_route_tables",
		Description:  "Azure route table",
		Resolver:     fetchNetworkRouteTables,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "id",
				Description: "Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "route_table_subnets",
				Description: "A collection of references to subnets.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveNetworkRouteTableSubnets,
			},
			{
				Name:        "disable_bgp_route_propagation",
				Description: "Whether to disable the routes learned by BGP on that route table.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("RouteTablePropertiesFormat.DisableBgpRoutePropagation"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the route table resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RouteTablePropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the route table.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RouteTablePropertiesFormat.ResourceGUID"),
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_network_route_table_routes",
				Description: "Azure route table route",
				Resolver:    fetchNetworkRouteTableRoutes,
				Columns: []schema.Column{
					{
						Name:        "route_table_cq_id",
						Description: "Unique CloudQuery ID of azure_network_route_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "address_prefix",
						Description: "The destination CIDR to which the route applies.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutePropertiesFormat.AddressPrefix"),
					},
					{
						Name:        "next_hop_type",
						Description: "The type of Azure hop the packet should be sent to.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutePropertiesFormat.NextHopType"),
					},
					{
						Name:        "next_hop_ip_address",
						Description: "The IP address packets should be forwarded to.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutePropertiesFormat.NextHopIPAddress"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the route resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutePropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "has_bgp_override",
						Description: "A value indicating whether this route overrides overlapping BGP routes regardless of LPM.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("RoutePropertiesFormat.HasBgpOverride"),
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type.",
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
func fetchNetworkRouteTables(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.RouteTables
	response, err := svc.ListAll(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func resolveNetworkRouteTableSubnets(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rt := resource.Item.(network.RouteTable)
	if rt.Subnets == nil {
		return nil
	}
	subnets := make([]string, 0, len(*rt.Subnets))
	for _, sn := range *rt.Subnets {
		subnets = append(subnets, *sn.ID)
	}
	return diag.WrapError(resource.Set(c.Name, subnets))
}
func fetchNetworkRouteTableRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	rt := parent.Item.(network.RouteTable)
	if rt.Routes == nil {
		return nil
	}
	res <- *rt.Routes
	return nil
}

package network

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkRouteFilters() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_route_filters",
		Description:  "Azure Network Route Filters",
		Resolver:     fetchNetworkRouteFilters,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id.",
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
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ipv6_peerings",
				Description: "A collection of references to express route circuit ipv6 peerings.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkRouteFilterIpv6Peerings,
			},
			{
				Name:        "location",
				Description: "Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "peerings",
				Description: "A collection of references to express route circuit peerings.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkRouteFilterPeerings,
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the route filter resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RouteFilterPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_network_route_filter_rules",
				Description: "Route Filter Rule Resource.",
				Resolver:    fetchNetworkRouteFilterRules,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"route_filter_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "route_filter_cq_id",
						Description: "Unique CloudQuery ID of azure_network_route_filters table (FK)",
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
						Name:        "access",
						Description: "The access type of the rule.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RouteFilterRulePropertiesFormat.Access"),
					},
					{
						Name:        "communities",
						Description: "The collection for bgp community values to filter on.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("RouteFilterRulePropertiesFormat.Communities"),
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "location",
						Description: "Resource location.",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the route filter rule resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RouteFilterRulePropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "route_filter_rule_type",
						Description: "The rule type of the rule.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RouteFilterRulePropertiesFormat.RouteFilterRuleType"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchNetworkRouteFilters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.RouteFilters
	response, err := svc.List(ctx)
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
func resolveNetworkRouteFilterIpv6Peerings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rf, ok := resource.Item.(network.RouteFilter)
	if !ok {
		return fmt.Errorf("expected to have network.RouteFilter but got %T", resource.Item)
	}
	if rf.Ipv6Peerings == nil {
		return nil
	}
	return resource.Set(c.Name, *rf.Ipv6Peerings)
}
func resolveNetworkRouteFilterPeerings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rf, ok := resource.Item.(network.RouteFilter)
	if !ok {
		return fmt.Errorf("expected to have network.RouteFilter but got %T", resource.Item)
	}
	if rf.Peerings == nil {
		return nil
	}
	return resource.Set(c.Name, *rf.Peerings)
}
func fetchNetworkRouteFilterRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	rf, ok := parent.Item.(network.RouteFilter)
	if !ok {
		return fmt.Errorf("expected to have network.RouteFilter but got %T", parent.Item)
	}
	if rf.RouteFilterPropertiesFormat != nil && rf.RouteFilterPropertiesFormat.Rules != nil {
		res <- *rf.RouteFilterPropertiesFormat.Rules
	}
	return nil
}

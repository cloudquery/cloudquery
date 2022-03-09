package search

import (
	"context"
	"fmt"
	"net"

	"github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SearchServices() *schema.Table {
	return &schema.Table{
		Name:         "azure_search_services",
		Description:  "Service describes an Azure Cognitive Search service and its current state.",
		Resolver:     fetchSearchServices,
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
				Name:        "replica_count",
				Description: "The number of replicas in the search service",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ServiceProperties.ReplicaCount"),
			},
			{
				Name:        "partition_count",
				Description: "The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ServiceProperties.PartitionCount"),
			},
			{
				Name:        "hosting_mode",
				Description: "Applicable only for the standard3 SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceProperties.HostingMode"),
			},
			{
				Name:        "public_network_access",
				Description: "This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceProperties.PublicNetworkAccess"),
			},
			{
				Name:        "status",
				Description: "The status of the search service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceProperties.Status"),
			},
			{
				Name:        "status_details",
				Description: "The details of the search service status.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceProperties.StatusDetails"),
			},
			{
				Name:        "provisioning_state",
				Description: "The state of the last provisioning operation performed on the search service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceProperties.ProvisioningState"),
			},
			{
				Name:        "network_rule_set_ip_rules",
				Description: "A list of IP restriction rules that defines the inbound network(s) with allowing access to the search service endpoint",
				Type:        schema.TypeInetArray,
				Resolver:    resolveSearchServicesNetworkRuleSetIpRules,
			},
			{
				Name:        "sku_name",
				Description: "The SKU of the search service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "identity_principal_id",
				Description: "The principal ID of resource identity.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The tenant ID of resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "identity_type",
				Description: "The identity type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "location",
				Description: "The geo-location where the resource lives",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "Fully qualified resource ID for the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the resource",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_search_service_private_endpoint_connections",
				Description: "PrivateEndpointConnection describes an existing Private Endpoint connection to the Azure Cognitive Search service.",
				Resolver:    fetchSearchServicePrivateEndpointConnections,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of azure_search_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "private_endpoint_id",
						Description: "The resource id of the private endpoint resource from Microsoft.Network provider.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateEndpoint.ID"),
					},
					{
						Name:        "private_link_connection_status",
						Description: "Status of the the private link service connection",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "private_link_connection_description",
						Description: "The description for the private link service connection state.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "private_link_connection_actions_required",
						Description: "A description of any extra actions that may be required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.ActionsRequired"),
					},
					{
						Name:        "id",
						Description: "Fully qualified resource ID for the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The name of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the resource",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_search_service_shared_private_link_resources",
				Description: "SharedPrivateLinkResource describes a Shared Private Link Resource managed by the Azure Cognitive Search service.",
				Resolver:    fetchSearchServiceSharedPrivateLinkResources,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of azure_search_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "private_link_resource_id",
						Description: "The resource id of the resource the shared private link resource is for.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkResourceID"),
					},
					{
						Name:        "group_id",
						Description: "The group id from the provider of resource the shared private link resource is for.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.GroupID"),
					},
					{
						Name:        "request_message",
						Description: "The request message for requesting approval of the shared private link resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.RequestMessage"),
					},
					{
						Name:        "resource_region",
						Description: "Optional",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.ResourceRegion"),
					},
					{
						Name:        "status",
						Description: "Status of the shared private link resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.Status"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the shared private link resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.ProvisioningState"),
					},
					{
						Name:        "id",
						Description: "Fully qualified resource ID for the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The name of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the resource",
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

func fetchSearchServices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Search.Service
	response, err := svc.ListBySubscription(ctx, nil)
	if err != nil {
		return err
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
func resolveSearchServicesNetworkRuleSetIpRules(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	service, ok := resource.Item.(search.Service)
	if !ok {
		return fmt.Errorf("expected to have search.Service but got %T", resource.Item)
	}
	if service.NetworkRuleSet == nil || service.NetworkRuleSet.IPRules == nil {
		return nil
	}
	ipRules := make([]net.IP, len(*service.NetworkRuleSet.IPRules))
	for _, ipRule := range *service.NetworkRuleSet.IPRules {
		ipStr := *ipRule.Value
		ip := net.ParseIP(ipStr)
		if ipStr != "" && ip == nil {
			return fmt.Errorf("failed to parse IP from %s", ipStr)
		}
		if ip.To4() != nil {
			ipRules = append(ipRules, ip.To4())
		}
	}
	return resource.Set(c.Name, ipRules)
}
func fetchSearchServicePrivateEndpointConnections(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(search.Service)
	if !ok {
		return fmt.Errorf("expected to have search.Service but got %T", parent.Item)
	}
	if service.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *service.PrivateEndpointConnections
	return nil
}
func fetchSearchServiceSharedPrivateLinkResources(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	service, ok := parent.Item.(search.Service)
	if !ok {
		return fmt.Errorf("expected to have search.Service but got %T", parent.Item)
	}
	if service.SharedPrivateLinkResources == nil {
		return nil
	}
	res <- *service.SharedPrivateLinkResources
	return nil
}

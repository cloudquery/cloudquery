package network

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkExpressRoutePorts() *schema.Table {
	return &schema.Table{
		Name:          "azure_network_express_route_ports",
		Description:   "Azure Network Express Route Ports",
		Resolver:      fetchNetworkExpressRoutePorts,
		Multiplex:     client.SubscriptionMultiplex,
		DeleteFilter:  client.DeleteSubscriptionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		IgnoreInTests: true,
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
				Name:        "allocation_date",
				Description: "Date of the physical port allocation to be used in Letter of Authorization.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRoutePortPropertiesFormat.AllocationDate"),
			},
			{
				Name:        "bandwidth_in_gbps",
				Description: "Bandwidth of procured ports in Gbps.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ExpressRoutePortPropertiesFormat.BandwidthInGbps"),
			},
			{
				Name:        "circuits",
				Description: "Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveExpressRoutePortCircuits,
			},
			{
				Name:        "encapsulation",
				Description: "Encapsulation method on physical ports.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRoutePortPropertiesFormat.Encapsulation"),
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ether_type",
				Description: "Ether type of the physical port.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRoutePortPropertiesFormat.EtherType"),
			},
			{
				Name:        "identity_principal_id",
				Description: "The principal id of the system assigned identity.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The tenant id of the system assigned identity.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "identity_type",
				Description: "The type of identity used for the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_user_assigned_identities",
				Description: "The list of user identities associated with resource.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Identity.UserAssignedIdentities"),
			},
			{
				Name:        "location",
				Description: "Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "mtu",
				Description: "Maximum transmission unit of the physical port pair(s).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRoutePortPropertiesFormat.Mtu"),
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "peering_location",
				Description: "The name of the peering location that the ExpressRoutePort is mapped to physically.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRoutePortPropertiesFormat.PeeringLocation"),
			},
			{
				Name:        "provisioned_bandwidth_in_gbps",
				Description: "Aggregate Gbps of associated circuit bandwidths.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("ExpressRoutePortPropertiesFormat.ProvisionedBandwidthInGbps"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the express route port resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRoutePortPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the express route port resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRoutePortPropertiesFormat.ResourceGUID"),
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
				Name:        "azure_network_express_route_links",
				Description: "ExpressRouteLink resource.",
				Resolver:    fetchNetworkExpressRouteLinks,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"express_route_port_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "express_route_port_cq_id",
						Description: "Unique CloudQuery ID of azure_network_express_route_ports table (FK)",
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
						Name:        "admin_state",
						Description: "Administrative state of the physical port.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.AdminState"),
					},
					{
						Name:        "connector_type",
						Description: "Physical fiber port type.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.ConnectorType"),
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "interface_name",
						Description: "Name of Azure router interface.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.InterfaceName"),
					},
					{
						Name:        "mac_sec_config_cak_secret_identifier",
						Description: "Keyvault Secret Identifier URL containing Mac security CAK key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.MacSecConfig.CakSecretIdentifier"),
					},
					{
						Name:        "mac_sec_config_cipher",
						Description: "Mac security cipher.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.MacSecConfig.Cipher"),
					},
					{
						Name:        "mac_sec_config_ckn_secret_identifier",
						Description: "Keyvault Secret Identifier URL containing Mac security CKN key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.MacSecConfig.CknSecretIdentifier"),
					},
					{
						Name:        "mac_sec_config_sci_state",
						Description: "Sci mode enabled/disabled.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.MacSecConfig.SciState"),
					},
					{
						Name:        "name",
						Description: "Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "patch_panel_id",
						Description: "Mapping between physical port to patch panel port.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.PatchPanelID"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the express route link resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "rack_id",
						Description: "Mapping of physical patch panel to rack.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.RackID"),
					},
					{
						Name:        "router_name",
						Description: "Name of Azure router associated with physical port.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteLinkPropertiesFormat.RouterName"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchNetworkExpressRoutePorts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.ExpressRoutePorts
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
func resolveExpressRoutePortCircuits(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	erp, ok := resource.Item.(network.ExpressRoutePort)
	if !ok {
		return fmt.Errorf("not a network.ExpressRoutePort instance: %T", resource.Item)
	}
	if erp.Circuits == nil {
		return nil
	}
	items := *erp.Circuits
	ids := make([]string, 0, len(items))
	for _, cir := range items {
		if cir.ID != nil {
			ids = append(ids, *cir.ID)
		}
	}
	return resource.Set(c.Name, ids)
}
func fetchNetworkExpressRouteLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	erp, ok := parent.Item.(network.ExpressRoutePort)
	if !ok {
		return fmt.Errorf("expected to have network.ExpressRoutePort but got %T", parent.Item)
	}
	if erp.ExpressRoutePortPropertiesFormat != nil && erp.ExpressRoutePortPropertiesFormat.Links != nil {
		res <- *erp.ExpressRoutePortPropertiesFormat.Links
	}
	return nil
}
